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

type EIPInitParameters struct {

	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIP *string `json:"associateWithPrivateIp,omitempty" tf:"associate_with_private_ip,omitempty"`

	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out Customer-owned IP addresses guide.
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup *string `json:"networkBorderGroup,omitempty" tf:"network_border_group,omitempty"`

	// EC2 IPv4 address pool identifier or amazon.
	// This option is only available for VPC EIPs.
	PublicIPv4Pool *string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean if the EIP is in a VPC or not.
	// Defaults to true unless the region supports EC2-Classic.
	VPC *bool `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type EIPObservation struct {

	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIP *string `json:"associateWithPrivateIp,omitempty" tf:"associate_with_private_ip,omitempty"`

	// ID representing the association of the address with an instance in a VPC.
	AssociationID *string `json:"associationId,omitempty" tf:"association_id,omitempty"`

	// Carrier IP address.
	CarrierIP *string `json:"carrierIp,omitempty" tf:"carrier_ip,omitempty"`

	// Customer owned IP.
	CustomerOwnedIP *string `json:"customerOwnedIp,omitempty" tf:"customer_owned_ip,omitempty"`

	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out Customer-owned IP addresses guide.
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// Indicates if this EIP is for use in VPC (vpc) or EC2-Classic (standard).
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Contains the EIP allocation ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// EC2 instance ID.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup *string `json:"networkBorderGroup,omitempty" tf:"network_border_group,omitempty"`

	// Network interface ID to associate with.
	NetworkInterface *string `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The Private DNS associated with the Elastic IP address (if in VPC).
	PrivateDNS *string `json:"privateDns,omitempty" tf:"private_dns,omitempty"`

	// Contains the private IP address (if in VPC).
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Public DNS associated with the Elastic IP address.
	PublicDNS *string `json:"publicDns,omitempty" tf:"public_dns,omitempty"`

	// Contains the public IP address.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// EC2 IPv4 address pool identifier or amazon.
	// This option is only available for VPC EIPs.
	PublicIPv4Pool *string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Boolean if the EIP is in a VPC or not.
	// Defaults to true unless the region supports EC2-Classic.
	VPC *bool `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

type EIPParameters struct {

	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	// +kubebuilder:validation:Optional
	AssociateWithPrivateIP *string `json:"associateWithPrivateIp,omitempty" tf:"associate_with_private_ip,omitempty"`

	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out Customer-owned IP addresses guide.
	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// EC2 instance ID.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	// +kubebuilder:validation:Optional
	NetworkBorderGroup *string `json:"networkBorderGroup,omitempty" tf:"network_border_group,omitempty"`

	// Network interface ID to associate with.
	// +crossplane:generate:reference:type=NetworkInterface
	// +kubebuilder:validation:Optional
	NetworkInterface *string `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// Reference to a NetworkInterface to populate networkInterface.
	// +kubebuilder:validation:Optional
	NetworkInterfaceRef *v1.Reference `json:"networkInterfaceRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterface.
	// +kubebuilder:validation:Optional
	NetworkInterfaceSelector *v1.Selector `json:"networkInterfaceSelector,omitempty" tf:"-"`

	// EC2 IPv4 address pool identifier or amazon.
	// This option is only available for VPC EIPs.
	// +kubebuilder:validation:Optional
	PublicIPv4Pool *string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean if the EIP is in a VPC or not.
	// Defaults to true unless the region supports EC2-Classic.
	// +kubebuilder:validation:Optional
	VPC *bool `json:"vpc,omitempty" tf:"vpc,omitempty"`
}

// EIPSpec defines the desired state of EIP
type EIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EIPParameters `json:"forProvider"`
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
	InitProvider EIPInitParameters `json:"initProvider,omitempty"`
}

// EIPStatus defines the observed state of EIP.
type EIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EIP is the Schema for the EIPs API. Provides an Elastic IP resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EIPSpec   `json:"spec"`
	Status            EIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EIPList contains a list of EIPs
type EIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EIP `json:"items"`
}

// Repository type metadata.
var (
	EIP_Kind             = "EIP"
	EIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EIP_Kind}.String()
	EIP_KindAPIVersion   = EIP_Kind + "." + CRDGroupVersion.String()
	EIP_GroupVersionKind = CRDGroupVersion.WithKind(EIP_Kind)
)

func init() {
	SchemeBuilder.Register(&EIP{}, &EIPList{})
}
