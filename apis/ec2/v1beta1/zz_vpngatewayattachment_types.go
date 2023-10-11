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

type VPNGatewayAttachmentInitParameters struct {
}

type VPNGatewayAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The ID of the Virtual Private Gateway.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type VPNGatewayAttachmentParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Private Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`
}

// VPNGatewayAttachmentSpec defines the desired state of VPNGatewayAttachment
type VPNGatewayAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayAttachmentParameters `json:"forProvider"`
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
	InitProvider VPNGatewayAttachmentInitParameters `json:"initProvider,omitempty"`
}

// VPNGatewayAttachmentStatus defines the observed state of VPNGatewayAttachment.
type VPNGatewayAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayAttachment is the Schema for the VPNGatewayAttachments API. Provides a Virtual Private Gateway attachment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPNGatewayAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewayAttachmentSpec   `json:"spec"`
	Status            VPNGatewayAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayAttachmentList contains a list of VPNGatewayAttachments
type VPNGatewayAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGatewayAttachment `json:"items"`
}

// Repository type metadata.
var (
	VPNGatewayAttachment_Kind             = "VPNGatewayAttachment"
	VPNGatewayAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGatewayAttachment_Kind}.String()
	VPNGatewayAttachment_KindAPIVersion   = VPNGatewayAttachment_Kind + "." + CRDGroupVersion.String()
	VPNGatewayAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VPNGatewayAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGatewayAttachment{}, &VPNGatewayAttachmentList{})
}
