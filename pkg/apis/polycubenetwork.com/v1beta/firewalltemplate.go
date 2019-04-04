package v1beta

import (
	k8sfirewall "github.com/SunSince90/polycube/src/components/k8s/utils/k8sfirewall"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirewallTemplate is a template of a firewall
type FirewallTemplate struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Status FirewallTemplateStatus `json:"status,omitempty"`
	// Spec of this policy
	Spec FirewallTemplateSpec `json:"spec,omitempty"`
}

// FirewallTemplateStatus defines the status of this firewall template
type FirewallTemplateStatus struct {
	Name string
}

// FirewallTemplateSpec contains the specifications of this firewall template
type FirewallTemplateSpec struct {
	DefaultActions map[string]FirewallTemplateDefaultAction
	Message        string `json:"message,omitempty"`
	Rules          []k8sfirewall.ChainRule
}

type FirewallTemplateDefaultAction struct {
	Action     FirewallTemplateDefaultActionType
	LastUpdate int64
}

type FirewallTemplateDefaultActionType string

const (
	Forward FirewallTemplateDefaultActionType = "forward"
	Drop    FirewallTemplateDefaultActionType = "drop"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirewallTemplateList contains a list of Firewall Templates.
type FirewallTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items contains the firewall tempaltes
	Items []FirewallTemplate `json:"items"`
}
