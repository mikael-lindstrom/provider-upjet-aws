// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeliveryChannelInitParameters struct {

	// The name of the S3 bucket used to store the configuration history.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameRef *v1.Reference `json:"s3BucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameSelector *v1.Selector `json:"s3BucketNameSelector,omitempty" tf:"-"`

	// The ARN of the AWS KMS key used to encrypt objects delivered by AWS Config. Must belong to the same Region as the destination S3 bucket.
	S3KMSKeyArn *string `json:"s3KmsKeyArn,omitempty" tf:"s3_kms_key_arn,omitempty"`

	// The prefix for the specified S3 bucket.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties *SnapshotDeliveryPropertiesInitParameters `json:"snapshotDeliveryProperties,omitempty" tf:"snapshot_delivery_properties,omitempty"`

	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`
}

type DeliveryChannelObservation struct {

	// The name of the delivery channel.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the S3 bucket used to store the configuration history.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The ARN of the AWS KMS key used to encrypt objects delivered by AWS Config. Must belong to the same Region as the destination S3 bucket.
	S3KMSKeyArn *string `json:"s3KmsKeyArn,omitempty" tf:"s3_kms_key_arn,omitempty"`

	// The prefix for the specified S3 bucket.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties *SnapshotDeliveryPropertiesObservation `json:"snapshotDeliveryProperties,omitempty" tf:"snapshot_delivery_properties,omitempty"`

	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`
}

type DeliveryChannelParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the S3 bucket used to store the configuration history.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +kubebuilder:validation:Optional
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameRef *v1.Reference `json:"s3BucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3BucketName.
	// +kubebuilder:validation:Optional
	S3BucketNameSelector *v1.Selector `json:"s3BucketNameSelector,omitempty" tf:"-"`

	// The ARN of the AWS KMS key used to encrypt objects delivered by AWS Config. Must belong to the same Region as the destination S3 bucket.
	// +kubebuilder:validation:Optional
	S3KMSKeyArn *string `json:"s3KmsKeyArn,omitempty" tf:"s3_kms_key_arn,omitempty"`

	// The prefix for the specified S3 bucket.
	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// Options for how AWS Config delivers configuration snapshots. See below
	// +kubebuilder:validation:Optional
	SnapshotDeliveryProperties *SnapshotDeliveryPropertiesParameters `json:"snapshotDeliveryProperties,omitempty" tf:"snapshot_delivery_properties,omitempty"`

	// The ARN of the SNS topic that AWS Config delivers notifications to.
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`
}

type SnapshotDeliveryPropertiesInitParameters struct {

	// - The frequency with which AWS Config recurringly delivers configuration snapshotsE.g., One_Hour or Three_Hours. Valid values are listed here.
	DeliveryFrequency *string `json:"deliveryFrequency,omitempty" tf:"delivery_frequency,omitempty"`
}

type SnapshotDeliveryPropertiesObservation struct {

	// - The frequency with which AWS Config recurringly delivers configuration snapshotsE.g., One_Hour or Three_Hours. Valid values are listed here.
	DeliveryFrequency *string `json:"deliveryFrequency,omitempty" tf:"delivery_frequency,omitempty"`
}

type SnapshotDeliveryPropertiesParameters struct {

	// - The frequency with which AWS Config recurringly delivers configuration snapshotsE.g., One_Hour or Three_Hours. Valid values are listed here.
	// +kubebuilder:validation:Optional
	DeliveryFrequency *string `json:"deliveryFrequency,omitempty" tf:"delivery_frequency,omitempty"`
}

// DeliveryChannelSpec defines the desired state of DeliveryChannel
type DeliveryChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeliveryChannelParameters `json:"forProvider"`
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
	InitProvider DeliveryChannelInitParameters `json:"initProvider,omitempty"`
}

// DeliveryChannelStatus defines the observed state of DeliveryChannel.
type DeliveryChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeliveryChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DeliveryChannel is the Schema for the DeliveryChannels API. Provides an AWS Config Delivery Channel.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DeliveryChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeliveryChannelSpec   `json:"spec"`
	Status            DeliveryChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeliveryChannelList contains a list of DeliveryChannels
type DeliveryChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeliveryChannel `json:"items"`
}

// Repository type metadata.
var (
	DeliveryChannel_Kind             = "DeliveryChannel"
	DeliveryChannel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeliveryChannel_Kind}.String()
	DeliveryChannel_KindAPIVersion   = DeliveryChannel_Kind + "." + CRDGroupVersion.String()
	DeliveryChannel_GroupVersionKind = CRDGroupVersion.WithKind(DeliveryChannel_Kind)
)

func init() {
	SchemeBuilder.Register(&DeliveryChannel{}, &DeliveryChannelList{})
}
