// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ClusterCapacityProvidersDefaultCapacityProviderStrategyInitParameters struct {

	// The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to 0.
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// Name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider,omitempty" tf:"capacity_provider,omitempty"`

	// The relative percentage of the total number of launched tasks that should use the specified capacity provider. The weight value is taken into consideration after the base count of tasks has been satisfied. Defaults to 0.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ClusterCapacityProvidersDefaultCapacityProviderStrategyObservation struct {

	// The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to 0.
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// Name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider,omitempty" tf:"capacity_provider,omitempty"`

	// The relative percentage of the total number of launched tasks that should use the specified capacity provider. The weight value is taken into consideration after the base count of tasks has been satisfied. Defaults to 0.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ClusterCapacityProvidersDefaultCapacityProviderStrategyParameters struct {

	// The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to 0.
	// +kubebuilder:validation:Optional
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// Name of the capacity provider.
	// +kubebuilder:validation:Optional
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider,omitempty"`

	// The relative percentage of the total number of launched tasks that should use the specified capacity provider. The weight value is taken into consideration after the base count of tasks has been satisfied. Defaults to 0.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ClusterCapacityProvidersInitParameters struct {

	// Set of names of one or more capacity providers to associate with the cluster. Valid values also include FARGATE and FARGATE_SPOT.
	CapacityProviders []*string `json:"capacityProviders,omitempty" tf:"capacity_providers,omitempty"`

	// Set of capacity provider strategies to use by default for the cluster. Detailed below.
	DefaultCapacityProviderStrategy []ClusterCapacityProvidersDefaultCapacityProviderStrategyInitParameters `json:"defaultCapacityProviderStrategy,omitempty" tf:"default_capacity_provider_strategy,omitempty"`
}

type ClusterCapacityProvidersObservation struct {

	// Set of names of one or more capacity providers to associate with the cluster. Valid values also include FARGATE and FARGATE_SPOT.
	CapacityProviders []*string `json:"capacityProviders,omitempty" tf:"capacity_providers,omitempty"`

	// Name of the ECS cluster to manage capacity providers for.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Set of capacity provider strategies to use by default for the cluster. Detailed below.
	DefaultCapacityProviderStrategy []ClusterCapacityProvidersDefaultCapacityProviderStrategyObservation `json:"defaultCapacityProviderStrategy,omitempty" tf:"default_capacity_provider_strategy,omitempty"`

	// Same as cluster_name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterCapacityProvidersParameters struct {

	// Set of names of one or more capacity providers to associate with the cluster. Valid values also include FARGATE and FARGATE_SPOT.
	// +kubebuilder:validation:Optional
	CapacityProviders []*string `json:"capacityProviders,omitempty" tf:"capacity_providers,omitempty"`

	// Name of the ECS cluster to manage capacity providers for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ecs/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in ecs to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in ecs to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// Set of capacity provider strategies to use by default for the cluster. Detailed below.
	// +kubebuilder:validation:Optional
	DefaultCapacityProviderStrategy []ClusterCapacityProvidersDefaultCapacityProviderStrategyParameters `json:"defaultCapacityProviderStrategy,omitempty" tf:"default_capacity_provider_strategy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ClusterCapacityProvidersSpec defines the desired state of ClusterCapacityProviders
type ClusterCapacityProvidersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterCapacityProvidersParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClusterCapacityProvidersInitParameters `json:"initProvider,omitempty"`
}

// ClusterCapacityProvidersStatus defines the observed state of ClusterCapacityProviders.
type ClusterCapacityProvidersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterCapacityProvidersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterCapacityProviders is the Schema for the ClusterCapacityProviderss API. Provides an ECS cluster capacity providers resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterCapacityProviders struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterCapacityProvidersSpec   `json:"spec"`
	Status            ClusterCapacityProvidersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterCapacityProvidersList contains a list of ClusterCapacityProviderss
type ClusterCapacityProvidersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterCapacityProviders `json:"items"`
}

// Repository type metadata.
var (
	ClusterCapacityProviders_Kind             = "ClusterCapacityProviders"
	ClusterCapacityProviders_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterCapacityProviders_Kind}.String()
	ClusterCapacityProviders_KindAPIVersion   = ClusterCapacityProviders_Kind + "." + CRDGroupVersion.String()
	ClusterCapacityProviders_GroupVersionKind = CRDGroupVersion.WithKind(ClusterCapacityProviders_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterCapacityProviders{}, &ClusterCapacityProvidersList{})
}
