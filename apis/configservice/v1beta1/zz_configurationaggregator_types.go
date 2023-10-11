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

type AccountAggregationSourceInitParameters struct {

	// List of 12-digit account IDs of the account(s) being aggregated.
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	// If true, aggregate existing AWS Config regions and future regions.
	AllRegions *bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`

	// List of source regions being aggregated.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type AccountAggregationSourceObservation struct {

	// List of 12-digit account IDs of the account(s) being aggregated.
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	// If true, aggregate existing AWS Config regions and future regions.
	AllRegions *bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`

	// List of source regions being aggregated.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type AccountAggregationSourceParameters struct {

	// List of 12-digit account IDs of the account(s) being aggregated.
	// +kubebuilder:validation:Optional
	AccountIds []*string `json:"accountIds" tf:"account_ids,omitempty"`

	// If true, aggregate existing AWS Config regions and future regions.
	// +kubebuilder:validation:Optional
	AllRegions *bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`

	// List of source regions being aggregated.
	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type ConfigurationAggregatorInitParameters struct {

	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource []AccountAggregationSourceInitParameters `json:"accountAggregationSource,omitempty" tf:"account_aggregation_source,omitempty"`

	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource []OrganizationAggregationSourceInitParameters `json:"organizationAggregationSource,omitempty" tf:"organization_aggregation_source,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationAggregatorObservation struct {

	// The account(s) to aggregate config data from as documented below.
	AccountAggregationSource []AccountAggregationSourceObservation `json:"accountAggregationSource,omitempty" tf:"account_aggregation_source,omitempty"`

	// The ARN of the aggregator
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization to aggregate config data from as documented below.
	OrganizationAggregationSource []OrganizationAggregationSourceObservation `json:"organizationAggregationSource,omitempty" tf:"organization_aggregation_source,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConfigurationAggregatorParameters struct {

	// The account(s) to aggregate config data from as documented below.
	// +kubebuilder:validation:Optional
	AccountAggregationSource []AccountAggregationSourceParameters `json:"accountAggregationSource,omitempty" tf:"account_aggregation_source,omitempty"`

	// The organization to aggregate config data from as documented below.
	// +kubebuilder:validation:Optional
	OrganizationAggregationSource []OrganizationAggregationSourceParameters `json:"organizationAggregationSource,omitempty" tf:"organization_aggregation_source,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OrganizationAggregationSourceInitParameters struct {

	// If true, aggregate existing AWS Config regions and future regions.
	AllRegions *bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`

	// List of source regions being aggregated.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type OrganizationAggregationSourceObservation struct {

	// If true, aggregate existing AWS Config regions and future regions.
	AllRegions *bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`

	// List of source regions being aggregated.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// ARN of the IAM role used to retrieve AWS Organization details associated with the aggregator account.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type OrganizationAggregationSourceParameters struct {

	// If true, aggregate existing AWS Config regions and future regions.
	// +kubebuilder:validation:Optional
	AllRegions *bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`

	// List of source regions being aggregated.
	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// ARN of the IAM role used to retrieve AWS Organization details associated with the aggregator account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// ConfigurationAggregatorSpec defines the desired state of ConfigurationAggregator
type ConfigurationAggregatorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationAggregatorParameters `json:"forProvider"`
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
	InitProvider ConfigurationAggregatorInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationAggregatorStatus defines the observed state of ConfigurationAggregator.
type ConfigurationAggregatorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationAggregatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationAggregator is the Schema for the ConfigurationAggregators API. Manages an AWS Config Configuration Aggregator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigurationAggregator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationAggregatorSpec   `json:"spec"`
	Status            ConfigurationAggregatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationAggregatorList contains a list of ConfigurationAggregators
type ConfigurationAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationAggregator `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationAggregator_Kind             = "ConfigurationAggregator"
	ConfigurationAggregator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationAggregator_Kind}.String()
	ConfigurationAggregator_KindAPIVersion   = ConfigurationAggregator_Kind + "." + CRDGroupVersion.String()
	ConfigurationAggregator_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationAggregator_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationAggregator{}, &ConfigurationAggregatorList{})
}
