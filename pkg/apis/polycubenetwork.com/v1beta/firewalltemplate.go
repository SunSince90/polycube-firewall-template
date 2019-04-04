package v1beta

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
	Message string `json:"message,omitempty"`
}

// FirewallTemplateList contains a list of Firewall Templates.
type FirewallTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`
	// Items contains the firewall tempaltes
	Items []FirewallTemplate `json:"items"`
}
