// Copyright (c) 2020, 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package v1alpha1

import (
	oamrt "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetricsTraitKind is the Kind of the MetricsTrait
const MetricsTraitKind string = "MetricsTrait"

func init() {
	SchemeBuilder.Register(&MetricsTrait{}, &MetricsTraitList{})
}

// MetricsTraitList contains a list of metrics traits.
// +kubebuilder:object:root=true
type MetricsTraitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricsTrait `json:"items"`
}

// MetricsTrait specifies the metrics trait API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type MetricsTrait struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetricsTraitSpec   `json:"spec,omitempty"`
	Status MetricsTraitStatus `json:"status,omitempty"`
}

// MetricsTraitSpec specifies the desired state of a metrics trait.
type MetricsTraitSpec struct {
	// The HTTP port for the related metrics trait. Defaults to 8080.
	Port *int `json:"port,omitempty"`

	// The HTTP ports for the related metrics trait. Defaults to 8080.
	Ports []PortSpec `json:"ports,omitempty"`

	// The HTTP path for the related metrics endpoint. Defaults to /metrics.
	Path *string `json:"path,omitempty"`

	// The name of an opaque secret (i.e. username and password) within the workload's namespace for metrics endpoint access.
	Secret *string `json:"secret,omitempty"`

	// The prometheus deployment used to scrape the related metrics endpoints.
	// Defaults to istio-system/prometheus
	Scraper *string `json:"scraper,omitempty"`

	// Enabled specifies whether metrics collection is enabled. Defaults to true.
	//+optional
	Enabled *bool `json:"enabled,omitempty"`

	// A reference to the workload used to generate this metrics trait.
	WorkloadReference oamrt.TypedReference `json:"workloadRef"`
}

type PortSpec struct {
	// The HTTP port for the related metrics trait. Defaults to 8080.
	Port *int `json:"port,omitempty"`

	// The HTTP path for the related metrics endpoint. Defaults to /metrics.
	Path *string `json:"path,omitempty"`
}

// MetricsTraitStatus defines the observed state of MetricsTrait and related resources.
type MetricsTraitStatus struct {
	// Important: Run code generation after modifying this file.

	// The reconcile status of this metrics trait
	oamrt.ConditionedStatus `json:",inline"`

	// Related resources affected by this metrics trait
	Resources []QualifiedResourceRelation `json:"resources,omitempty"`
}
