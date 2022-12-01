// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update wave.
func (c *Client) UpdateWave(ctx context.Context, params *UpdateWaveInput, optFns ...func(*Options)) (*UpdateWaveOutput, error) {
	if params == nil {
		params = &UpdateWaveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWave", params, optFns, c.addOperationUpdateWaveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWaveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWaveInput struct {

	// Wave ID.
	//
	// This member is required.
	WaveID *string

	// Wave description.
	Description *string

	// Wave name.
	Name *string

	noSmithyDocumentSerde
}

type UpdateWaveOutput struct {

	// Wave ARN.
	Arn *string

	// Wave creation dateTime.
	CreationDateTime *string

	// Wave description.
	Description *string

	// Wave archival status.
	IsArchived *bool

	// Wave last modified dateTime.
	LastModifiedDateTime *string

	// Wave name.
	Name *string

	// Wave tags.
	Tags map[string]string

	// Wave aggregated status.
	WaveAggregatedStatus *types.WaveAggregatedStatus

	// Wave ID.
	WaveID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWaveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWave{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWave{}, middleware.After)
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
	if err = addOpUpdateWaveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWave(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWave(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "UpdateWave",
	}
}