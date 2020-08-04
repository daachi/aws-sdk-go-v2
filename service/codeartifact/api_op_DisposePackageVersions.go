// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisposePackageVersionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that contains the repository you want to dispose.
	//
	// Domain is a required field
	Domain *string `location:"querystring" locationName:"domain" min:"2" type:"string" required:"true"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `location:"querystring" locationName:"domain-owner" min:"12" type:"string"`

	// The expected status of the package version to dispose. Valid values are:
	//
	//    * Published
	//
	//    * Unfinished
	//
	//    * Unlisted
	//
	//    * Archived
	//
	//    * Disposed
	ExpectedStatus PackageVersionStatus `locationName:"expectedStatus" type:"string" enum:"true"`

	// A format that specifies the type of package versions you want to dispose.
	// The valid values are:
	//
	//    * npm
	//
	//    * pypi
	//
	//    * maven
	//
	// Format is a required field
	Format PackageFormat `location:"querystring" locationName:"format" type:"string" required:"true" enum:"true"`

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//    * The namespace of a Maven package is its groupId.
	//
	//    * The namespace of an npm package is its scope.
	//
	//    * A Python package does not contain a corresponding component, so Python
	//    packages do not have a namespace.
	Namespace *string `location:"querystring" locationName:"namespace" min:"1" type:"string"`

	// The name of the package with the versions you want to dispose.
	//
	// Package is a required field
	Package *string `location:"querystring" locationName:"package" min:"1" type:"string" required:"true"`

	// The name of the repository that contains the package versions you want to
	// dispose.
	//
	// Repository is a required field
	Repository *string `location:"querystring" locationName:"repository" min:"2" type:"string" required:"true"`

	// The revisions of the package versions you want to dispose.
	VersionRevisions map[string]string `locationName:"versionRevisions" type:"map"`

	// The versions of the package you want to dispose.
	//
	// Versions is a required field
	Versions []string `locationName:"versions" type:"list" required:"true"`
}

// String returns the string representation
func (s DisposePackageVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisposePackageVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisposePackageVersionsInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 2))
	}
	if s.DomainOwner != nil && len(*s.DomainOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainOwner", 12))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}

	if s.Package == nil {
		invalidParams.Add(aws.NewErrParamRequired("Package"))
	}
	if s.Package != nil && len(*s.Package) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Package", 1))
	}

	if s.Repository == nil {
		invalidParams.Add(aws.NewErrParamRequired("Repository"))
	}
	if s.Repository != nil && len(*s.Repository) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Repository", 2))
	}

	if s.Versions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Versions"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisposePackageVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.ExpectedStatus) > 0 {
		v := s.ExpectedStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expectedStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.VersionRevisions != nil {
		v := s.VersionRevisions

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "versionRevisions", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Versions != nil {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain-owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Package != nil {
		v := *s.Package

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "package", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Repository != nil {
		v := *s.Repository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "repository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisposePackageVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A PackageVersionError object that contains a map of errors codes for the
	// disposed package versions that failed. The possible error codes are:
	//
	//    * ALREADY_EXISTS
	//
	//    * MISMATCHED_REVISION
	//
	//    * MISMATCHED_STATUS
	//
	//    * NOT_ALLOWED
	//
	//    * NOT_FOUND
	//
	//    * SKIPPED
	FailedVersions map[string]PackageVersionError `locationName:"failedVersions" type:"map"`

	// A list of the package versions that were successfully disposed.
	SuccessfulVersions map[string]SuccessfulPackageVersionInfo `locationName:"successfulVersions" type:"map"`
}

// String returns the string representation
func (s DisposePackageVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisposePackageVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FailedVersions != nil {
		v := s.FailedVersions

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "failedVersions", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.SuccessfulVersions != nil {
		v := s.SuccessfulVersions

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "successfulVersions", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	return nil
}

const opDisposePackageVersions = "DisposePackageVersions"

// DisposePackageVersionsRequest returns a request value for making API operation for
// CodeArtifact.
//
// Deletes the assets in package versions and sets the package versions' status
// to Disposed. A disposed package version cannot be restored in your repository
// because its assets are deleted.
//
// To view all disposed package versions in a repository, use ListackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)
// and set the status (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html#API_ListPackageVersions_RequestSyntax)
// parameter to Disposed.
//
// To view information about a disposed package version, use ListPackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)
// and set the status (https://docs.aws.amazon.com/API_ListPackageVersions.html#codeartifact-ListPackageVersions-response-status)
// parameter to Disposed.
//
//    // Example sending a request using DisposePackageVersionsRequest.
//    req := client.DisposePackageVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeartifact-2018-09-22/DisposePackageVersions
func (c *Client) DisposePackageVersionsRequest(input *DisposePackageVersionsInput) DisposePackageVersionsRequest {
	op := &aws.Operation{
		Name:       opDisposePackageVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/package/versions/dispose",
	}

	if input == nil {
		input = &DisposePackageVersionsInput{}
	}

	req := c.newRequest(op, input, &DisposePackageVersionsOutput{})

	return DisposePackageVersionsRequest{Request: req, Input: input, Copy: c.DisposePackageVersionsRequest}
}

// DisposePackageVersionsRequest is the request type for the
// DisposePackageVersions API operation.
type DisposePackageVersionsRequest struct {
	*aws.Request
	Input *DisposePackageVersionsInput
	Copy  func(*DisposePackageVersionsInput) DisposePackageVersionsRequest
}

// Send marshals and sends the DisposePackageVersions API request.
func (r DisposePackageVersionsRequest) Send(ctx context.Context) (*DisposePackageVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisposePackageVersionsResponse{
		DisposePackageVersionsOutput: r.Request.Data.(*DisposePackageVersionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisposePackageVersionsResponse is the response type for the
// DisposePackageVersions API operation.
type DisposePackageVersionsResponse struct {
	*DisposePackageVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisposePackageVersions request.
func (r *DisposePackageVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}