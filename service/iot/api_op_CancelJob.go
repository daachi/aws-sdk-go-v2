// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CancelJobInput struct {
	_ struct{} `type:"structure"`

	// An optional comment string describing why the job was canceled.
	Comment *string `locationName:"comment" type:"string"`

	// (Optional) If true job executions with status "IN_PROGRESS" and "QUEUED"
	// are canceled, otherwise only job executions with status "QUEUED" are canceled.
	// The default is false.
	//
	// Canceling a job which is "IN_PROGRESS", will cause a device which is executing
	// the job to be unable to update the job execution status. Use caution and
	// ensure that each device executing a job which is canceled is able to recover
	// to a valid state.
	Force *bool `location:"querystring" locationName:"force" type:"boolean"`

	// The unique identifier you assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// (Optional)A reason code string that explains why the job was canceled.
	ReasonCode *string `locationName:"reasonCode" type:"string"`
}

// String returns the string representation
func (s CancelJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Comment != nil {
		v := *s.Comment

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "comment", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReasonCode != nil {
		v := *s.ReasonCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reasonCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Force != nil {
		v := *s.Force

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "force", protocol.BoolValue(v), metadata)
	}
	return nil
}

type CancelJobOutput struct {
	_ struct{} `type:"structure"`

	// A short text description of the job.
	Description *string `locationName:"description" type:"string"`

	// The job ARN.
	JobArn *string `locationName:"jobArn" type:"string"`

	// The unique identifier you assigned to this job when it was created.
	JobId *string `locationName:"jobId" min:"1" type:"string"`
}

// String returns the string representation
func (s CancelJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobArn != nil {
		v := *s.JobArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCancelJob = "CancelJob"

// CancelJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Cancels a job.
//
//    // Example sending a request using CancelJobRequest.
//    req := client.CancelJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CancelJobRequest(input *CancelJobInput) CancelJobRequest {
	op := &aws.Operation{
		Name:       opCancelJob,
		HTTPMethod: "PUT",
		HTTPPath:   "/jobs/{jobId}/cancel",
	}

	if input == nil {
		input = &CancelJobInput{}
	}

	req := c.newRequest(op, input, &CancelJobOutput{})
	return CancelJobRequest{Request: req, Input: input, Copy: c.CancelJobRequest}
}

// CancelJobRequest is the request type for the
// CancelJob API operation.
type CancelJobRequest struct {
	*aws.Request
	Input *CancelJobInput
	Copy  func(*CancelJobInput) CancelJobRequest
}

// Send marshals and sends the CancelJob API request.
func (r CancelJobRequest) Send(ctx context.Context) (*CancelJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelJobResponse{
		CancelJobOutput: r.Request.Data.(*CancelJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelJobResponse is the response type for the
// CancelJob API operation.
type CancelJobResponse struct {
	*CancelJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelJob request.
func (r *CancelJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}