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

type ConfigurationInitParameters struct {

	// A collection of parameters for this group configuration item. See below for details.
	Parameters []ParametersInitParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Specifies the type of group configuration item.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigurationObservation struct {

	// A collection of parameters for this group configuration item. See below for details.
	Parameters []ParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Specifies the type of group configuration item.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigurationParameters struct {

	// A collection of parameters for this group configuration item. See below for details.
	// +kubebuilder:validation:Optional
	Parameters []ParametersParameters `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Specifies the type of group configuration item.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type GroupInitParameters struct {

	// A configuration associates the resource group with an AWS service and specifies how the service can interact with the resources in the group. See below for details.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A description of the resource group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A resource_query block. Resource queries are documented below.
	ResourceQuery []ResourceQueryInitParameters `json:"resourceQuery,omitempty" tf:"resource_query,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GroupObservation struct {

	// The ARN assigned by AWS for this resource group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A configuration associates the resource group with an AWS service and specifies how the service can interact with the resources in the group. See below for details.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A description of the resource group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A resource_query block. Resource queries are documented below.
	ResourceQuery []ResourceQueryObservation `json:"resourceQuery,omitempty" tf:"resource_query,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GroupParameters struct {

	// A configuration associates the resource group with an AWS service and specifies how the service can interact with the resources in the group. See below for details.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A description of the resource group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A resource_query block. Resource queries are documented below.
	// +kubebuilder:validation:Optional
	ResourceQuery []ResourceQueryParameters `json:"resourceQuery,omitempty" tf:"resource_query,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParametersInitParameters struct {

	// The name of the group configuration parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value or values to be used for the specified parameter.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ParametersObservation struct {

	// The name of the group configuration parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value or values to be used for the specified parameter.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ParametersParameters struct {

	// The name of the group configuration parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value or values to be used for the specified parameter.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ResourceQueryInitParameters struct {

	// The resource query as a JSON string.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The type of the resource query. Defaults to TAG_FILTERS_1_0.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceQueryObservation struct {

	// The resource query as a JSON string.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The type of the resource query. Defaults to TAG_FILTERS_1_0.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceQueryParameters struct {

	// The resource query as a JSON string.
	// +kubebuilder:validation:Optional
	Query *string `json:"query" tf:"query,omitempty"`

	// The type of the resource query. Defaults to TAG_FILTERS_1_0.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. Provides a Resource Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
