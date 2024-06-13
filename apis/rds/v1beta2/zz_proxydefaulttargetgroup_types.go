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

type ConnectionPoolConfigInitParameters struct {

	// The number of seconds for a proxy to wait for a connection to become available in the connection pool. Only applies when the proxy has opened its maximum number of connections and all connections are busy with client sessions.
	ConnectionBorrowTimeout *float64 `json:"connectionBorrowTimeout,omitempty" tf:"connection_borrow_timeout,omitempty"`

	// One or more SQL statements for the proxy to run when opening each new database connection. Typically used with SET statements to make sure that each connection has identical settings such as time zone and character set. This setting is empty by default. For multiple statements, use semicolons as the separator. You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	InitQuery *string `json:"initQuery,omitempty" tf:"init_query,omitempty"`

	// The maximum size of the connection pool for each target in a target group. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent,omitempty" tf:"max_connections_percent,omitempty"`

	// Controls how actively the proxy closes idle database connections in the connection pool. A high value enables the proxy to leave a high percentage of idle connections open. A low value causes the proxy to close idle client connections and return the underlying database connections to the connection pool. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent,omitempty" tf:"max_idle_connections_percent,omitempty"`

	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection. Including an item in the list exempts that class of SQL operations from the pinning behavior. Currently, the only allowed value is EXCLUDE_VARIABLE_SETS.
	// +listType=set
	SessionPinningFilters []*string `json:"sessionPinningFilters,omitempty" tf:"session_pinning_filters,omitempty"`
}

type ConnectionPoolConfigObservation struct {

	// The number of seconds for a proxy to wait for a connection to become available in the connection pool. Only applies when the proxy has opened its maximum number of connections and all connections are busy with client sessions.
	ConnectionBorrowTimeout *float64 `json:"connectionBorrowTimeout,omitempty" tf:"connection_borrow_timeout,omitempty"`

	// One or more SQL statements for the proxy to run when opening each new database connection. Typically used with SET statements to make sure that each connection has identical settings such as time zone and character set. This setting is empty by default. For multiple statements, use semicolons as the separator. You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	InitQuery *string `json:"initQuery,omitempty" tf:"init_query,omitempty"`

	// The maximum size of the connection pool for each target in a target group. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent,omitempty" tf:"max_connections_percent,omitempty"`

	// Controls how actively the proxy closes idle database connections in the connection pool. A high value enables the proxy to leave a high percentage of idle connections open. A low value causes the proxy to close idle client connections and return the underlying database connections to the connection pool. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent,omitempty" tf:"max_idle_connections_percent,omitempty"`

	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection. Including an item in the list exempts that class of SQL operations from the pinning behavior. Currently, the only allowed value is EXCLUDE_VARIABLE_SETS.
	// +listType=set
	SessionPinningFilters []*string `json:"sessionPinningFilters,omitempty" tf:"session_pinning_filters,omitempty"`
}

type ConnectionPoolConfigParameters struct {

	// The number of seconds for a proxy to wait for a connection to become available in the connection pool. Only applies when the proxy has opened its maximum number of connections and all connections are busy with client sessions.
	// +kubebuilder:validation:Optional
	ConnectionBorrowTimeout *float64 `json:"connectionBorrowTimeout,omitempty" tf:"connection_borrow_timeout,omitempty"`

	// One or more SQL statements for the proxy to run when opening each new database connection. Typically used with SET statements to make sure that each connection has identical settings such as time zone and character set. This setting is empty by default. For multiple statements, use semicolons as the separator. You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	// +kubebuilder:validation:Optional
	InitQuery *string `json:"initQuery,omitempty" tf:"init_query,omitempty"`

	// The maximum size of the connection pool for each target in a target group. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	// +kubebuilder:validation:Optional
	MaxConnectionsPercent *float64 `json:"maxConnectionsPercent,omitempty" tf:"max_connections_percent,omitempty"`

	// Controls how actively the proxy closes idle database connections in the connection pool. A high value enables the proxy to leave a high percentage of idle connections open. A low value causes the proxy to close idle client connections and return the underlying database connections to the connection pool. For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB cluster used by the target group.
	// +kubebuilder:validation:Optional
	MaxIdleConnectionsPercent *float64 `json:"maxIdleConnectionsPercent,omitempty" tf:"max_idle_connections_percent,omitempty"`

	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection. Including an item in the list exempts that class of SQL operations from the pinning behavior. Currently, the only allowed value is EXCLUDE_VARIABLE_SETS.
	// +kubebuilder:validation:Optional
	// +listType=set
	SessionPinningFilters []*string `json:"sessionPinningFilters,omitempty" tf:"session_pinning_filters,omitempty"`
}

type ProxyDefaultTargetGroupInitParameters struct {

	// The settings that determine the size and behavior of the connection pool for the target group.
	ConnectionPoolConfig *ConnectionPoolConfigInitParameters `json:"connectionPoolConfig,omitempty" tf:"connection_pool_config,omitempty"`

	// Name of the RDS DB Proxy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Proxy
	DBProxyName *string `json:"dbProxyName,omitempty" tf:"db_proxy_name,omitempty"`

	// Reference to a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameRef *v1.Reference `json:"dbProxyNameRef,omitempty" tf:"-"`

	// Selector for a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameSelector *v1.Selector `json:"dbProxyNameSelector,omitempty" tf:"-"`
}

type ProxyDefaultTargetGroupObservation struct {

	// The Amazon Resource Name (ARN) representing the target group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The settings that determine the size and behavior of the connection pool for the target group.
	ConnectionPoolConfig *ConnectionPoolConfigObservation `json:"connectionPoolConfig,omitempty" tf:"connection_pool_config,omitempty"`

	// Name of the RDS DB Proxy.
	DBProxyName *string `json:"dbProxyName,omitempty" tf:"db_proxy_name,omitempty"`

	// Name of the RDS DB Proxy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the default target group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProxyDefaultTargetGroupParameters struct {

	// The settings that determine the size and behavior of the connection pool for the target group.
	// +kubebuilder:validation:Optional
	ConnectionPoolConfig *ConnectionPoolConfigParameters `json:"connectionPoolConfig,omitempty" tf:"connection_pool_config,omitempty"`

	// Name of the RDS DB Proxy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Proxy
	// +kubebuilder:validation:Optional
	DBProxyName *string `json:"dbProxyName,omitempty" tf:"db_proxy_name,omitempty"`

	// Reference to a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameRef *v1.Reference `json:"dbProxyNameRef,omitempty" tf:"-"`

	// Selector for a Proxy in rds to populate dbProxyName.
	// +kubebuilder:validation:Optional
	DBProxyNameSelector *v1.Selector `json:"dbProxyNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ProxyDefaultTargetGroupSpec defines the desired state of ProxyDefaultTargetGroup
type ProxyDefaultTargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyDefaultTargetGroupParameters `json:"forProvider"`
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
	InitProvider ProxyDefaultTargetGroupInitParameters `json:"initProvider,omitempty"`
}

// ProxyDefaultTargetGroupStatus defines the observed state of ProxyDefaultTargetGroup.
type ProxyDefaultTargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyDefaultTargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProxyDefaultTargetGroup is the Schema for the ProxyDefaultTargetGroups API. Manage an RDS DB proxy default target group resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ProxyDefaultTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxyDefaultTargetGroupSpec   `json:"spec"`
	Status            ProxyDefaultTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyDefaultTargetGroupList contains a list of ProxyDefaultTargetGroups
type ProxyDefaultTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyDefaultTargetGroup `json:"items"`
}

// Repository type metadata.
var (
	ProxyDefaultTargetGroup_Kind             = "ProxyDefaultTargetGroup"
	ProxyDefaultTargetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProxyDefaultTargetGroup_Kind}.String()
	ProxyDefaultTargetGroup_KindAPIVersion   = ProxyDefaultTargetGroup_Kind + "." + CRDGroupVersion.String()
	ProxyDefaultTargetGroup_GroupVersionKind = CRDGroupVersion.WithKind(ProxyDefaultTargetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ProxyDefaultTargetGroup{}, &ProxyDefaultTargetGroupList{})
}
