// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Cluster.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.IAMRoles),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.IAMRoleRefs,
			Selector:      mg.Spec.ForProvider.IAMRoleSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoles")
	}
	mg.Spec.ForProvider.IAMRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.IAMRoleRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterParameterGroup", "ClusterParameterGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneClusterParameterGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef,
			Selector:     mg.Spec.ForProvider.NeptuneClusterParameterGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneClusterParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
			Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationSourceIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ReplicationSourceIdentifierRef,
			Selector:     mg.Spec.ForProvider.ReplicationSourceIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationSourceIdentifier")
	}
	mg.Spec.ForProvider.ReplicationSourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationSourceIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterSnapshot", "ClusterSnapshotList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.SnapshotIdentifierRef,
			Selector:     mg.Spec.ForProvider.SnapshotIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotIdentifier")
	}
	mg.Spec.ForProvider.SnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
			Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.IAMRoles),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.IAMRoleRefs,
			Selector:      mg.Spec.InitProvider.IAMRoleSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRoles")
	}
	mg.Spec.InitProvider.IAMRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.IAMRoleRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.KMSKeyArnRef,
			Selector:     mg.Spec.InitProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyArn")
	}
	mg.Spec.InitProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterParameterGroup", "ClusterParameterGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NeptuneClusterParameterGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NeptuneClusterParameterGroupNameRef,
			Selector:     mg.Spec.InitProvider.NeptuneClusterParameterGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NeptuneClusterParameterGroupName")
	}
	mg.Spec.InitProvider.NeptuneClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NeptuneClusterParameterGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "SubnetGroup", "SubnetGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NeptuneSubnetGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.NeptuneSubnetGroupNameRef,
			Selector:     mg.Spec.InitProvider.NeptuneSubnetGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.InitProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ReplicationSourceIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ReplicationSourceIdentifierRef,
			Selector:     mg.Spec.InitProvider.ReplicationSourceIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ReplicationSourceIdentifier")
	}
	mg.Spec.InitProvider.ReplicationSourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ReplicationSourceIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("neptune.aws.upbound.io", "v1beta1", "ClusterSnapshot", "ClusterSnapshotList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnapshotIdentifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.SnapshotIdentifierRef,
			Selector:     mg.Spec.InitProvider.SnapshotIdentifierSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnapshotIdentifier")
	}
	mg.Spec.InitProvider.SnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnapshotIdentifierRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCSecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.VPCSecurityGroupIDRefs,
			Selector:      mg.Spec.InitProvider.VPCSecurityGroupIDSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCSecurityGroupIds")
	}
	mg.Spec.InitProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}
