/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`
}

type FirewallConfigParameters struct {

	// +kubebuilder:validation:Optional
	FirewallFailOpen *string `json:"firewallFailOpen,omitempty" tf:"firewall_fail_open,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

// FirewallConfigSpec defines the desired state of FirewallConfig
type FirewallConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallConfigParameters `json:"forProvider"`
}

// FirewallConfigStatus defines the observed state of FirewallConfig.
type FirewallConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallConfig is the Schema for the FirewallConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FirewallConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallConfigSpec   `json:"spec"`
	Status            FirewallConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallConfigList contains a list of FirewallConfigs
type FirewallConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallConfig `json:"items"`
}

// Repository type metadata.
var (
	FirewallConfig_Kind             = "FirewallConfig"
	FirewallConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallConfig_Kind}.String()
	FirewallConfig_KindAPIVersion   = FirewallConfig_Kind + "." + CRDGroupVersion.String()
	FirewallConfig_GroupVersionKind = CRDGroupVersion.WithKind(FirewallConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallConfig{}, &FirewallConfigList{})
}
