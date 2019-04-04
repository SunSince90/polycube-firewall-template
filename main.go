package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	k8sfirewall "github.com/SunSince90/polycube/src/components/k8s/utils/k8sfirewall"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta "github.com/SunSince90/polycube-firewall-template/pkg/apis/polycubenetwork.com/v1beta"

	log "github.com/Sirupsen/logrus"

	"github.com/SunSince90/polycube-firewall-template/controller"
	fwt_clientset "github.com/SunSince90/polycube-firewall-template/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubernetesClient() (kubernetes.Interface, fwt_clientset.Interface) {

	kubeconfig := os.Getenv("HOME") + "/.kube/config"

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	var err1 error
	clientset, err1 := kubernetes.NewForConfig(config)
	if err1 != nil {
		panic(err1.Error())
	}

	fclientset, err := fwt_clientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	log.Info("Successfully constructed k8s client")
	return clientset, fclientset
}

func main() {
	log.Infoln("Hello, World!")

	kclientset, fclientset := getKubernetesClient()
	c := controller.NewPcnFirewallTemplateController(kclientset, fclientset)

	// use a channel to synchronize the finalization for a graceful shutdown
	stopCh := make(chan struct{})
	defer close(stopCh)

	// run the controller loop to process items
	go c.Run(stopCh)

	go func() {
		time.Sleep(10 * time.Second)
		log.Infoln("Going to put a new one!")
		multiple(fclientset)
	}()

	// use a channel to handle OS signals to terminate and gracefully shut
	// down processing
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}

func multiple(f fwt_clientset.Interface) {
	rules := []k8sfirewall.ChainRule{}
	for i := 0; i < 1000; i++ {
		rules = append(rules, k8sfirewall.ChainRule{
			Id:        int32(i),
			Src:       "255.255.255.255",
			Dst:       "255.255.255.255",
			Conntrack: "established",
			L4proto:   "TCP",
			Sport:     8080,
			Dport:     9090,
		})
	}
	firewallClientSet := f.PolycubenetworkV1beta().FirewallTemplates(meta_v1.NamespaceDefault)
	fwt := v1beta.FirewallTemplate{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: "hello-world",
		},
		Spec: v1beta.FirewallTemplateSpec{
			DefaultActions: map[string]v1beta.FirewallTemplateDefaultAction{
				"ingress": v1beta.FirewallTemplateDefaultAction{
					Action:     v1beta.Forward,
					LastUpdate: time.Now().Unix(),
				},
			},
			Rules: rules,
		},
	}

	log.Infoln("Putting rules at", time.Now().Unix())
	_, err := firewallClientSet.Create(&fwt)
	if err != nil {
		log.Infoln("error:", err)
	}
}
