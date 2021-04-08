// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2instanceconnect

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpSendSerialConsoleSSHPublicKey struct {
}

func (*validateOpSendSerialConsoleSSHPublicKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendSerialConsoleSSHPublicKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendSerialConsoleSSHPublicKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendSerialConsoleSSHPublicKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendSSHPublicKey struct {
}

func (*validateOpSendSSHPublicKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendSSHPublicKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendSSHPublicKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendSSHPublicKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpSendSerialConsoleSSHPublicKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendSerialConsoleSSHPublicKey{}, middleware.After)
}

func addOpSendSSHPublicKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendSSHPublicKey{}, middleware.After)
}

func validateOpSendSerialConsoleSSHPublicKeyInput(v *SendSerialConsoleSSHPublicKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendSerialConsoleSSHPublicKeyInput"}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.SSHPublicKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SSHPublicKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendSSHPublicKeyInput(v *SendSSHPublicKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendSSHPublicKeyInput"}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.InstanceOSUser == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceOSUser"))
	}
	if v.SSHPublicKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SSHPublicKey"))
	}
	if v.AvailabilityZone == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZone"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}