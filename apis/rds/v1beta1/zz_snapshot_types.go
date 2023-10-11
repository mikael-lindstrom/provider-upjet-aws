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

type SnapshotInitParameters struct {

	// List of AWS Account ids to share snapshot with, use all to make snaphot public.
	SharedAccounts []*string `json:"sharedAccounts,omitempty" tf:"shared_accounts,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SnapshotObservation struct {

	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage *float64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The DB Instance Identifier from which to take the snapshot.
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty" tf:"db_instance_identifier,omitempty"`

	// The Amazon Resource Name (ARN) for the DB snapshot.
	DBSnapshotArn *string `json:"dbSnapshotArn,omitempty" tf:"db_snapshot_arn,omitempty"`

	// Specifies whether the DB snapshot is encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Specifies the name of the database engine.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Specifies the version of the database engine.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The ARN for the KMS encryption key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// License model information for the restored DB instance.
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`

	// Provides the option group name for the DB snapshot.
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// List of AWS Account ids to share snapshot with, use all to make snaphot public.
	SharedAccounts []*string `json:"sharedAccounts,omitempty" tf:"shared_accounts,omitempty"`

	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDBSnapshotIdentifier *string `json:"sourceDbSnapshotIdentifier,omitempty" tf:"source_db_snapshot_identifier,omitempty"`

	// The region that the DB snapshot was created in or copied from.
	SourceRegion *string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`

	// Specifies the status of this DB snapshot.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies the storage type associated with DB snapshot.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Provides the VPC ID associated with the DB snapshot.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SnapshotParameters struct {

	// The DB Instance Identifier from which to take the snapshot.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty" tf:"db_instance_identifier,omitempty"`

	// Reference to a Instance in rds to populate dbInstanceIdentifier.
	// +kubebuilder:validation:Optional
	DBInstanceIdentifierRef *v1.Reference `json:"dbInstanceIdentifierRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate dbInstanceIdentifier.
	// +kubebuilder:validation:Optional
	DBInstanceIdentifierSelector *v1.Selector `json:"dbInstanceIdentifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// List of AWS Account ids to share snapshot with, use all to make snaphot public.
	// +kubebuilder:validation:Optional
	SharedAccounts []*string `json:"sharedAccounts,omitempty" tf:"shared_accounts,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Snapshot is the Schema for the Snapshots API. Manages an RDS database instance snapshot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
