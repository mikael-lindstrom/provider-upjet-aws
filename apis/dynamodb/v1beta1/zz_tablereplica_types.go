// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableReplicaInitParameters_2 struct {

	// ARN of the main or global table which this resource will replicate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dynamodb/v1beta2.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	GlobalTableArn *string `json:"globalTableArn,omitempty" tf:"global_table_arn,omitempty"`

	// Reference to a Table in dynamodb to populate globalTableArn.
	// +kubebuilder:validation:Optional
	GlobalTableArnRef *v1.Reference `json:"globalTableArnRef,omitempty" tf:"-"`

	// Selector for a Table in dynamodb to populate globalTableArn.
	// +kubebuilder:validation:Optional
	GlobalTableArnSelector *v1.Selector `json:"globalTableArnSelector,omitempty" tf:"-"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Whether to enable Point In Time Recovery for the replica. Default is false.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// Storage class of the table replica. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS. If not used, the table replica will use the same class as the global table.
	TableClassOverride *string `json:"tableClassOverride,omitempty" tf:"table_class_override,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TableReplicaObservation_2 struct {

	// ARN of the table replica.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the main or global table which this resource will replicate.
	GlobalTableArn *string `json:"globalTableArn,omitempty" tf:"global_table_arn,omitempty"`

	// Name of the table and region of the main global table joined with a semicolon (e.g., TableName:us-east-1).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Whether to enable Point In Time Recovery for the replica. Default is false.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// Storage class of the table replica. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS. If not used, the table replica will use the same class as the global table.
	TableClassOverride *string `json:"tableClassOverride,omitempty" tf:"table_class_override,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TableReplicaParameters_2 struct {

	// ARN of the main or global table which this resource will replicate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dynamodb/v1beta2.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	GlobalTableArn *string `json:"globalTableArn,omitempty" tf:"global_table_arn,omitempty"`

	// Reference to a Table in dynamodb to populate globalTableArn.
	// +kubebuilder:validation:Optional
	GlobalTableArnRef *v1.Reference `json:"globalTableArnRef,omitempty" tf:"-"`

	// Selector for a Table in dynamodb to populate globalTableArn.
	// +kubebuilder:validation:Optional
	GlobalTableArnSelector *v1.Selector `json:"globalTableArnSelector,omitempty" tf:"-"`

	// ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, alias/aws/dynamodb. Note: This attribute will not be populated with the ARN of default keys.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Whether to enable Point In Time Recovery for the replica. Default is false.
	// +kubebuilder:validation:Optional
	PointInTimeRecovery *bool `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Storage class of the table replica. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS. If not used, the table replica will use the same class as the global table.
	// +kubebuilder:validation:Optional
	TableClassOverride *string `json:"tableClassOverride,omitempty" tf:"table_class_override,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TableReplicaSpec defines the desired state of TableReplica
type TableReplicaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableReplicaParameters_2 `json:"forProvider"`
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
	InitProvider TableReplicaInitParameters_2 `json:"initProvider,omitempty"`
}

// TableReplicaStatus defines the observed state of TableReplica.
type TableReplicaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableReplicaObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TableReplica is the Schema for the TableReplicas API. Provides a DynamoDB table replica resource
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TableReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableReplicaSpec   `json:"spec"`
	Status            TableReplicaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableReplicaList contains a list of TableReplicas
type TableReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableReplica `json:"items"`
}

// Repository type metadata.
var (
	TableReplica_Kind             = "TableReplica"
	TableReplica_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableReplica_Kind}.String()
	TableReplica_KindAPIVersion   = TableReplica_Kind + "." + CRDGroupVersion.String()
	TableReplica_GroupVersionKind = CRDGroupVersion.WithKind(TableReplica_Kind)
)

func init() {
	SchemeBuilder.Register(&TableReplica{}, &TableReplicaList{})
}
