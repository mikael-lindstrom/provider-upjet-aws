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

type DestinationConfigurationInitParameters struct {

	// S3 destination configuration where recorded videos will be stored.
	S3 *S3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type DestinationConfigurationObservation struct {

	// S3 destination configuration where recorded videos will be stored.
	S3 *S3Observation `json:"s3,omitempty" tf:"s3,omitempty"`
}

type DestinationConfigurationParameters struct {

	// S3 destination configuration where recorded videos will be stored.
	// +kubebuilder:validation:Optional
	S3 *S3Parameters `json:"s3" tf:"s3,omitempty"`
}

type RecordingConfigurationInitParameters struct {

	// Object containing destination configuration for where recorded video will be stored.
	DestinationConfiguration *DestinationConfigurationInitParameters `json:"destinationConfiguration,omitempty" tf:"destination_configuration,omitempty"`

	// Recording Configuration name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	RecordingReconnectWindowSeconds *float64 `json:"recordingReconnectWindowSeconds,omitempty" tf:"recording_reconnect_window_seconds,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
	ThumbnailConfiguration *ThumbnailConfigurationInitParameters `json:"thumbnailConfiguration,omitempty" tf:"thumbnail_configuration,omitempty"`
}

type RecordingConfigurationObservation struct {

	// ARN of the Recording Configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Object containing destination configuration for where recorded video will be stored.
	DestinationConfiguration *DestinationConfigurationObservation `json:"destinationConfiguration,omitempty" tf:"destination_configuration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Recording Configuration name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	RecordingReconnectWindowSeconds *float64 `json:"recordingReconnectWindowSeconds,omitempty" tf:"recording_reconnect_window_seconds,omitempty"`

	// The current state of the Recording Configuration.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
	ThumbnailConfiguration *ThumbnailConfigurationObservation `json:"thumbnailConfiguration,omitempty" tf:"thumbnail_configuration,omitempty"`
}

type RecordingConfigurationParameters struct {

	// Object containing destination configuration for where recorded video will be stored.
	// +kubebuilder:validation:Optional
	DestinationConfiguration *DestinationConfigurationParameters `json:"destinationConfiguration,omitempty" tf:"destination_configuration,omitempty"`

	// Recording Configuration name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	// +kubebuilder:validation:Optional
	RecordingReconnectWindowSeconds *float64 `json:"recordingReconnectWindowSeconds,omitempty" tf:"recording_reconnect_window_seconds,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
	// +kubebuilder:validation:Optional
	ThumbnailConfiguration *ThumbnailConfigurationParameters `json:"thumbnailConfiguration,omitempty" tf:"thumbnail_configuration,omitempty"`
}

type S3InitParameters struct {

	// S3 bucket name where recorded videos will be stored.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`
}

type S3Observation struct {

	// S3 bucket name where recorded videos will be stored.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`
}

type S3Parameters struct {

	// S3 bucket name where recorded videos will be stored.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`
}

type ThumbnailConfigurationInitParameters struct {

	// Thumbnail recording mode. Valid values: DISABLED, INTERVAL.
	RecordingMode *string `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// The targeted thumbnail-generation interval in seconds.
	TargetIntervalSeconds *float64 `json:"targetIntervalSeconds,omitempty" tf:"target_interval_seconds,omitempty"`
}

type ThumbnailConfigurationObservation struct {

	// Thumbnail recording mode. Valid values: DISABLED, INTERVAL.
	RecordingMode *string `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// The targeted thumbnail-generation interval in seconds.
	TargetIntervalSeconds *float64 `json:"targetIntervalSeconds,omitempty" tf:"target_interval_seconds,omitempty"`
}

type ThumbnailConfigurationParameters struct {

	// Thumbnail recording mode. Valid values: DISABLED, INTERVAL.
	// +kubebuilder:validation:Optional
	RecordingMode *string `json:"recordingMode,omitempty" tf:"recording_mode,omitempty"`

	// The targeted thumbnail-generation interval in seconds.
	// +kubebuilder:validation:Optional
	TargetIntervalSeconds *float64 `json:"targetIntervalSeconds,omitempty" tf:"target_interval_seconds,omitempty"`
}

// RecordingConfigurationSpec defines the desired state of RecordingConfiguration
type RecordingConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordingConfigurationParameters `json:"forProvider"`
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
	InitProvider RecordingConfigurationInitParameters `json:"initProvider,omitempty"`
}

// RecordingConfigurationStatus defines the observed state of RecordingConfiguration.
type RecordingConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RecordingConfiguration is the Schema for the RecordingConfigurations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RecordingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationConfiguration) || (has(self.initProvider) && has(self.initProvider.destinationConfiguration))",message="spec.forProvider.destinationConfiguration is a required parameter"
	Spec   RecordingConfigurationSpec   `json:"spec"`
	Status RecordingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordingConfigurationList contains a list of RecordingConfigurations
type RecordingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecordingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	RecordingConfiguration_Kind             = "RecordingConfiguration"
	RecordingConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecordingConfiguration_Kind}.String()
	RecordingConfiguration_KindAPIVersion   = RecordingConfiguration_Kind + "." + CRDGroupVersion.String()
	RecordingConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(RecordingConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&RecordingConfiguration{}, &RecordingConfigurationList{})
}
