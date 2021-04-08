// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpConfigureLogs struct {
}

func (*validateOpConfigureLogs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConfigureLogs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConfigureLogsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConfigureLogsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAsset struct {
}

func (*validateOpCreateAsset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAsset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAssetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAssetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePackagingConfiguration struct {
}

func (*validateOpCreatePackagingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePackagingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePackagingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePackagingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePackagingGroup struct {
}

func (*validateOpCreatePackagingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePackagingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePackagingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePackagingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAsset struct {
}

func (*validateOpDeleteAsset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAsset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAssetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAssetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePackagingConfiguration struct {
}

func (*validateOpDeletePackagingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePackagingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePackagingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePackagingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePackagingGroup struct {
}

func (*validateOpDeletePackagingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePackagingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePackagingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePackagingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAsset struct {
}

func (*validateOpDescribeAsset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAsset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAssetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAssetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePackagingConfiguration struct {
}

func (*validateOpDescribePackagingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePackagingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePackagingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePackagingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePackagingGroup struct {
}

func (*validateOpDescribePackagingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePackagingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePackagingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePackagingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdatePackagingGroup struct {
}

func (*validateOpUpdatePackagingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePackagingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePackagingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePackagingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpConfigureLogsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConfigureLogs{}, middleware.After)
}

func addOpCreateAssetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAsset{}, middleware.After)
}

func addOpCreatePackagingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePackagingConfiguration{}, middleware.After)
}

func addOpCreatePackagingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePackagingGroup{}, middleware.After)
}

func addOpDeleteAssetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAsset{}, middleware.After)
}

func addOpDeletePackagingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePackagingConfiguration{}, middleware.After)
}

func addOpDeletePackagingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePackagingGroup{}, middleware.After)
}

func addOpDescribeAssetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAsset{}, middleware.After)
}

func addOpDescribePackagingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePackagingConfiguration{}, middleware.After)
}

func addOpDescribePackagingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePackagingGroup{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdatePackagingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePackagingGroup{}, middleware.After)
}

func validateAuthorization(v *types.Authorization) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Authorization"}
	if v.CdnIdentifierSecret == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CdnIdentifierSecret"))
	}
	if v.SecretsRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecretsRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCmafEncryption(v *types.CmafEncryption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CmafEncryption"}
	if v.SpekeKeyProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpekeKeyProvider"))
	} else if v.SpekeKeyProvider != nil {
		if err := validateSpekeKeyProvider(v.SpekeKeyProvider); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCmafPackage(v *types.CmafPackage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CmafPackage"}
	if v.Encryption != nil {
		if err := validateCmafEncryption(v.Encryption); err != nil {
			invalidParams.AddNested("Encryption", err.(smithy.InvalidParamsError))
		}
	}
	if v.HlsManifests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HlsManifests"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDashEncryption(v *types.DashEncryption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DashEncryption"}
	if v.SpekeKeyProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpekeKeyProvider"))
	} else if v.SpekeKeyProvider != nil {
		if err := validateSpekeKeyProvider(v.SpekeKeyProvider); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDashPackage(v *types.DashPackage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DashPackage"}
	if v.DashManifests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DashManifests"))
	}
	if v.Encryption != nil {
		if err := validateDashEncryption(v.Encryption); err != nil {
			invalidParams.AddNested("Encryption", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHlsEncryption(v *types.HlsEncryption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HlsEncryption"}
	if v.SpekeKeyProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpekeKeyProvider"))
	} else if v.SpekeKeyProvider != nil {
		if err := validateSpekeKeyProvider(v.SpekeKeyProvider); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHlsPackage(v *types.HlsPackage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HlsPackage"}
	if v.Encryption != nil {
		if err := validateHlsEncryption(v.Encryption); err != nil {
			invalidParams.AddNested("Encryption", err.(smithy.InvalidParamsError))
		}
	}
	if v.HlsManifests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("HlsManifests"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMssEncryption(v *types.MssEncryption) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MssEncryption"}
	if v.SpekeKeyProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SpekeKeyProvider"))
	} else if v.SpekeKeyProvider != nil {
		if err := validateSpekeKeyProvider(v.SpekeKeyProvider); err != nil {
			invalidParams.AddNested("SpekeKeyProvider", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMssPackage(v *types.MssPackage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MssPackage"}
	if v.Encryption != nil {
		if err := validateMssEncryption(v.Encryption); err != nil {
			invalidParams.AddNested("Encryption", err.(smithy.InvalidParamsError))
		}
	}
	if v.MssManifests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MssManifests"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSpekeKeyProvider(v *types.SpekeKeyProvider) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SpekeKeyProvider"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.SystemIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SystemIds"))
	}
	if v.Url == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Url"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpConfigureLogsInput(v *ConfigureLogsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigureLogsInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAssetInput(v *CreateAssetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAssetInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.PackagingGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PackagingGroupId"))
	}
	if v.SourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceArn"))
	}
	if v.SourceRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePackagingConfigurationInput(v *CreatePackagingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePackagingConfigurationInput"}
	if v.CmafPackage != nil {
		if err := validateCmafPackage(v.CmafPackage); err != nil {
			invalidParams.AddNested("CmafPackage", err.(smithy.InvalidParamsError))
		}
	}
	if v.DashPackage != nil {
		if err := validateDashPackage(v.DashPackage); err != nil {
			invalidParams.AddNested("DashPackage", err.(smithy.InvalidParamsError))
		}
	}
	if v.HlsPackage != nil {
		if err := validateHlsPackage(v.HlsPackage); err != nil {
			invalidParams.AddNested("HlsPackage", err.(smithy.InvalidParamsError))
		}
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.MssPackage != nil {
		if err := validateMssPackage(v.MssPackage); err != nil {
			invalidParams.AddNested("MssPackage", err.(smithy.InvalidParamsError))
		}
	}
	if v.PackagingGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PackagingGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePackagingGroupInput(v *CreatePackagingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePackagingGroupInput"}
	if v.Authorization != nil {
		if err := validateAuthorization(v.Authorization); err != nil {
			invalidParams.AddNested("Authorization", err.(smithy.InvalidParamsError))
		}
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAssetInput(v *DeleteAssetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAssetInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePackagingConfigurationInput(v *DeletePackagingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePackagingConfigurationInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePackagingGroupInput(v *DeletePackagingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePackagingGroupInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAssetInput(v *DescribeAssetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAssetInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePackagingConfigurationInput(v *DescribePackagingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePackagingConfigurationInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePackagingGroupInput(v *DescribePackagingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePackagingGroupInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdatePackagingGroupInput(v *UpdatePackagingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePackagingGroupInput"}
	if v.Authorization != nil {
		if err := validateAuthorization(v.Authorization); err != nil {
			invalidParams.AddNested("Authorization", err.(smithy.InvalidParamsError))
		}
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}