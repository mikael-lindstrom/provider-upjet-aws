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

type HealthCheckInitParameters struct {

	// Whether health checks are enabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of consecutive health check successes required before considering a target healthy. The range is 2-10. Defaults to 3.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// Approximate amount of time, in seconds, between health checks of an individual target. The range is 5-300. For lambda target groups, it needs to be greater than the timeout of the underlying lambda. Defaults to 30.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// 299" or "0-99"). Required for HTTP/HTTPS/GRPC ALB. Only applies to Application Load Balancers (i.e., HTTP/HTTPS/GRPC) not Network Load Balancers (i.e., TCP).
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (May be required) Destination for the health check request. Required for HTTP/HTTPS ALB and HTTP NLB. Only applies to HTTP/HTTPS.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port the load balancer uses when performing health checks on targets. Default is traffic-port.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol the load balancer uses when performing health checks on targets. Must be either TCP, HTTP, or HTTPS. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Amount of time, in seconds, during which no response from a target means a failed health check. The range is 2–120 seconds. For target groups with a protocol of HTTP, the default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If the target type is lambda, the default is 30 seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Number of consecutive health check failures required before considering a target unhealthy. The range is 2-10. Defaults to 3.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckObservation struct {

	// Whether health checks are enabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of consecutive health check successes required before considering a target healthy. The range is 2-10. Defaults to 3.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// Approximate amount of time, in seconds, between health checks of an individual target. The range is 5-300. For lambda target groups, it needs to be greater than the timeout of the underlying lambda. Defaults to 30.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// 299" or "0-99"). Required for HTTP/HTTPS/GRPC ALB. Only applies to Application Load Balancers (i.e., HTTP/HTTPS/GRPC) not Network Load Balancers (i.e., TCP).
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (May be required) Destination for the health check request. Required for HTTP/HTTPS ALB and HTTP NLB. Only applies to HTTP/HTTPS.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port the load balancer uses when performing health checks on targets. Default is traffic-port.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol the load balancer uses when performing health checks on targets. Must be either TCP, HTTP, or HTTPS. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Amount of time, in seconds, during which no response from a target means a failed health check. The range is 2–120 seconds. For target groups with a protocol of HTTP, the default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If the target type is lambda, the default is 30 seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Number of consecutive health check failures required before considering a target unhealthy. The range is 2-10. Defaults to 3.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckParameters struct {

	// Whether health checks are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Number of consecutive health check successes required before considering a target healthy. The range is 2-10. Defaults to 3.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// Approximate amount of time, in seconds, between health checks of an individual target. The range is 5-300. For lambda target groups, it needs to be greater than the timeout of the underlying lambda. Defaults to 30.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// 299" or "0-99"). Required for HTTP/HTTPS/GRPC ALB. Only applies to Application Load Balancers (i.e., HTTP/HTTPS/GRPC) not Network Load Balancers (i.e., TCP).
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// (May be required) Destination for the health check request. Required for HTTP/HTTPS ALB and HTTP NLB. Only applies to HTTP/HTTPS.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port the load balancer uses when performing health checks on targets. Default is traffic-port.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol the load balancer uses when performing health checks on targets. Must be either TCP, HTTP, or HTTPS. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Amount of time, in seconds, during which no response from a target means a failed health check. The range is 2–120 seconds. For target groups with a protocol of HTTP, the default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If the target type is lambda, the default is 30 seconds.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Number of consecutive health check failures required before considering a target unhealthy. The range is 2-10. Defaults to 3.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LBTargetGroupInitParameters struct {

	// Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See doc for more information. Default is false.
	ConnectionTermination *bool `json:"connectionTermination,omitempty" tf:"connection_termination,omitempty"`

	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay *string `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`

	// Health Check configuration block. Detailed below.
	HealthCheck []HealthCheckInitParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The type of IP addresses used by the target group, only supported when target type is set to ip. Possible values are ipv4 or ipv6.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when target_type is lambda. Default is false.
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`

	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is round_robin or least_outstanding_requests. The default is round_robin.
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type,omitempty"`

	// Indicates whether cross zone load balancing is enabled. The value is "true", "false" or "use_load_balancer_configuration". The default is "use_load_balancer_configuration".
	LoadBalancingCrossZoneEnabled *string `json:"loadBalancingCrossZoneEnabled,omitempty" tf:"load_balancing_cross_zone_enabled,omitempty"`

	// Name of the target group. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (May be required, Forces new resource) Port on which targets receive traffic, unless overridden when registering a specific target. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether client IP preservation is enabled. See doc for more information.
	PreserveClientIP *string `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// (May be required, Forces new resource) Protocol to use for routing traffic to the targets. Should be one of GENEVE, HTTP, HTTPS, TCP, TCP_UDP, TLS, or UDP. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Only applicable when protocol is HTTP or HTTPS. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See doc for more information. Default is false.
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`

	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart *float64 `json:"slowStart,omitempty" tf:"slow_start,omitempty"`

	// Stickiness configuration block. Detailed below.
	Stickiness []LBTargetGroupStickinessInitParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
	TargetFailover []TargetFailoverInitParameters `json:"targetFailover,omitempty" tf:"target_failover,omitempty"`

	// (May be required, Forces new resource) Type of target that you must specify when registering targets with this target group. See doc for supported values. The default is instance.
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`
}

type LBTargetGroupObservation struct {

	// ARN of the Target Group (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN suffix for use with CloudWatch Metrics.
	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	// Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See doc for more information. Default is false.
	ConnectionTermination *bool `json:"connectionTermination,omitempty" tf:"connection_termination,omitempty"`

	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay *string `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`

	// Health Check configuration block. Detailed below.
	HealthCheck []HealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// ARN of the Target Group (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of IP addresses used by the target group, only supported when target type is set to ip. Possible values are ipv4 or ipv6.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when target_type is lambda. Default is false.
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`

	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is round_robin or least_outstanding_requests. The default is round_robin.
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type,omitempty"`

	// Indicates whether cross zone load balancing is enabled. The value is "true", "false" or "use_load_balancer_configuration". The default is "use_load_balancer_configuration".
	LoadBalancingCrossZoneEnabled *string `json:"loadBalancingCrossZoneEnabled,omitempty" tf:"load_balancing_cross_zone_enabled,omitempty"`

	// Name of the target group. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (May be required, Forces new resource) Port on which targets receive traffic, unless overridden when registering a specific target. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether client IP preservation is enabled. See doc for more information.
	PreserveClientIP *string `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// (May be required, Forces new resource) Protocol to use for routing traffic to the targets. Should be one of GENEVE, HTTP, HTTPS, TCP, TCP_UDP, TLS, or UDP. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Only applicable when protocol is HTTP or HTTPS. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See doc for more information. Default is false.
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`

	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart *float64 `json:"slowStart,omitempty" tf:"slow_start,omitempty"`

	// Stickiness configuration block. Detailed below.
	Stickiness []LBTargetGroupStickinessObservation `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
	TargetFailover []TargetFailoverObservation `json:"targetFailover,omitempty" tf:"target_failover,omitempty"`

	// (May be required, Forces new resource) Type of target that you must specify when registering targets with this target group. See doc for supported values. The default is instance.
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// Identifier of the VPC in which to create the target group. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type LBTargetGroupParameters struct {

	// Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See doc for more information. Default is false.
	// +kubebuilder:validation:Optional
	ConnectionTermination *bool `json:"connectionTermination,omitempty" tf:"connection_termination,omitempty"`

	// Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	// +kubebuilder:validation:Optional
	DeregistrationDelay *string `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`

	// Health Check configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// The type of IP addresses used by the target group, only supported when target type is set to ip. Possible values are ipv4 or ipv6.
	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when target_type is lambda. Default is false.
	// +kubebuilder:validation:Optional
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`

	// Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is round_robin or least_outstanding_requests. The default is round_robin.
	// +kubebuilder:validation:Optional
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type,omitempty"`

	// Indicates whether cross zone load balancing is enabled. The value is "true", "false" or "use_load_balancer_configuration". The default is "use_load_balancer_configuration".
	// +kubebuilder:validation:Optional
	LoadBalancingCrossZoneEnabled *string `json:"loadBalancingCrossZoneEnabled,omitempty" tf:"load_balancing_cross_zone_enabled,omitempty"`

	// Name of the target group. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (May be required, Forces new resource) Port on which targets receive traffic, unless overridden when registering a specific target. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Whether client IP preservation is enabled. See doc for more information.
	// +kubebuilder:validation:Optional
	PreserveClientIP *string `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// (May be required, Forces new resource) Protocol to use for routing traffic to the targets. Should be one of GENEVE, HTTP, HTTPS, TCP, TCP_UDP, TLS, or UDP. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Only applicable when protocol is HTTP or HTTPS. The protocol version. Specify GRPC to send requests to targets using gRPC. Specify HTTP2 to send requests to targets using HTTP/2. The default is HTTP1, which sends requests to targets using HTTP/1.1
	// +kubebuilder:validation:Optional
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// Whether to enable support for proxy protocol v2 on Network Load Balancers. See doc for more information. Default is false.
	// +kubebuilder:validation:Optional
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	// +kubebuilder:validation:Optional
	SlowStart *float64 `json:"slowStart,omitempty" tf:"slow_start,omitempty"`

	// Stickiness configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	Stickiness []LBTargetGroupStickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
	// +kubebuilder:validation:Optional
	TargetFailover []TargetFailoverParameters `json:"targetFailover,omitempty" tf:"target_failover,omitempty"`

	// (May be required, Forces new resource) Type of target that you must specify when registering targets with this target group. See doc for supported values. The default is instance.
	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// Identifier of the VPC in which to create the target group. Required when target_type is instance, ip or alb. Does not apply when target_type is lambda.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type LBTargetGroupStickinessInitParameters struct {

	// Only used when the type is lb_cookie. The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// Name of the application based cookie. AWSALB, AWSALBAPP, and AWSALBTG prefixes are reserved and cannot be used. Only needed when type is app_cookie.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Whether health checks are enabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The type of sticky sessions. The only current possible values are lb_cookie, app_cookie for ALBs, source_ip for NLBs, and source_ip_dest_ip, source_ip_dest_ip_proto for GWLBs.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LBTargetGroupStickinessObservation struct {

	// Only used when the type is lb_cookie. The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// Name of the application based cookie. AWSALB, AWSALBAPP, and AWSALBTG prefixes are reserved and cannot be used. Only needed when type is app_cookie.
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Whether health checks are enabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The type of sticky sessions. The only current possible values are lb_cookie, app_cookie for ALBs, source_ip for NLBs, and source_ip_dest_ip, source_ip_dest_ip_proto for GWLBs.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LBTargetGroupStickinessParameters struct {

	// Only used when the type is lb_cookie. The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	// +kubebuilder:validation:Optional
	CookieDuration *float64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// Name of the application based cookie. AWSALB, AWSALBAPP, and AWSALBTG prefixes are reserved and cannot be used. Only needed when type is app_cookie.
	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// Whether health checks are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The type of sticky sessions. The only current possible values are lb_cookie, app_cookie for ALBs, source_ip for NLBs, and source_ip_dest_ip, source_ip_dest_ip_proto for GWLBs.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type TargetFailoverInitParameters struct {

	// Indicates how the GWLB handles existing flows when a target is deregistered. Possible values are rebalance and no_rebalance. Must match the attribute value set for on_unhealthy. Default: no_rebalance.
	OnDeregistration *string `json:"onDeregistration,omitempty" tf:"on_deregistration,omitempty"`

	// Indicates how the GWLB handles existing flows when a target is unhealthy. Possible values are rebalance and no_rebalance. Must match the attribute value set for on_deregistration. Default: no_rebalance.
	OnUnhealthy *string `json:"onUnhealthy,omitempty" tf:"on_unhealthy,omitempty"`
}

type TargetFailoverObservation struct {

	// Indicates how the GWLB handles existing flows when a target is deregistered. Possible values are rebalance and no_rebalance. Must match the attribute value set for on_unhealthy. Default: no_rebalance.
	OnDeregistration *string `json:"onDeregistration,omitempty" tf:"on_deregistration,omitempty"`

	// Indicates how the GWLB handles existing flows when a target is unhealthy. Possible values are rebalance and no_rebalance. Must match the attribute value set for on_deregistration. Default: no_rebalance.
	OnUnhealthy *string `json:"onUnhealthy,omitempty" tf:"on_unhealthy,omitempty"`
}

type TargetFailoverParameters struct {

	// Indicates how the GWLB handles existing flows when a target is deregistered. Possible values are rebalance and no_rebalance. Must match the attribute value set for on_unhealthy. Default: no_rebalance.
	// +kubebuilder:validation:Optional
	OnDeregistration *string `json:"onDeregistration" tf:"on_deregistration,omitempty"`

	// Indicates how the GWLB handles existing flows when a target is unhealthy. Possible values are rebalance and no_rebalance. Must match the attribute value set for on_deregistration. Default: no_rebalance.
	// +kubebuilder:validation:Optional
	OnUnhealthy *string `json:"onUnhealthy" tf:"on_unhealthy,omitempty"`
}

// LBTargetGroupSpec defines the desired state of LBTargetGroup
type LBTargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBTargetGroupParameters `json:"forProvider"`
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
	InitProvider LBTargetGroupInitParameters `json:"initProvider,omitempty"`
}

// LBTargetGroupStatus defines the observed state of LBTargetGroup.
type LBTargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBTargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBTargetGroup is the Schema for the LBTargetGroups API. Provides a Target Group resource for use with Load Balancers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LBTargetGroupSpec   `json:"spec"`
	Status LBTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBTargetGroupList contains a list of LBTargetGroups
type LBTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBTargetGroup `json:"items"`
}

// Repository type metadata.
var (
	LBTargetGroup_Kind             = "LBTargetGroup"
	LBTargetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBTargetGroup_Kind}.String()
	LBTargetGroup_KindAPIVersion   = LBTargetGroup_Kind + "." + CRDGroupVersion.String()
	LBTargetGroup_GroupVersionKind = CRDGroupVersion.WithKind(LBTargetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&LBTargetGroup{}, &LBTargetGroupList{})
}
