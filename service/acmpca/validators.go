// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateCertificateAuthorityAuditReport struct {
}

func (*validateOpCreateCertificateAuthorityAuditReport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCertificateAuthorityAuditReport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCertificateAuthorityAuditReportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCertificateAuthorityAuditReportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCertificateAuthority struct {
}

func (*validateOpCreateCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePermission struct {
}

func (*validateOpCreatePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCertificateAuthority struct {
}

func (*validateOpDeleteCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePermission struct {
}

func (*validateOpDeletePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePolicy struct {
}

func (*validateOpDeletePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCertificateAuthorityAuditReport struct {
}

func (*validateOpDescribeCertificateAuthorityAuditReport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCertificateAuthorityAuditReport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCertificateAuthorityAuditReportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCertificateAuthorityAuditReportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCertificateAuthority struct {
}

func (*validateOpDescribeCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCertificateAuthorityCertificate struct {
}

func (*validateOpGetCertificateAuthorityCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCertificateAuthorityCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCertificateAuthorityCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCertificateAuthorityCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCertificateAuthorityCsr struct {
}

func (*validateOpGetCertificateAuthorityCsr) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCertificateAuthorityCsr) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCertificateAuthorityCsrInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCertificateAuthorityCsrInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCertificate struct {
}

func (*validateOpGetCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPolicy struct {
}

func (*validateOpGetPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpImportCertificateAuthorityCertificate struct {
}

func (*validateOpImportCertificateAuthorityCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpImportCertificateAuthorityCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ImportCertificateAuthorityCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpImportCertificateAuthorityCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpIssueCertificate struct {
}

func (*validateOpIssueCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpIssueCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*IssueCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpIssueCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPermissions struct {
}

func (*validateOpListPermissions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPermissions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPermissionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPermissionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTags struct {
}

func (*validateOpListTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutPolicy struct {
}

func (*validateOpPutPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRestoreCertificateAuthority struct {
}

func (*validateOpRestoreCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRestoreCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RestoreCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRestoreCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRevokeCertificate struct {
}

func (*validateOpRevokeCertificate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRevokeCertificate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RevokeCertificateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRevokeCertificateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagCertificateAuthority struct {
}

func (*validateOpTagCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagCertificateAuthority struct {
}

func (*validateOpUntagCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateCertificateAuthority struct {
}

func (*validateOpUpdateCertificateAuthority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateCertificateAuthority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateCertificateAuthorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateCertificateAuthorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateCertificateAuthorityAuditReportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCertificateAuthorityAuditReport{}, middleware.After)
}

func addOpCreateCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCertificateAuthority{}, middleware.After)
}

func addOpCreatePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePermission{}, middleware.After)
}

func addOpDeleteCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCertificateAuthority{}, middleware.After)
}

func addOpDeletePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePermission{}, middleware.After)
}

func addOpDeletePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePolicy{}, middleware.After)
}

func addOpDescribeCertificateAuthorityAuditReportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCertificateAuthorityAuditReport{}, middleware.After)
}

func addOpDescribeCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCertificateAuthority{}, middleware.After)
}

func addOpGetCertificateAuthorityCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCertificateAuthorityCertificate{}, middleware.After)
}

func addOpGetCertificateAuthorityCsrValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCertificateAuthorityCsr{}, middleware.After)
}

func addOpGetCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCertificate{}, middleware.After)
}

func addOpGetPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPolicy{}, middleware.After)
}

func addOpImportCertificateAuthorityCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpImportCertificateAuthorityCertificate{}, middleware.After)
}

func addOpIssueCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpIssueCertificate{}, middleware.After)
}

func addOpListPermissionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPermissions{}, middleware.After)
}

func addOpListTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTags{}, middleware.After)
}

func addOpPutPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutPolicy{}, middleware.After)
}

func addOpRestoreCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRestoreCertificateAuthority{}, middleware.After)
}

func addOpRevokeCertificateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRevokeCertificate{}, middleware.After)
}

func addOpTagCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagCertificateAuthority{}, middleware.After)
}

func addOpUntagCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagCertificateAuthority{}, middleware.After)
}

func addOpUpdateCertificateAuthorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateCertificateAuthority{}, middleware.After)
}

func validateAccessDescription(v *types.AccessDescription) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AccessDescription"}
	if v.AccessMethod == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessMethod"))
	}
	if v.AccessLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessLocation"))
	} else if v.AccessLocation != nil {
		if err := validateGeneralName(v.AccessLocation); err != nil {
			invalidParams.AddNested("AccessLocation", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAccessDescriptionList(v []types.AccessDescription) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AccessDescriptionList"}
	for i := range v {
		if err := validateAccessDescription(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateApiPassthrough(v *types.ApiPassthrough) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApiPassthrough"}
	if v.Extensions != nil {
		if err := validateExtensions(v.Extensions); err != nil {
			invalidParams.AddNested("Extensions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCertificateAuthorityConfiguration(v *types.CertificateAuthorityConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CertificateAuthorityConfiguration"}
	if len(v.KeyAlgorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("KeyAlgorithm"))
	}
	if len(v.SigningAlgorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SigningAlgorithm"))
	}
	if v.Subject == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subject"))
	}
	if v.CsrExtensions != nil {
		if err := validateCsrExtensions(v.CsrExtensions); err != nil {
			invalidParams.AddNested("CsrExtensions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCertificatePolicyList(v []types.PolicyInformation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CertificatePolicyList"}
	for i := range v {
		if err := validatePolicyInformation(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrlConfiguration(v *types.CrlConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrlConfiguration"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCsrExtensions(v *types.CsrExtensions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CsrExtensions"}
	if v.SubjectInformationAccess != nil {
		if err := validateAccessDescriptionList(v.SubjectInformationAccess); err != nil {
			invalidParams.AddNested("SubjectInformationAccess", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEdiPartyName(v *types.EdiPartyName) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EdiPartyName"}
	if v.PartyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartyName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExtensions(v *types.Extensions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Extensions"}
	if v.CertificatePolicies != nil {
		if err := validateCertificatePolicyList(v.CertificatePolicies); err != nil {
			invalidParams.AddNested("CertificatePolicies", err.(smithy.InvalidParamsError))
		}
	}
	if v.SubjectAlternativeNames != nil {
		if err := validateGeneralNameList(v.SubjectAlternativeNames); err != nil {
			invalidParams.AddNested("SubjectAlternativeNames", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGeneralName(v *types.GeneralName) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GeneralName"}
	if v.OtherName != nil {
		if err := validateOtherName(v.OtherName); err != nil {
			invalidParams.AddNested("OtherName", err.(smithy.InvalidParamsError))
		}
	}
	if v.EdiPartyName != nil {
		if err := validateEdiPartyName(v.EdiPartyName); err != nil {
			invalidParams.AddNested("EdiPartyName", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGeneralNameList(v []types.GeneralName) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GeneralNameList"}
	for i := range v {
		if err := validateGeneralName(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOtherName(v *types.OtherName) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OtherName"}
	if v.TypeId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TypeId"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePolicyInformation(v *types.PolicyInformation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PolicyInformation"}
	if v.CertPolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertPolicyId"))
	}
	if v.PolicyQualifiers != nil {
		if err := validatePolicyQualifierInfoList(v.PolicyQualifiers); err != nil {
			invalidParams.AddNested("PolicyQualifiers", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePolicyQualifierInfo(v *types.PolicyQualifierInfo) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PolicyQualifierInfo"}
	if len(v.PolicyQualifierId) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyQualifierId"))
	}
	if v.Qualifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Qualifier"))
	} else if v.Qualifier != nil {
		if err := validateQualifier(v.Qualifier); err != nil {
			invalidParams.AddNested("Qualifier", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePolicyQualifierInfoList(v []types.PolicyQualifierInfo) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PolicyQualifierInfoList"}
	for i := range v {
		if err := validatePolicyQualifierInfo(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateQualifier(v *types.Qualifier) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Qualifier"}
	if v.CpsUri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CpsUri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRevocationConfiguration(v *types.RevocationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RevocationConfiguration"}
	if v.CrlConfiguration != nil {
		if err := validateCrlConfiguration(v.CrlConfiguration); err != nil {
			invalidParams.AddNested("CrlConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateValidity(v *types.Validity) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Validity"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCertificateAuthorityAuditReportInput(v *CreateCertificateAuthorityAuditReportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCertificateAuthorityAuditReportInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.S3BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3BucketName"))
	}
	if len(v.AuditReportResponseFormat) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AuditReportResponseFormat"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCertificateAuthorityInput(v *CreateCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCertificateAuthorityInput"}
	if v.CertificateAuthorityConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityConfiguration"))
	} else if v.CertificateAuthorityConfiguration != nil {
		if err := validateCertificateAuthorityConfiguration(v.CertificateAuthorityConfiguration); err != nil {
			invalidParams.AddNested("CertificateAuthorityConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RevocationConfiguration != nil {
		if err := validateRevocationConfiguration(v.RevocationConfiguration); err != nil {
			invalidParams.AddNested("RevocationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.CertificateAuthorityType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityType"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePermissionInput(v *CreatePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePermissionInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.Principal == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Principal"))
	}
	if v.Actions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Actions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCertificateAuthorityInput(v *DeleteCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCertificateAuthorityInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePermissionInput(v *DeletePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePermissionInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.Principal == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Principal"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePolicyInput(v *DeletePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePolicyInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCertificateAuthorityAuditReportInput(v *DescribeCertificateAuthorityAuditReportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCertificateAuthorityAuditReportInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.AuditReportId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AuditReportId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeCertificateAuthorityInput(v *DescribeCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeCertificateAuthorityInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCertificateAuthorityCertificateInput(v *GetCertificateAuthorityCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCertificateAuthorityCertificateInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCertificateAuthorityCsrInput(v *GetCertificateAuthorityCsrInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCertificateAuthorityCsrInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCertificateInput(v *GetCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCertificateInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.CertificateArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPolicyInput(v *GetPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPolicyInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpImportCertificateAuthorityCertificateInput(v *ImportCertificateAuthorityCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportCertificateAuthorityCertificateInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.Certificate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Certificate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpIssueCertificateInput(v *IssueCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IssueCertificateInput"}
	if v.ApiPassthrough != nil {
		if err := validateApiPassthrough(v.ApiPassthrough); err != nil {
			invalidParams.AddNested("ApiPassthrough", err.(smithy.InvalidParamsError))
		}
	}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.Csr == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Csr"))
	}
	if len(v.SigningAlgorithm) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SigningAlgorithm"))
	}
	if v.Validity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Validity"))
	} else if v.Validity != nil {
		if err := validateValidity(v.Validity); err != nil {
			invalidParams.AddNested("Validity", err.(smithy.InvalidParamsError))
		}
	}
	if v.ValidityNotBefore != nil {
		if err := validateValidity(v.ValidityNotBefore); err != nil {
			invalidParams.AddNested("ValidityNotBefore", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPermissionsInput(v *ListPermissionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPermissionsInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsInput(v *ListTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutPolicyInput(v *PutPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutPolicyInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRestoreCertificateAuthorityInput(v *RestoreCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RestoreCertificateAuthorityInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRevokeCertificateInput(v *RevokeCertificateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RevokeCertificateInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.CertificateSerial == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateSerial"))
	}
	if len(v.RevocationReason) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RevocationReason"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagCertificateAuthorityInput(v *TagCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagCertificateAuthorityInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagCertificateAuthorityInput(v *UntagCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagCertificateAuthorityInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateCertificateAuthorityInput(v *UpdateCertificateAuthorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateCertificateAuthorityInput"}
	if v.CertificateAuthorityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if v.RevocationConfiguration != nil {
		if err := validateRevocationConfiguration(v.RevocationConfiguration); err != nil {
			invalidParams.AddNested("RevocationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}