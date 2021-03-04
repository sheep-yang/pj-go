package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AdminSpec defines the desired state of Admin
type AdminSpec struct {
Username string `json:"username,omitempty"`
Password string `json:"password,omitempty"`
Email    string `json:"email,omitempty"`
}


// AdminStatus defines the observed state of Admin
type AdminStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Admin is the Schema for the admins API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=admins,scope=Namespaced
type Admin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdminSpec   `json:"spec,omitempty"`
	Status AdminStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AdminList contains a list of Admin
type AdminList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Admin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Admin{}, &AdminList{})
}
