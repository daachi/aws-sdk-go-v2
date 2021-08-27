// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Amazon S3 on Outposts Access Points simplify managing data access at scale for
// shared datasets in S3 on Outposts. S3 on Outposts uses endpoints to connect to
// Outposts buckets so that you can perform actions within your virtual private
// cloud (VPC). For more information, see  Accessing S3 on Outposts using VPC only
// access points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/AccessingS3Outposts.html).
type Endpoint struct {

	//
	AccessType EndpointAccessType

	// The VPC CIDR committed by this endpoint.
	CidrBlock *string

	// The time the endpoint was created.
	CreationTime *time.Time

	// The ID of the customer-owned IPv4 pool used for the endpoint.
	CustomerOwnedIpv4Pool *string

	// The Amazon Resource Name (ARN) of the endpoint.
	EndpointArn *string

	// The network interface of the endpoint.
	NetworkInterfaces []NetworkInterface

	// The ID of the AWS Outposts.
	OutpostsId *string

	// The ID of the security group used for the endpoint.
	SecurityGroupId *string

	// The status of the endpoint.
	Status EndpointStatus

	// The ID of the subnet used for the endpoint.
	SubnetId *string

	// The ID of the VPC used for the endpoint.
	VpcId *string

	noSmithyDocumentSerde
}

// The container for the network interface.
type NetworkInterface struct {

	// The ID for the network interface.
	NetworkInterfaceId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde