/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/provider-aws/apis/acm/v1beta1"
	v1beta1acmpca "github.com/upbound/provider-aws/apis/acmpca/v1beta1"
	v1beta1amp "github.com/upbound/provider-aws/apis/amp/v1beta1"
	v1beta1apigateway "github.com/upbound/provider-aws/apis/apigateway/v1beta1"
	v1beta1apigatewayv2 "github.com/upbound/provider-aws/apis/apigatewayv2/v1beta1"
	v1beta1appautoscaling "github.com/upbound/provider-aws/apis/appautoscaling/v1beta1"
	v1beta1athena "github.com/upbound/provider-aws/apis/athena/v1beta1"
	v1beta1autoscaling "github.com/upbound/provider-aws/apis/autoscaling/v1beta1"
	v1beta1backup "github.com/upbound/provider-aws/apis/backup/v1beta1"
	v1beta1cloudfront "github.com/upbound/provider-aws/apis/cloudfront/v1beta1"
	v1beta1cloudsearch "github.com/upbound/provider-aws/apis/cloudsearch/v1beta1"
	v1beta1cloudwatch "github.com/upbound/provider-aws/apis/cloudwatch/v1beta1"
	v1beta1cloudwatchlogs "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1"
	v1beta1codecommit "github.com/upbound/provider-aws/apis/codecommit/v1beta1"
	v1beta1cognitoidentity "github.com/upbound/provider-aws/apis/cognitoidentity/v1beta1"
	v1beta1cognitoidp "github.com/upbound/provider-aws/apis/cognitoidp/v1beta1"
	v1beta1dax "github.com/upbound/provider-aws/apis/dax/v1beta1"
	v1beta1deploy "github.com/upbound/provider-aws/apis/deploy/v1beta1"
	v1beta1docdb "github.com/upbound/provider-aws/apis/docdb/v1beta1"
	v1beta1dynamodb "github.com/upbound/provider-aws/apis/dynamodb/v1beta1"
	v1beta1ec2 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1ecr "github.com/upbound/provider-aws/apis/ecr/v1beta1"
	v1beta1ecrpublic "github.com/upbound/provider-aws/apis/ecrpublic/v1beta1"
	v1beta1ecs "github.com/upbound/provider-aws/apis/ecs/v1beta1"
	v1beta1efs "github.com/upbound/provider-aws/apis/efs/v1beta1"
	v1beta1eks "github.com/upbound/provider-aws/apis/eks/v1beta1"
	v1beta1elasticache "github.com/upbound/provider-aws/apis/elasticache/v1beta1"
	v1beta1elb "github.com/upbound/provider-aws/apis/elb/v1beta1"
	v1beta1elbv2 "github.com/upbound/provider-aws/apis/elbv2/v1beta1"
	v1beta1firehose "github.com/upbound/provider-aws/apis/firehose/v1beta1"
	v1beta1gamelift "github.com/upbound/provider-aws/apis/gamelift/v1beta1"
	v1beta1globalaccelerator "github.com/upbound/provider-aws/apis/globalaccelerator/v1beta1"
	v1beta1glue "github.com/upbound/provider-aws/apis/glue/v1beta1"
	v1beta1grafana "github.com/upbound/provider-aws/apis/grafana/v1beta1"
	v1beta1iam "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1iot "github.com/upbound/provider-aws/apis/iot/v1beta1"
	v1beta1kafka "github.com/upbound/provider-aws/apis/kafka/v1beta1"
	v1beta1kinesis "github.com/upbound/provider-aws/apis/kinesis/v1beta1"
	v1beta1kinesisanalytics "github.com/upbound/provider-aws/apis/kinesisanalytics/v1beta1"
	v1beta1kinesisanalyticsv2 "github.com/upbound/provider-aws/apis/kinesisanalyticsv2/v1beta1"
	v1beta1kinesisvideo "github.com/upbound/provider-aws/apis/kinesisvideo/v1beta1"
	v1beta1kms "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta1lakeformation "github.com/upbound/provider-aws/apis/lakeformation/v1beta1"
	v1beta1lambda "github.com/upbound/provider-aws/apis/lambda/v1beta1"
	v1beta1lexmodels "github.com/upbound/provider-aws/apis/lexmodels/v1beta1"
	v1beta1licensemanager "github.com/upbound/provider-aws/apis/licensemanager/v1beta1"
	v1beta1mq "github.com/upbound/provider-aws/apis/mq/v1beta1"
	v1beta1neptune "github.com/upbound/provider-aws/apis/neptune/v1beta1"
	v1beta1opensearch "github.com/upbound/provider-aws/apis/opensearch/v1beta1"
	v1beta1ram "github.com/upbound/provider-aws/apis/ram/v1beta1"
	v1beta1rds "github.com/upbound/provider-aws/apis/rds/v1beta1"
	v1beta1redshift "github.com/upbound/provider-aws/apis/redshift/v1beta1"
	v1beta1resourcegroups "github.com/upbound/provider-aws/apis/resourcegroups/v1beta1"
	v1beta1route53 "github.com/upbound/provider-aws/apis/route53/v1beta1"
	v1beta1route53resolver "github.com/upbound/provider-aws/apis/route53resolver/v1beta1"
	v1beta1s3 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	v1beta1secretsmanager "github.com/upbound/provider-aws/apis/secretsmanager/v1beta1"
	v1beta1servicediscovery "github.com/upbound/provider-aws/apis/servicediscovery/v1beta1"
	v1beta1sfn "github.com/upbound/provider-aws/apis/sfn/v1beta1"
	v1beta1signer "github.com/upbound/provider-aws/apis/signer/v1beta1"
	v1beta1sns "github.com/upbound/provider-aws/apis/sns/v1beta1"
	v1beta1sqs "github.com/upbound/provider-aws/apis/sqs/v1beta1"
	v1beta1transfer "github.com/upbound/provider-aws/apis/transfer/v1beta1"
	v1beta1apis "github.com/upbound/provider-aws/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1acmpca.SchemeBuilder.AddToScheme,
		v1beta1amp.SchemeBuilder.AddToScheme,
		v1beta1apigateway.SchemeBuilder.AddToScheme,
		v1beta1apigatewayv2.SchemeBuilder.AddToScheme,
		v1beta1appautoscaling.SchemeBuilder.AddToScheme,
		v1beta1athena.SchemeBuilder.AddToScheme,
		v1beta1autoscaling.SchemeBuilder.AddToScheme,
		v1beta1backup.SchemeBuilder.AddToScheme,
		v1beta1cloudfront.SchemeBuilder.AddToScheme,
		v1beta1cloudsearch.SchemeBuilder.AddToScheme,
		v1beta1cloudwatch.SchemeBuilder.AddToScheme,
		v1beta1cloudwatchlogs.SchemeBuilder.AddToScheme,
		v1beta1codecommit.SchemeBuilder.AddToScheme,
		v1beta1cognitoidentity.SchemeBuilder.AddToScheme,
		v1beta1cognitoidp.SchemeBuilder.AddToScheme,
		v1beta1dax.SchemeBuilder.AddToScheme,
		v1beta1deploy.SchemeBuilder.AddToScheme,
		v1beta1docdb.SchemeBuilder.AddToScheme,
		v1beta1dynamodb.SchemeBuilder.AddToScheme,
		v1beta1ec2.SchemeBuilder.AddToScheme,
		v1beta1ecr.SchemeBuilder.AddToScheme,
		v1beta1ecrpublic.SchemeBuilder.AddToScheme,
		v1beta1ecs.SchemeBuilder.AddToScheme,
		v1beta1efs.SchemeBuilder.AddToScheme,
		v1beta1eks.SchemeBuilder.AddToScheme,
		v1beta1elasticache.SchemeBuilder.AddToScheme,
		v1beta1elb.SchemeBuilder.AddToScheme,
		v1beta1elbv2.SchemeBuilder.AddToScheme,
		v1beta1firehose.SchemeBuilder.AddToScheme,
		v1beta1gamelift.SchemeBuilder.AddToScheme,
		v1beta1globalaccelerator.SchemeBuilder.AddToScheme,
		v1beta1glue.SchemeBuilder.AddToScheme,
		v1beta1grafana.SchemeBuilder.AddToScheme,
		v1beta1iam.SchemeBuilder.AddToScheme,
		v1beta1iot.SchemeBuilder.AddToScheme,
		v1beta1kafka.SchemeBuilder.AddToScheme,
		v1beta1kinesis.SchemeBuilder.AddToScheme,
		v1beta1kinesisanalytics.SchemeBuilder.AddToScheme,
		v1beta1kinesisanalyticsv2.SchemeBuilder.AddToScheme,
		v1beta1kinesisvideo.SchemeBuilder.AddToScheme,
		v1beta1kms.SchemeBuilder.AddToScheme,
		v1beta1lakeformation.SchemeBuilder.AddToScheme,
		v1beta1lambda.SchemeBuilder.AddToScheme,
		v1beta1lexmodels.SchemeBuilder.AddToScheme,
		v1beta1licensemanager.SchemeBuilder.AddToScheme,
		v1beta1mq.SchemeBuilder.AddToScheme,
		v1beta1neptune.SchemeBuilder.AddToScheme,
		v1beta1opensearch.SchemeBuilder.AddToScheme,
		v1beta1ram.SchemeBuilder.AddToScheme,
		v1beta1rds.SchemeBuilder.AddToScheme,
		v1beta1redshift.SchemeBuilder.AddToScheme,
		v1beta1resourcegroups.SchemeBuilder.AddToScheme,
		v1beta1route53.SchemeBuilder.AddToScheme,
		v1beta1route53resolver.SchemeBuilder.AddToScheme,
		v1beta1s3.SchemeBuilder.AddToScheme,
		v1beta1secretsmanager.SchemeBuilder.AddToScheme,
		v1beta1servicediscovery.SchemeBuilder.AddToScheme,
		v1beta1sfn.SchemeBuilder.AddToScheme,
		v1beta1signer.SchemeBuilder.AddToScheme,
		v1beta1sns.SchemeBuilder.AddToScheme,
		v1beta1sqs.SchemeBuilder.AddToScheme,
		v1beta1transfer.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
