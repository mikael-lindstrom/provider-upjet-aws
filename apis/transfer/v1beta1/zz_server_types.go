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

type EndpointDetailsInitParameters struct {

	// A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when endpoint_type is set to VPC.
	AddressAllocationIds []*string `json:"addressAllocationIds,omitempty" tf:"address_allocation_ids,omitempty"`

	// A list of security groups IDs that are available to attach to your server's endpoint. If no security groups are specified, the VPC's default security groups are automatically assigned to your endpoint. This property can only be used when endpoint_type is set to VPC.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when endpoint_type is set to VPC.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The ID of the VPC endpoint. This property can only be used when endpoint_type is set to VPC_ENDPOINT
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`
}

type EndpointDetailsObservation struct {

	// A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when endpoint_type is set to VPC.
	AddressAllocationIds []*string `json:"addressAllocationIds,omitempty" tf:"address_allocation_ids,omitempty"`

	// A list of security groups IDs that are available to attach to your server's endpoint. If no security groups are specified, the VPC's default security groups are automatically assigned to your endpoint. This property can only be used when endpoint_type is set to VPC.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when endpoint_type is set to VPC.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The ID of the VPC endpoint. This property can only be used when endpoint_type is set to VPC_ENDPOINT
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when endpoint_type is set to VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type EndpointDetailsParameters struct {

	// A list of address allocation IDs that are required to attach an Elastic IP address to your SFTP server's endpoint. This property can only be used when endpoint_type is set to VPC.
	// +kubebuilder:validation:Optional
	AddressAllocationIds []*string `json:"addressAllocationIds,omitempty" tf:"address_allocation_ids,omitempty"`

	// A list of security groups IDs that are available to attach to your server's endpoint. If no security groups are specified, the VPC's default security groups are automatically assigned to your endpoint. This property can only be used when endpoint_type is set to VPC.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// A list of subnet IDs that are required to host your SFTP server endpoint in your VPC. This property can only be used when endpoint_type is set to VPC.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The ID of the VPC endpoint. This property can only be used when endpoint_type is set to VPC_ENDPOINT
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The VPC ID of the virtual private cloud in which the SFTP server's endpoint will be hosted. This property can only be used when endpoint_type is set to VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type OnPartialUploadInitParameters struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole *string `json:"executionRole,omitempty" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	WorkflowID *string `json:"workflowId,omitempty" tf:"workflow_id,omitempty"`
}

type OnPartialUploadObservation struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole *string `json:"executionRole,omitempty" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	WorkflowID *string `json:"workflowId,omitempty" tf:"workflow_id,omitempty"`
}

type OnPartialUploadParameters struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	// +kubebuilder:validation:Optional
	ExecutionRole *string `json:"executionRole" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	// +kubebuilder:validation:Optional
	WorkflowID *string `json:"workflowId" tf:"workflow_id,omitempty"`
}

type OnUploadInitParameters struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole *string `json:"executionRole,omitempty" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	WorkflowID *string `json:"workflowId,omitempty" tf:"workflow_id,omitempty"`
}

type OnUploadObservation struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole *string `json:"executionRole,omitempty" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	WorkflowID *string `json:"workflowId,omitempty" tf:"workflow_id,omitempty"`
}

type OnUploadParameters struct {

	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	// +kubebuilder:validation:Optional
	ExecutionRole *string `json:"executionRole" tf:"execution_role,omitempty"`

	// A unique identifier for the workflow.
	// +kubebuilder:validation:Optional
	WorkflowID *string `json:"workflowId" tf:"workflow_id,omitempty"`
}

type ProtocolDetailsInitParameters struct {

	// Indicates the transport method for the AS2 messages. Currently, only HTTP is supported.
	As2Transports []*string `json:"as2Transports,omitempty" tf:"as2_transports,omitempty"`

	// Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
	PassiveIP *string `json:"passiveIp,omitempty" tf:"passive_ip,omitempty"`

	// Use to ignore the error that is generated when the client attempts to use SETSTAT on a file you are uploading to an S3 bucket. Valid values: DEFAULT, ENABLE_NO_OP.
	SetStatOption *string `json:"setStatOption,omitempty" tf:"set_stat_option,omitempty"`

	// A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: DISABLED, ENABLED, ENFORCED.
	TLSSessionResumptionMode *string `json:"tlsSessionResumptionMode,omitempty" tf:"tls_session_resumption_mode,omitempty"`
}

type ProtocolDetailsObservation struct {

	// Indicates the transport method for the AS2 messages. Currently, only HTTP is supported.
	As2Transports []*string `json:"as2Transports,omitempty" tf:"as2_transports,omitempty"`

	// Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
	PassiveIP *string `json:"passiveIp,omitempty" tf:"passive_ip,omitempty"`

	// Use to ignore the error that is generated when the client attempts to use SETSTAT on a file you are uploading to an S3 bucket. Valid values: DEFAULT, ENABLE_NO_OP.
	SetStatOption *string `json:"setStatOption,omitempty" tf:"set_stat_option,omitempty"`

	// A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: DISABLED, ENABLED, ENFORCED.
	TLSSessionResumptionMode *string `json:"tlsSessionResumptionMode,omitempty" tf:"tls_session_resumption_mode,omitempty"`
}

type ProtocolDetailsParameters struct {

	// Indicates the transport method for the AS2 messages. Currently, only HTTP is supported.
	// +kubebuilder:validation:Optional
	As2Transports []*string `json:"as2Transports,omitempty" tf:"as2_transports,omitempty"`

	// Indicates passive mode, for FTP and FTPS protocols. Enter a single IPv4 address, such as the public IP address of a firewall, router, or load balancer.
	// +kubebuilder:validation:Optional
	PassiveIP *string `json:"passiveIp,omitempty" tf:"passive_ip,omitempty"`

	// Use to ignore the error that is generated when the client attempts to use SETSTAT on a file you are uploading to an S3 bucket. Valid values: DEFAULT, ENABLE_NO_OP.
	// +kubebuilder:validation:Optional
	SetStatOption *string `json:"setStatOption,omitempty" tf:"set_stat_option,omitempty"`

	// A property used with Transfer Family servers that use the FTPS protocol. Provides a mechanism to resume or share a negotiated secret key between the control and data connection for an FTPS session. Valid values: DISABLED, ENABLED, ENFORCED.
	// +kubebuilder:validation:Optional
	TLSSessionResumptionMode *string `json:"tlsSessionResumptionMode,omitempty" tf:"tls_session_resumption_mode,omitempty"`
}

type ServerInitParameters struct {

	// The domain of the storage system that is used for file transfers. Valid values are: S3 and EFS. The default value is S3.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
	EndpointDetails []EndpointDetailsInitParameters `json:"endpointDetails,omitempty" tf:"endpoint_details,omitempty"`

	// The type of endpoint that you want your SFTP server connect to. If you connect to a VPC (or VPC_ENDPOINT), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set PUBLIC.  Defaults to PUBLIC.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is false. This option only applies to servers configured with a SERVICE_MANAGED identity_provider_type.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The ARN for a lambda function to use for the Identity provider.
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// The mode of authentication enabled for this service. The default value is SERVICE_MANAGED, which allows you to store and access SFTP user credentials within the service. API_GATEWAY indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using AWS_DIRECTORY_SERVICE will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the AWS_LAMBDA value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the function argument.
	IdentityProviderType *string `json:"identityProviderType,omitempty" tf:"identity_provider_type,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an identity_provider_type of API_GATEWAY.
	InvocationRole *string `json:"invocationRole,omitempty" tf:"invocation_role,omitempty"`

	// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// The protocol settings that are configured for your server.
	ProtocolDetails []ProtocolDetailsInitParameters `json:"protocolDetails,omitempty" tf:"protocol_details,omitempty"`

	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to SFTP . The available protocols are:
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Specifies the name of the security policy that is attached to the server. Possible values are TransferSecurityPolicy-2018-11, TransferSecurityPolicy-2020-06, TransferSecurityPolicy-FIPS-2020-06 and TransferSecurityPolicy-2022-03. Default value is: TransferSecurityPolicy-2018-11.
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// - URL of the service endpoint used to authenticate users with an identity_provider_type of API_GATEWAY.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Specifies the workflow details. See Workflow Details below.
	WorkflowDetails []WorkflowDetailsInitParameters `json:"workflowDetails,omitempty" tf:"workflow_details,omitempty"`
}

type ServerObservation struct {

	// Amazon Resource Name (ARN) of Transfer Server
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when protocols is set to FTPS
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// The directory service ID of the directory service you want to connect to with an identity_provider_type of AWS_DIRECTORY_SERVICE.
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// The domain of the storage system that is used for file transfers. Valid values are: S3 and EFS. The default value is S3.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The endpoint of the Transfer Server (e.g., s-12345678.server.transfer.REGION.amazonaws.com)
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
	EndpointDetails []EndpointDetailsObservation `json:"endpointDetails,omitempty" tf:"endpoint_details,omitempty"`

	// The type of endpoint that you want your SFTP server connect to. If you connect to a VPC (or VPC_ENDPOINT), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set PUBLIC.  Defaults to PUBLIC.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is false. This option only applies to servers configured with a SERVICE_MANAGED identity_provider_type.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The ARN for a lambda function to use for the Identity provider.
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the ssh-keygen -l -E md5 -f my-new-server-key command.
	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty" tf:"host_key_fingerprint,omitempty"`

	// The Server ID of the Transfer Server (e.g., s-12345678)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The mode of authentication enabled for this service. The default value is SERVICE_MANAGED, which allows you to store and access SFTP user credentials within the service. API_GATEWAY indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using AWS_DIRECTORY_SERVICE will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the AWS_LAMBDA value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the function argument.
	IdentityProviderType *string `json:"identityProviderType,omitempty" tf:"identity_provider_type,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an identity_provider_type of API_GATEWAY.
	InvocationRole *string `json:"invocationRole,omitempty" tf:"invocation_role,omitempty"`

	// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// The protocol settings that are configured for your server.
	ProtocolDetails []ProtocolDetailsObservation `json:"protocolDetails,omitempty" tf:"protocol_details,omitempty"`

	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to SFTP . The available protocols are:
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Specifies the name of the security policy that is attached to the server. Possible values are TransferSecurityPolicy-2018-11, TransferSecurityPolicy-2020-06, TransferSecurityPolicy-FIPS-2020-06 and TransferSecurityPolicy-2022-03. Default value is: TransferSecurityPolicy-2018-11.
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// - URL of the service endpoint used to authenticate users with an identity_provider_type of API_GATEWAY.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Specifies the workflow details. See Workflow Details below.
	WorkflowDetails []WorkflowDetailsObservation `json:"workflowDetails,omitempty" tf:"workflow_details,omitempty"`
}

type ServerParameters struct {

	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. This is required when protocols is set to FTPS
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Reference to a Certificate in acm to populate certificate.
	// +kubebuilder:validation:Optional
	CertificateRef *v1.Reference `json:"certificateRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificate.
	// +kubebuilder:validation:Optional
	CertificateSelector *v1.Selector `json:"certificateSelector,omitempty" tf:"-"`

	// The directory service ID of the directory service you want to connect to with an identity_provider_type of AWS_DIRECTORY_SERVICE.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// The domain of the storage system that is used for file transfers. Valid values are: S3 and EFS. The default value is S3.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
	// +kubebuilder:validation:Optional
	EndpointDetails []EndpointDetailsParameters `json:"endpointDetails,omitempty" tf:"endpoint_details,omitempty"`

	// The type of endpoint that you want your SFTP server connect to. If you connect to a VPC (or VPC_ENDPOINT), your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set PUBLIC.  Defaults to PUBLIC.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is false. This option only applies to servers configured with a SERVICE_MANAGED identity_provider_type.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The ARN for a lambda function to use for the Identity provider.
	// +kubebuilder:validation:Optional
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// RSA, ECDSA, or ED25519 private key (e.g., as generated by the ssh-keygen -t rsa -b 2048 -N "" -m PEM -f my-new-server-key, ssh-keygen -t ecdsa -b 256 -N "" -m PEM -f my-new-server-key or ssh-keygen -t ed25519 -N "" -f my-new-server-key commands).
	// +kubebuilder:validation:Optional
	HostKeySecretRef *v1.SecretKeySelector `json:"hostKeySecretRef,omitempty" tf:"-"`

	// The mode of authentication enabled for this service. The default value is SERVICE_MANAGED, which allows you to store and access SFTP user credentials within the service. API_GATEWAY indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice. Using AWS_DIRECTORY_SERVICE will allow for authentication against AWS Managed Active Directory or Microsoft Active Directory in your on-premises environment, or in AWS using AD Connectors. Use the AWS_LAMBDA value to directly use a Lambda function as your identity provider. If you choose this value, you must specify the ARN for the lambda function in the function argument.
	// +kubebuilder:validation:Optional
	IdentityProviderType *string `json:"identityProviderType,omitempty" tf:"identity_provider_type,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an identity_provider_type of API_GATEWAY.
	// +kubebuilder:validation:Optional
	InvocationRole *string `json:"invocationRole,omitempty" tf:"invocation_role,omitempty"`

	// Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
	// +kubebuilder:validation:Optional
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// Specify a string to display when users connect to a server. This string is displayed after the user authenticates. The SFTP protocol does not support post-authentication display banners.
	// +kubebuilder:validation:Optional
	PostAuthenticationLoginBannerSecretRef *v1.SecretKeySelector `json:"postAuthenticationLoginBannerSecretRef,omitempty" tf:"-"`

	// Specify a string to display when users connect to a server. This string is displayed before the user authenticates.
	// +kubebuilder:validation:Optional
	PreAuthenticationLoginBannerSecretRef *v1.SecretKeySelector `json:"preAuthenticationLoginBannerSecretRef,omitempty" tf:"-"`

	// The protocol settings that are configured for your server.
	// +kubebuilder:validation:Optional
	ProtocolDetails []ProtocolDetailsParameters `json:"protocolDetails,omitempty" tf:"protocol_details,omitempty"`

	// Specifies the file transfer protocol or protocols over which your file transfer protocol client can connect to your server's endpoint. This defaults to SFTP . The available protocols are:
	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the name of the security policy that is attached to the server. Possible values are TransferSecurityPolicy-2018-11, TransferSecurityPolicy-2020-06, TransferSecurityPolicy-FIPS-2020-06 and TransferSecurityPolicy-2022-03. Default value is: TransferSecurityPolicy-2018-11.
	// +kubebuilder:validation:Optional
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// - URL of the service endpoint used to authenticate users with an identity_provider_type of API_GATEWAY.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Specifies the workflow details. See Workflow Details below.
	// +kubebuilder:validation:Optional
	WorkflowDetails []WorkflowDetailsParameters `json:"workflowDetails,omitempty" tf:"workflow_details,omitempty"`
}

type WorkflowDetailsInitParameters struct {

	// A trigger that starts a workflow if a file is only partially uploaded. See Workflow Detail below.
	OnPartialUpload []OnPartialUploadInitParameters `json:"onPartialUpload,omitempty" tf:"on_partial_upload,omitempty"`

	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded. See Workflow Detail below.
	OnUpload []OnUploadInitParameters `json:"onUpload,omitempty" tf:"on_upload,omitempty"`
}

type WorkflowDetailsObservation struct {

	// A trigger that starts a workflow if a file is only partially uploaded. See Workflow Detail below.
	OnPartialUpload []OnPartialUploadObservation `json:"onPartialUpload,omitempty" tf:"on_partial_upload,omitempty"`

	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded. See Workflow Detail below.
	OnUpload []OnUploadObservation `json:"onUpload,omitempty" tf:"on_upload,omitempty"`
}

type WorkflowDetailsParameters struct {

	// A trigger that starts a workflow if a file is only partially uploaded. See Workflow Detail below.
	// +kubebuilder:validation:Optional
	OnPartialUpload []OnPartialUploadParameters `json:"onPartialUpload,omitempty" tf:"on_partial_upload,omitempty"`

	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded. See Workflow Detail below.
	// +kubebuilder:validation:Optional
	OnUpload []OnUploadParameters `json:"onUpload,omitempty" tf:"on_upload,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
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
	InitProvider ServerInitParameters `json:"initProvider,omitempty"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Provides a AWS Transfer Server resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
