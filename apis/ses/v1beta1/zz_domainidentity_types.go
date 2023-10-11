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

type DomainIdentityInitParameters struct {
}

type DomainIdentityObservation struct {

	// The ARN of the domain identity.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done.  Find out more about verifying domains in Amazon
	// SES in the AWS SES
	// docs.
	VerificationToken *string `json:"verificationToken,omitempty" tf:"verification_token,omitempty"`
}

type DomainIdentityParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainIdentitySpec defines the desired state of DomainIdentity
type DomainIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainIdentityParameters `json:"forProvider"`
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
	InitProvider DomainIdentityInitParameters `json:"initProvider,omitempty"`
}

// DomainIdentityStatus defines the observed state of DomainIdentity.
type DomainIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainIdentity is the Schema for the DomainIdentitys API. Provides an SES domain identity resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainIdentitySpec   `json:"spec"`
	Status            DomainIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainIdentityList contains a list of DomainIdentitys
type DomainIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainIdentity `json:"items"`
}

// Repository type metadata.
var (
	DomainIdentity_Kind             = "DomainIdentity"
	DomainIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainIdentity_Kind}.String()
	DomainIdentity_KindAPIVersion   = DomainIdentity_Kind + "." + CRDGroupVersion.String()
	DomainIdentity_GroupVersionKind = CRDGroupVersion.WithKind(DomainIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainIdentity{}, &DomainIdentityList{})
}
