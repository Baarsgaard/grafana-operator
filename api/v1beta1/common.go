package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ValueFrom struct {
	TargetPath string          `json:"targetPath"`
	ValueFrom  ValueFromSource `json:"valueFrom"`
}

// +kubebuilder:validation:XValidation:rule="(has(self.configMapKeyRef) && !has(self.secretKeyRef)) || (!has(self.configMapKeyRef) && has(self.secretKeyRef))", message="Either configMapKeyRef or secretKeyRef must be set"
type ValueFromSource struct {
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *v1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// Selects a key of a Secret.
	// +optional
	SecretKeyRef *v1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="((!has(oldSelf.uid) && !has(self.uid)) || (has(oldSelf.uid) && has(self.uid)))", message="spec.uid is immutable"
type GrafanaUIDSpec struct {
	// Manually specify the UID
	// +optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="spec.uid is immutable"
	CustomUID string `json:"uid,omitempty"`
}

type GrafanaCommonSpec struct {
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=duration
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$"
	// +kubebuilder:default="10m"
	ResyncPeriod metav1.Duration `json:"resyncPeriod,omitempty"`

	// selects Grafanas for import
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	InstanceSelector *metav1.LabelSelector `json:"instanceSelector"`

	// +optional
	AllowCrossNamespaceImport *bool `json:"allowCrossNamespaceImport,omitempty"`
}

// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
// WARN: Run "make" to regenerate code after modifying this file

// The most recent observed state of a Grafana resource
type GrafanaCommonStatus struct {
	// Detect resource changes
	Hash string `json:"hash,omitempty"`
	// If the instanceSelector can't find matching grafana instances
	NoMatchingInstances bool   `json:"NoMatchingInstances,omitempty"`
	LastMessage         string `json:"lastMessage,omitempty"`
	// Last time the resource was synchronized
	LastResync metav1.Time        `json:"lastResync,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
