// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a wireless gateway with a certificate.
func (c *Client) AssociateWirelessGatewayWithCertificate(ctx context.Context, params *AssociateWirelessGatewayWithCertificateInput, optFns ...func(*Options)) (*AssociateWirelessGatewayWithCertificateOutput, error) {
	if params == nil {
		params = &AssociateWirelessGatewayWithCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWirelessGatewayWithCertificate", params, optFns, addOperationAssociateWirelessGatewayWithCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWirelessGatewayWithCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWirelessGatewayWithCertificateInput struct {

	// The ID of the resource to update.
	//
	// This member is required.
	Id *string

	// The ID of the certificate to associate with the wireless gateway.
	//
	// This member is required.
	IotCertificateId *string
}

type AssociateWirelessGatewayWithCertificateOutput struct {

	// The ID of the certificate associated with the wireless gateway.
	IotCertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateWirelessGatewayWithCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateWirelessGatewayWithCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateWirelessGatewayWithCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateWirelessGatewayWithCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWirelessGatewayWithCertificate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateWirelessGatewayWithCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "AssociateWirelessGatewayWithCertificate",
	}
}