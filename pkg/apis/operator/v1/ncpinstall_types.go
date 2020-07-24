package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NcpInstallSpec defines the desired state of NcpInstall
type NcpInstallSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Replicas number for nsx-ncp deployment
	// Operator will ignore the value if NCP HA is disabled
	// +kubebuilder:validation:Minimum=1
	// +optional
	NcpReplicas int32 `json:"ncpReplicas,omitempty"`
}

// NcpInstallStatus defines the observed state of NcpInstall
type NcpInstallStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions describes the state of NCP installation.
	// +optional
	Conditions []InstallCondition `json:"conditions,omitempty"`
}

type InstallCondition = configv1.ClusterOperatorStatusCondition

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NcpInstall is the Schema for the ncpinstalls API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ncpinstalls,scope=Namespaced
type NcpInstall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NcpInstallSpec   `json:"spec,omitempty"`
	Status NcpInstallStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NcpInstallList contains a list of NcpInstall
type NcpInstallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NcpInstall `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NcpInstall{}, &NcpInstallList{})
}
