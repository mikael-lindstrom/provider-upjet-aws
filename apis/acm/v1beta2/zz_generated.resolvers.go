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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Certificate) ResolveReferences( // ResolveReferences of this Certificate.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("acmpca.aws.upbound.io", "v1beta2", "CertificateAuthority", "CertificateAuthorityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateAuthorityArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CertificateAuthorityArnRef,
			Selector:     mg.Spec.ForProvider.CertificateAuthorityArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateAuthorityArn")
	}
	mg.Spec.ForProvider.CertificateAuthorityArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateAuthorityArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acmpca.aws.upbound.io", "v1beta2", "CertificateAuthority", "CertificateAuthorityList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CertificateAuthorityArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.CertificateAuthorityArnRef,
			Selector:     mg.Spec.InitProvider.CertificateAuthorityArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CertificateAuthorityArn")
	}
	mg.Spec.InitProvider.CertificateAuthorityArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateAuthorityArnRef = rsp.ResolvedReference

	return nil
}
