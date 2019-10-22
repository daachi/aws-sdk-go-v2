// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDigestInput struct {
	_ struct{} `type:"structure"`

	// The name of the ledger.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDigestInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDigestInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDigestInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDigestInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDigestOutput struct {
	_ struct{} `type:"structure"`

	// The 256-bit hash value representing the digest returned by a GetDigest request.
	//
	// Digest is automatically base64 encoded/decoded by the SDK.
	//
	// Digest is a required field
	Digest []byte `min:"32" type:"blob" required:"true"`

	// The latest block location covered by the digest that you requested. An address
	// is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	//
	// DigestTipAddress is a required field
	DigestTipAddress *ValueHolder `type:"structure" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s GetDigestOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDigestOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Digest != nil {
		v := s.Digest

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Digest", protocol.QuotedValue{ValueMarshaler: protocol.BytesValue(v)}, metadata)
	}
	if s.DigestTipAddress != nil {
		v := s.DigestTipAddress

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DigestTipAddress", v, metadata)
	}
	return nil
}

const opGetDigest = "GetDigest"

// GetDigestRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Returns the digest of a ledger at the latest committed block in the journal.
// The response includes a 256-bit hash value and a block address.
//
//    // Example sending a request using GetDigestRequest.
//    req := client.GetDigestRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/GetDigest
func (c *Client) GetDigestRequest(input *GetDigestInput) GetDigestRequest {
	op := &aws.Operation{
		Name:       opGetDigest,
		HTTPMethod: "POST",
		HTTPPath:   "/ledgers/{name}/digest",
	}

	if input == nil {
		input = &GetDigestInput{}
	}

	req := c.newRequest(op, input, &GetDigestOutput{})
	return GetDigestRequest{Request: req, Input: input, Copy: c.GetDigestRequest}
}

// GetDigestRequest is the request type for the
// GetDigest API operation.
type GetDigestRequest struct {
	*aws.Request
	Input *GetDigestInput
	Copy  func(*GetDigestInput) GetDigestRequest
}

// Send marshals and sends the GetDigest API request.
func (r GetDigestRequest) Send(ctx context.Context) (*GetDigestResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDigestResponse{
		GetDigestOutput: r.Request.Data.(*GetDigestOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDigestResponse is the response type for the
// GetDigest API operation.
type GetDigestResponse struct {
	*GetDigestOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDigest request.
func (r *GetDigestResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}