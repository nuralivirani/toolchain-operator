package v1alpha1

import (
	toolchainv1alpha1 "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheInstallationSpec defines the desired state of CheInstallation
// +k8s:openapi-gen=true
type CheInstallationSpec struct {
	// The configuration required for che operator
	CheOperatorSpec CheOperator `json:"cheOperatorSpec"`
}

type CheOperator struct {
	// The namespace where the Che operator will be installed
	Namespace string `json:"namespace"`
}

// CheInstallationStatus defines the observed state of CheInstallation
// +k8s:openapi-gen=true
type CheInstallationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current CheInstallation conditions
	// Supported condition types:
	// CheInstalled, FailedToInstallChe
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []toolchainv1alpha1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CheInstallation is the Schema for the cheinstallations API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cheinstallations,scope=Cluster
type CheInstallation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CheInstallationSpec   `json:"spec,omitempty"`
	Status CheInstallationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CheInstallationList contains a list of CheInstallation
type CheInstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CheInstallation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CheInstallation{}, &CheInstallationList{})
}
