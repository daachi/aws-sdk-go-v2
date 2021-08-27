// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes user password(s) and/or access string.
func (c *Client) ModifyUser(ctx context.Context, params *ModifyUserInput, optFns ...func(*Options)) (*ModifyUserOutput, error) {
	if params == nil {
		params = &ModifyUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyUser", params, optFns, c.addOperationModifyUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyUserInput struct {

	// The ID of the user.
	//
	// This member is required.
	UserId *string

	// Access permissions string used for this user.
	AccessString *string

	// Adds additional user permissions to the access string.
	AppendAccessString *string

	// Indicates no password is required for the user.
	NoPasswordRequired *bool

	// The passwords belonging to the user. You are allowed up to two.
	Passwords []string

	noSmithyDocumentSerde
}

type ModifyUserOutput struct {

	// The Amazon Resource Name (ARN) of the user.
	ARN *string

	// Access permissions string used for this user.
	AccessString *string

	// Denotes whether the user requires a password to authenticate.
	Authentication *types.Authentication

	// The current supported value is Redis.
	Engine *string

	// Indicates the user status. Can be "active", "modifying" or "deleting".
	Status *string

	// Returns a list of the user group IDs the user belongs to.
	UserGroupIds []string

	// The ID of the user.
	UserId *string

	// The username of the user.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyUser{}, middleware.After)
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
	if err = addOpModifyUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyUser",
	}
}