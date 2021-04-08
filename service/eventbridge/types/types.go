// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains details about an API destination.
type ApiDestination struct {

	// The ARN of the API destination.
	ApiDestinationArn *string

	// The state of the API destination.
	ApiDestinationState ApiDestinationState

	// The ARN of the connection specified for the API destination.
	ConnectionArn *string

	// A time stamp for the time that the API destination was created.
	CreationTime *time.Time

	// The method to use to connect to the HTTP endpoint.
	HttpMethod ApiDestinationHttpMethod

	// The URL to the endpoint for the API destination.
	InvocationEndpoint *string

	// The maximum number of invocations per second to send to the HTTP endpoint.
	InvocationRateLimitPerSecond *int32

	// A time stamp for the time that the API destination was last modified.
	LastModifiedTime *time.Time

	// The name of the API destination.
	Name *string
}

// An Archive object that contains details about an archive.
type Archive struct {

	// The name of the archive.
	ArchiveName *string

	// The time stamp for the time that the archive was created.
	CreationTime *time.Time

	// The number of events in the archive.
	EventCount int64

	// The ARN of the event bus associated with the archive. Only events from this
	// event bus are sent to the archive.
	EventSourceArn *string

	// The number of days to retain events in the archive before they are deleted.
	RetentionDays *int32

	// The size of the archive, in bytes.
	SizeBytes int64

	// The current state of the archive.
	State ArchiveState

	// A description for the reason that the archive is in the current state.
	StateReason *string
}

// This structure specifies the VPC subnets and security groups for the task, and
// whether a public IP address is to be used. This structure is relevant only for
// ECS tasks that use the awsvpc network mode.
type AwsVpcConfiguration struct {

	// Specifies the subnets associated with the task. These subnets must all be in the
	// same VPC. You can specify as many as 16 subnets.
	//
	// This member is required.
	Subnets []string

	// Specifies whether the task's elastic network interface receives a public IP
	// address. You can specify ENABLED only when LaunchType in EcsParameters is set to
	// FARGATE.
	AssignPublicIp AssignPublicIp

	// Specifies the security groups associated with the task. These security groups
	// must all be in the same VPC. You can specify as many as five security groups. If
	// you do not specify a security group, the default security group for the VPC is
	// used.
	SecurityGroups []string
}

// The array properties for the submitted job, such as the size of the array. The
// array size can be between 2 and 10,000. If you specify array properties for a
// job, it becomes an array job. This parameter is used only if the target is an
// AWS Batch job.
type BatchArrayProperties struct {

	// The size of the array, if this is an array batch job. Valid values are integers
	// between 2 and 10,000.
	Size int32
}

// The custom parameters to be used when the target is an AWS Batch job.
type BatchParameters struct {

	// The ARN or name of the job definition to use if the event target is an AWS Batch
	// job. This job definition must already exist.
	//
	// This member is required.
	JobDefinition *string

	// The name to use for this execution of the job, if the target is an AWS Batch
	// job.
	//
	// This member is required.
	JobName *string

	// The array properties for the submitted job, such as the size of the array. The
	// array size can be between 2 and 10,000. If you specify array properties for a
	// job, it becomes an array job. This parameter is used only if the target is an
	// AWS Batch job.
	ArrayProperties *BatchArrayProperties

	// The retry strategy to use for failed jobs, if the target is an AWS Batch job.
	// The retry strategy is the number of times to retry the failed job execution.
	// Valid values are 1–10. When you specify a retry strategy here, it overrides the
	// retry strategy defined in the job definition.
	RetryStrategy *BatchRetryStrategy
}

// The retry strategy to use for failed jobs, if the target is an AWS Batch job. If
// you specify a retry strategy here, it overrides the retry strategy defined in
// the job definition.
type BatchRetryStrategy struct {

	// The number of times to attempt to retry, if the job fails. Valid values are
	// 1–10.
	Attempts int32
}

// A JSON string which you can use to limit the event bus permissions you are
// granting to only accounts that fulfill the condition. Currently, the only
// supported condition is membership in a certain AWS organization. The string must
// contain Type, Key, and Value fields. The Value field specifies the ID of the AWS
// organization. Following is an example value for Condition: '{"Type" :
// "StringEquals", "Key": "aws:PrincipalOrgID", "Value": "o-1234567890"}'
type Condition struct {

	// Specifies the key for the condition. Currently the only supported key is
	// aws:PrincipalOrgID.
	//
	// This member is required.
	Key *string

	// Specifies the type of condition. Currently the only supported value is
	// StringEquals.
	//
	// This member is required.
	Type *string

	// Specifies the value for the key. Currently, this must be the ID of the
	// organization.
	//
	// This member is required.
	Value *string
}

// Contains information about a connection.
type Connection struct {

	// The authorization type specified for the connection.
	AuthorizationType ConnectionAuthorizationType

	// The ARN of the connection.
	ConnectionArn *string

	// The state of the connection.
	ConnectionState ConnectionState

	// A time stamp for the time that the connection was created.
	CreationTime *time.Time

	// A time stamp for the time that the connection was last authorized.
	LastAuthorizedTime *time.Time

	// A time stamp for the time that the connection was last modified.
	LastModifiedTime *time.Time

	// The name of the connection.
	Name *string

	// The reason that the connection is in the connection state.
	StateReason *string
}

// Contains the authorization parameters for the connection if API Key is specified
// as the authorization type.
type ConnectionApiKeyAuthResponseParameters struct {

	// The name of the header to use for the APIKeyValue used for authorization.
	ApiKeyName *string
}

// Contains the authorization parameters to use for the connection.
type ConnectionAuthResponseParameters struct {

	// The API Key parameters to use for authorization.
	ApiKeyAuthParameters *ConnectionApiKeyAuthResponseParameters

	// The authorization parameters for Basic authorization.
	BasicAuthParameters *ConnectionBasicAuthResponseParameters

	// Additional parameters for the connection that are passed through with every
	// invocation to the HTTP endpoint.
	InvocationHttpParameters *ConnectionHttpParameters

	// The OAuth parameters to use for authorization.
	OAuthParameters *ConnectionOAuthResponseParameters
}

// Contains the authorization parameters for the connection if Basic is specified
// as the authorization type.
type ConnectionBasicAuthResponseParameters struct {

	// The user name to use for Basic authorization.
	Username *string
}

// Additional parameter included in the body. You can include up to 100 additional
// body parameters per request. An event payload cannot exceed 64 KB.
type ConnectionBodyParameter struct {

	// Specified whether the value is secret.
	IsValueSecret bool

	// The key for the parameter.
	Key *string

	// The value associated with the key.
	Value *string
}

// Additional parameter included in the header. You can include up to 100
// additional header parameters per request. An event payload cannot exceed 64 KB.
type ConnectionHeaderParameter struct {

	// Specified whether the value is a secret.
	IsValueSecret bool

	// The key for the parameter.
	Key *string

	// The value associated with the key.
	Value *string
}

// Contains additional parameters for the connection.
type ConnectionHttpParameters struct {

	// Contains additional body string parameters for the connection.
	BodyParameters []ConnectionBodyParameter

	// Contains additional header parameters for the connection.
	HeaderParameters []ConnectionHeaderParameter

	// Contains additional query string parameters for the connection.
	QueryStringParameters []ConnectionQueryStringParameter
}

// Contains the client response parameters for the connection when OAuth is
// specified as the authorization type.
type ConnectionOAuthClientResponseParameters struct {

	// The client ID associated with the response to the connection request.
	ClientID *string
}

// Contains the response parameters when OAuth is specified as the authorization
// type.
type ConnectionOAuthResponseParameters struct {

	// The URL to the HTTP endpoint that authorized the request.
	AuthorizationEndpoint *string

	// A ConnectionOAuthClientResponseParameters object that contains details about the
	// client parameters returned when OAuth is specified as the authorization type.
	ClientParameters *ConnectionOAuthClientResponseParameters

	// The method used to connect to the HTTP endpoint.
	HttpMethod ConnectionOAuthHttpMethod

	// The additional HTTP parameters used for the OAuth authorization request.
	OAuthHttpParameters *ConnectionHttpParameters
}

// Additional query string parameter for the connection. You can include up to 100
// additional query string parameters per request. Each additional parameter counts
// towards the event payload size, which cannot exceed 64 KB.
type ConnectionQueryStringParameter struct {

	// Specifies whether the value is secret.
	IsValueSecret bool

	// The key for a query string parameter.
	Key *string

	// The value associated with the key for the query string parameter.
	Value *string
}

// Contains the API key authorization parameters for the connection.
type CreateConnectionApiKeyAuthRequestParameters struct {

	// The name of the API key to use for authorization.
	//
	// This member is required.
	ApiKeyName *string

	// The value for the API key to use for authorization.
	//
	// This member is required.
	ApiKeyValue *string
}

// Contains the authorization parameters for the connection.
type CreateConnectionAuthRequestParameters struct {

	// A CreateConnectionApiKeyAuthRequestParameters object that contains the API key
	// authorization parameters to use for the connection.
	ApiKeyAuthParameters *CreateConnectionApiKeyAuthRequestParameters

	// A CreateConnectionBasicAuthRequestParameters object that contains the Basic
	// authorization parameters to use for the connection.
	BasicAuthParameters *CreateConnectionBasicAuthRequestParameters

	// A ConnectionHttpParameters object that contains the API key authorization
	// parameters to use for the connection. Note that if you include additional
	// parameters for the target of a rule via HttpParameters, including query strings,
	// the parameters added for the connection take precedence.
	InvocationHttpParameters *ConnectionHttpParameters

	// A CreateConnectionOAuthRequestParameters object that contains the OAuth
	// authorization parameters to use for the connection.
	OAuthParameters *CreateConnectionOAuthRequestParameters
}

// Contains the Basic authorization parameters to use for the connection.
type CreateConnectionBasicAuthRequestParameters struct {

	// The password associated with the user name to use for Basic authorization.
	//
	// This member is required.
	Password *string

	// The user name to use for Basic authorization.
	//
	// This member is required.
	Username *string
}

// Contains the Basic authorization parameters to use for the connection.
type CreateConnectionOAuthClientRequestParameters struct {

	// The client ID to use for OAuth authorization for the connection.
	//
	// This member is required.
	ClientID *string

	// The client secret associated with the client ID to use for OAuth authorization
	// for the connection.
	//
	// This member is required.
	ClientSecret *string
}

// Contains the OAuth authorization parameters to use for the connection.
type CreateConnectionOAuthRequestParameters struct {

	// The URL to the authorization endpoint when OAuth is specified as the
	// authorization type.
	//
	// This member is required.
	AuthorizationEndpoint *string

	// A CreateConnectionOAuthClientRequestParameters object that contains the client
	// parameters for OAuth authorization.
	//
	// This member is required.
	ClientParameters *CreateConnectionOAuthClientRequestParameters

	// The method to use for the authorization request.
	//
	// This member is required.
	HttpMethod ConnectionOAuthHttpMethod

	// A ConnectionHttpParameters object that contains details about the additional
	// parameters to use for the connection.
	OAuthHttpParameters *ConnectionHttpParameters
}

// A DeadLetterConfig object that contains information about a dead-letter queue
// configuration.
type DeadLetterConfig struct {

	// The ARN of the SQS queue specified as the target for the dead-letter queue.
	Arn *string
}

// The custom parameters to be used when the target is an Amazon ECS task.
type EcsParameters struct {

	// The ARN of the task definition to use if the event target is an Amazon ECS task.
	//
	// This member is required.
	TaskDefinitionArn *string

	// Specifies an ECS task group for the task. The maximum length is 255 characters.
	Group *string

	// Specifies the launch type on which your task is running. The launch type that
	// you specify here must match one of the launch type (compatibilities) of the
	// target task. The FARGATE value is supported only in the Regions where AWS
	// Fargate with Amazon ECS is supported. For more information, see AWS Fargate on
	// Amazon ECS
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS-Fargate.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LaunchType LaunchType

	// Use this structure if the ECS task uses the awsvpc network mode. This structure
	// specifies the VPC subnets and security groups associated with the task, and
	// whether a public IP address is to be used. This structure is required if
	// LaunchType is FARGATE because the awsvpc mode is required for Fargate tasks. If
	// you specify NetworkConfiguration when the target ECS task does not use the
	// awsvpc network mode, the task fails.
	NetworkConfiguration *NetworkConfiguration

	// Specifies the platform version for the task. Specify only the numeric portion of
	// the platform version, such as 1.1.0. This structure is used only if LaunchType
	// is FARGATE. For more information about valid platform versions, see AWS Fargate
	// Platform Versions
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string

	// The number of tasks to create based on TaskDefinition. The default is 1.
	TaskCount *int32
}

// An event bus receives events from a source and routes them to rules associated
// with that event bus. Your account's default event bus receives rules from AWS
// services. A custom event bus can receive rules from AWS services as well as your
// custom applications and services. A partner event bus receives events from an
// event source created by an SaaS partner. These events come from the partners
// services or applications.
type EventBus struct {

	// The ARN of the event bus.
	Arn *string

	// The name of the event bus.
	Name *string

	// The permissions policy of the event bus, describing which other AWS accounts can
	// write events to this event bus.
	Policy *string
}

// A partner event source is created by an SaaS partner. If a customer creates a
// partner event bus that matches this event source, that AWS account can receive
// events from the partner's applications or services.
type EventSource struct {

	// The ARN of the event source.
	Arn *string

	// The name of the partner that created the event source.
	CreatedBy *string

	// The date and time the event source was created.
	CreationTime *time.Time

	// The date and time that the event source will expire, if the AWS account doesn't
	// create a matching event bus for it.
	ExpirationTime *time.Time

	// The name of the event source.
	Name *string

	// The state of the event source. If it is ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If it is
	// PENDING, either you haven't yet created a matching event bus, or that event bus
	// is deactivated. If it is DELETED, you have created a matching event bus, but the
	// event source has since been deleted.
	State EventSourceState
}

// These are custom parameter to be used when the target is an API Gateway REST
// APIs or EventBridge ApiDestinations. In the latter case, these are merged with
// any InvocationParameters specified on the Connection, with any values from the
// Connection taking precedence.
type HttpParameters struct {

	// The headers that need to be sent as part of request invoking the API Gateway
	// REST API or EventBridge ApiDestination.
	HeaderParameters map[string]string

	// The path parameter values to be used to populate API Gateway REST API or
	// EventBridge ApiDestination path wildcards ("*").
	PathParameterValues []string

	// The query string keys/values that need to be sent as part of request invoking
	// the API Gateway REST API or EventBridge ApiDestination.
	QueryStringParameters map[string]string
}

// Contains the parameters needed for you to provide custom input to a target based
// on one or more pieces of data extracted from the event.
type InputTransformer struct {

	// Input template where you specify placeholders that will be filled with the
	// values of the keys from InputPathsMap to customize the data sent to the target.
	// Enclose each InputPathsMaps value in brackets: <value> The InputTemplate must be
	// valid JSON. If InputTemplate is a JSON object (surrounded by curly braces), the
	// following restrictions apply:
	//
	// * The placeholder cannot be used as an object
	// key.
	//
	// The following example shows the syntax for using InputPathsMap and
	// InputTemplate.  "InputTransformer":
	//     {
	//
	//     "InputPathsMap": {"instance":
	// "$.detail.instance","status": "$.detail.status"},
	//
	//     "InputTemplate": " is in
	// state "
	//
	// } To have the InputTemplate include quote marks within a JSON string,
	// escape each quote marks with a slash, as in the following example:
	// "InputTransformer":
	//     {
	//
	//     "InputPathsMap": {"instance":
	// "$.detail.instance","status": "$.detail.status"},
	//
	//     "InputTemplate": " is in
	// state """
	//
	// } The InputTemplate can also be valid JSON with varibles in quotes or
	// out, as in the following example:  "InputTransformer":
	//     {
	//
	//
	// "InputPathsMap": {"instance": "$.detail.instance","status":
	// "$.detail.status"},
	//
	//     "InputTemplate": '{"myInstance": ,"myStatus": " is in
	// state """}'
	//
	//     }
	//
	// This member is required.
	InputTemplate *string

	// Map of JSON paths to be extracted from the event. You can then insert these in
	// the template in InputTemplate to produce the output you want to be sent to the
	// target. InputPathsMap is an array key-value pairs, where each value is a valid
	// JSON path. You can have as many as 100 key-value pairs. You must use JSON dot
	// notation, not bracket notation. The keys cannot start with "AWS."
	InputPathsMap map[string]string
}

// This object enables you to specify a JSON path to extract from the event and use
// as the partition key for the Amazon Kinesis data stream, so that you can control
// the shard to which the event goes. If you do not include this parameter, the
// default is to use the eventId as the partition key.
type KinesisParameters struct {

	// The JSON path to be extracted from the event and used as the partition key. For
	// more information, see Amazon Kinesis Streams Key Concepts
	// (https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html#partition-key)
	// in the Amazon Kinesis Streams Developer Guide.
	//
	// This member is required.
	PartitionKeyPath *string
}

// This structure specifies the network configuration for an ECS task.
type NetworkConfiguration struct {

	// Use this structure to specify the VPC subnets and security groups for the task,
	// and whether a public IP address is to be used. This structure is relevant only
	// for ECS tasks that use the awsvpc network mode.
	AwsvpcConfiguration *AwsVpcConfiguration
}

// A partner event source is created by an SaaS partner. If a customer creates a
// partner event bus that matches this event source, that AWS account can receive
// events from the partner's applications or services.
type PartnerEventSource struct {

	// The ARN of the partner event source.
	Arn *string

	// The name of the partner event source.
	Name *string
}

// The AWS account that a partner event source has been offered to.
type PartnerEventSourceAccount struct {

	// The AWS account ID that the partner event source was offered to.
	Account *string

	// The date and time the event source was created.
	CreationTime *time.Time

	// The date and time that the event source will expire, if the AWS account doesn't
	// create a matching event bus for it.
	ExpirationTime *time.Time

	// The state of the event source. If it is ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If it is
	// PENDING, either you haven't yet created a matching event bus, or that event bus
	// is deactivated. If it is DELETED, you have created a matching event bus, but the
	// event source has since been deleted.
	State EventSourceState
}

// Represents an event to be submitted.
type PutEventsRequestEntry struct {

	// A valid JSON string. There is no other schema imposed. The JSON string may
	// contain fields and nested subobjects.
	Detail *string

	// Free-form string used to decide what fields to expect in the event detail.
	DetailType *string

	// The name or ARN of the event bus to receive the event. Only the rules that are
	// associated with this event bus are used to match the event. If you omit this,
	// the default event bus is used.
	EventBusName *string

	// AWS resources, identified by Amazon Resource Name (ARN), which the event
	// primarily concerns. Any number, including zero, may be present.
	Resources []string

	// The source of the event.
	Source *string

	// The time stamp of the event, per RFC3339
	// (https://www.rfc-editor.org/rfc/rfc3339.txt). If no time stamp is provided, the
	// time stamp of the PutEvents call is used.
	Time *time.Time

	// An AWS X-Ray trade header, which is an http header (X-Amzn-Trace-Id) that
	// contains the trace-id associated with the event. To learn more about X-Ray trace
	// headers, see Tracing header
	// (https://docs.aws.amazon.com/xray/latest/devguide/xray-concepts.html#xray-concepts-tracingheader)
	// in the AWS X-Ray Developer Guide.
	TraceHeader *string
}

// Represents an event that failed to be submitted.
type PutEventsResultEntry struct {

	// The error code that indicates why the event submission failed.
	ErrorCode *string

	// The error message that explains why the event submission failed.
	ErrorMessage *string

	// The ID of the event.
	EventId *string
}

// The details about an event generated by an SaaS partner.
type PutPartnerEventsRequestEntry struct {

	// A valid JSON string. There is no other schema imposed. The JSON string may
	// contain fields and nested subobjects.
	Detail *string

	// A free-form string used to decide what fields to expect in the event detail.
	DetailType *string

	// AWS resources, identified by Amazon Resource Name (ARN), which the event
	// primarily concerns. Any number, including zero, may be present.
	Resources []string

	// The event source that is generating the evntry.
	Source *string

	// The date and time of the event.
	Time *time.Time
}

// Represents an event that a partner tried to generate, but failed.
type PutPartnerEventsResultEntry struct {

	// The error code that indicates why the event submission failed.
	ErrorCode *string

	// The error message that explains why the event submission failed.
	ErrorMessage *string

	// The ID of the event.
	EventId *string
}

// Represents a target that failed to be added to a rule.
type PutTargetsResultEntry struct {

	// The error code that indicates why the target addition failed. If the value is
	// ConcurrentModificationException, too many requests were made at the same time.
	ErrorCode *string

	// The error message that explains why the target addition failed.
	ErrorMessage *string

	// The ID of the target.
	TargetId *string
}

// These are custom parameters to be used when the target is a Redshift cluster to
// invoke the Redshift Data API ExecuteStatement based on EventBridge events.
type RedshiftDataParameters struct {

	// The name of the database. Required when authenticating using temporary
	// credentials.
	//
	// This member is required.
	Database *string

	// The SQL statement text to run.
	//
	// This member is required.
	Sql *string

	// The database user name. Required when authenticating using temporary
	// credentials.
	DbUser *string

	// The name or ARN of the secret that enables access to the database. Required when
	// authenticating using AWS Secrets Manager.
	SecretManagerArn *string

	// The name of the SQL statement. You can name the SQL statement when you create it
	// to identify the query.
	StatementName *string

	// Indicates whether to send an event back to EventBridge after the SQL statement
	// runs.
	WithEvent bool
}

// Represents a target that failed to be removed from a rule.
type RemoveTargetsResultEntry struct {

	// The error code that indicates why the target removal failed. If the value is
	// ConcurrentModificationException, too many requests were made at the same time.
	ErrorCode *string

	// The error message that explains why the target removal failed.
	ErrorMessage *string

	// The ID of the target.
	TargetId *string
}

// A Replay object that contains details about a replay.
type Replay struct {

	// A time stamp for the time to start replaying events. Any event with a creation
	// time prior to the EventEndTime specified is replayed.
	EventEndTime *time.Time

	// A time stamp for the time that the last event was replayed.
	EventLastReplayedTime *time.Time

	// The ARN of the archive to replay event from.
	EventSourceArn *string

	// A time stamp for the time to start replaying events. This is determined by the
	// time in the event as described in Time
	// (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEventsRequestEntry.html#eventbridge-Type-PutEventsRequestEntry-Time).
	EventStartTime *time.Time

	// A time stamp for the time that the replay completed.
	ReplayEndTime *time.Time

	// The name of the replay.
	ReplayName *string

	// A time stamp for the time that the replay started.
	ReplayStartTime *time.Time

	// The current state of the replay.
	State ReplayState

	// A description of why the replay is in the current state.
	StateReason *string
}

// A ReplayDestination object that contains details about a replay.
type ReplayDestination struct {

	// The ARN of the event bus to replay event to. You can replay events only to the
	// event bus specified to create the archive.
	//
	// This member is required.
	Arn *string

	// A list of ARNs for rules to replay events to.
	FilterArns []string
}

// A RetryPolicy object that includes information about the retry policy settings.
type RetryPolicy struct {

	// The maximum amount of time, in seconds, to continue to make retry attempts.
	MaximumEventAgeInSeconds *int32

	// The maximum number of retry attempts to make before the request fails. Retry
	// attempts continue until either the maximum number of attempts is made or until
	// the duration of the MaximumEventAgeInSeconds is met.
	MaximumRetryAttempts *int32
}

// Contains information about a rule in Amazon EventBridge.
type Rule struct {

	// The Amazon Resource Name (ARN) of the rule.
	Arn *string

	// The description of the rule.
	Description *string

	// The name or ARN of the event bus associated with the rule. If you omit this, the
	// default event bus is used.
	EventBusName *string

	// The event pattern of the rule. For more information, see Events and Event
	// Patterns
	// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	EventPattern *string

	// If the rule was created on behalf of your account by an AWS service, this field
	// displays the principal name of the service that created the rule.
	ManagedBy *string

	// The name of the rule.
	Name *string

	// The Amazon Resource Name (ARN) of the role that is used for target invocation.
	RoleArn *string

	// The scheduling expression. For example, "cron(0 20 * * ? *)", "rate(5 minutes)".
	ScheduleExpression *string

	// The state of the rule.
	State RuleState
}

// This parameter contains the criteria (either InstanceIds or a tag) used to
// specify which EC2 instances are to be sent the command.
type RunCommandParameters struct {

	// Currently, we support including only one RunCommandTarget block, which specifies
	// either an array of InstanceIds or a tag.
	//
	// This member is required.
	RunCommandTargets []RunCommandTarget
}

// Information about the EC2 instances that are to be sent the command, specified
// as key-value pairs. Each RunCommandTarget block can include only one key, but
// this key may specify multiple values.
type RunCommandTarget struct {

	// Can be either tag: tag-key or InstanceIds.
	//
	// This member is required.
	Key *string

	// If Key is tag: tag-key, Values is a list of tag values. If Key is InstanceIds,
	// Values is a list of Amazon EC2 instance IDs.
	//
	// This member is required.
	Values []string
}

// Name/Value pair of a parameter to start execution of a SageMaker Model Building
// Pipeline.
type SageMakerPipelineParameter struct {

	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	//
	// This member is required.
	Name *string

	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	//
	// This member is required.
	Value *string
}

// These are custom parameters to use when the target is a SageMaker Model Building
// Pipeline that starts based on EventBridge events.
type SageMakerPipelineParameters struct {

	// List of Parameter names and values for SageMaker Model Building Pipeline
	// execution.
	PipelineParameterList []SageMakerPipelineParameter
}

// This structure includes the custom parameter to be used when the target is an
// SQS FIFO queue.
type SqsParameters struct {

	// The FIFO message group ID to use as the target.
	MessageGroupId *string
}

// A key-value pair associated with an AWS resource. In EventBridge, rules and
// event buses support tagging.
type Tag struct {

	// A string you can use to assign a value. The combination of tag keys and values
	// can help you organize and categorize your resources.
	//
	// This member is required.
	Key *string

	// The value for the specified tag key.
	//
	// This member is required.
	Value *string
}

// Targets are the resources to be invoked when a rule is triggered. For a complete
// list of services and resources that can be set as a target, see PutTargets. If
// you are setting the event bus of another account as the target, and that account
// granted permission to your account through an organization instead of directly
// by the account ID, then you must specify a RoleArn with proper permissions in
// the Target structure. For more information, see Sending and Receiving Events
// Between AWS Accounts
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html)
// in the Amazon EventBridge User Guide.
type Target struct {

	// The Amazon Resource Name (ARN) of the target.
	//
	// This member is required.
	Arn *string

	// The ID of the target.
	//
	// This member is required.
	Id *string

	// If the event target is an AWS Batch job, this contains the job definition, job
	// name, and other parameters. For more information, see Jobs
	// (https://docs.aws.amazon.com/batch/latest/userguide/jobs.html) in the AWS Batch
	// User Guide.
	BatchParameters *BatchParameters

	// The DeadLetterConfig that defines the target queue to send dead-letter queue
	// events to.
	DeadLetterConfig *DeadLetterConfig

	// Contains the Amazon ECS task definition and task count to be used, if the event
	// target is an Amazon ECS task. For more information about Amazon ECS tasks, see
	// Task Definitions
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	EcsParameters *EcsParameters

	// Contains the HTTP parameters to use when the target is a API Gateway REST
	// endpoint or EventBridge ApiDestination. If you specify an API Gateway REST API
	// or EventBridge ApiDestination as a target, you can use this parameter to specify
	// headers, path parameters, and query string keys/values as part of your target
	// invoking request. If you're using ApiDestinations, the corresponding Connection
	// can also have these values configured. In case of any conflicting keys, values
	// from the Connection take precedence.
	HttpParameters *HttpParameters

	// Valid JSON text passed to the target. In this case, nothing from the event
	// itself is passed to the target. For more information, see The JavaScript Object
	// Notation (JSON) Data Interchange Format
	// (http://www.rfc-editor.org/rfc/rfc7159.txt).
	Input *string

	// The value of the JSONPath that is used for extracting part of the matched event
	// when passing it to the target. You must use JSON dot notation, not bracket
	// notation. For more information about JSON paths, see JSONPath
	// (http://goessner.net/articles/JsonPath/).
	InputPath *string

	// Settings to enable you to provide custom input to a target based on certain
	// event data. You can extract one or more key-value pairs from the event and then
	// use that data to send customized input to the target.
	InputTransformer *InputTransformer

	// The custom parameter you can use to control the shard assignment, when the
	// target is a Kinesis data stream. If you do not include this parameter, the
	// default is to use the eventId as the partition key.
	KinesisParameters *KinesisParameters

	// Contains the Redshift Data API parameters to use when the target is a Redshift
	// cluster. If you specify a Redshift Cluster as a Target, you can use this to
	// specify parameters to invoke the Redshift Data API ExecuteStatement based on
	// EventBridge events.
	RedshiftDataParameters *RedshiftDataParameters

	// The RetryPolicy object that contains the retry policy configuration to use for
	// the dead-letter queue.
	RetryPolicy *RetryPolicy

	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when
	// the rule is triggered. If one rule triggers multiple targets, you can use a
	// different IAM role for each target.
	RoleArn *string

	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command.
	RunCommandParameters *RunCommandParameters

	// Contains the SageMaker Model Building Pipeline parameters to start execution of
	// a SageMaker Model Building Pipeline. If you specify a SageMaker Model Building
	// Pipeline as a target, you can use this to specify parameters to start a pipeline
	// execution based on EventBridge events.
	SageMakerPipelineParameters *SageMakerPipelineParameters

	// Contains the message group ID to use when the target is a FIFO queue. If you
	// specify an SQS FIFO queue as a target, the queue must have content-based
	// deduplication enabled.
	SqsParameters *SqsParameters
}

// Contains the API key authorization parameters to use to update the connection.
type UpdateConnectionApiKeyAuthRequestParameters struct {

	// The name of the API key to use for authorization.
	ApiKeyName *string

	// The value associated with teh API key to use for authorization.
	ApiKeyValue *string
}

// Contains the additional parameters to use for the connection.
type UpdateConnectionAuthRequestParameters struct {

	// A UpdateConnectionApiKeyAuthRequestParameters object that contains the
	// authorization parameters for API key authorization.
	ApiKeyAuthParameters *UpdateConnectionApiKeyAuthRequestParameters

	// A UpdateConnectionBasicAuthRequestParameters object that contains the
	// authorization parameters for Basic authorization.
	BasicAuthParameters *UpdateConnectionBasicAuthRequestParameters

	// A ConnectionHttpParameters object that contains the additional parameters to use
	// for the connection.
	InvocationHttpParameters *ConnectionHttpParameters

	// A UpdateConnectionOAuthRequestParameters object that contains the authorization
	// parameters for OAuth authorization.
	OAuthParameters *UpdateConnectionOAuthRequestParameters
}

// Contains the Basic authorization parameters for the connection.
type UpdateConnectionBasicAuthRequestParameters struct {

	// The password associated with the user name to use for Basic authorization.
	Password *string

	// The user name to use for Basic authorization.
	Username *string
}

// Contains the OAuth authorization parameters to use for the connection.
type UpdateConnectionOAuthClientRequestParameters struct {

	// The client ID to use for OAuth authorization.
	ClientID *string

	// The client secret assciated with the client ID to use for OAuth authorization.
	ClientSecret *string
}

// Contains the OAuth request parameters to use for the connection.
type UpdateConnectionOAuthRequestParameters struct {

	// The URL to the authorization endpoint when OAuth is specified as the
	// authorization type.
	AuthorizationEndpoint *string

	// A UpdateConnectionOAuthClientRequestParameters object that contains the client
	// parameters to use for the connection when OAuth is specified as the
	// authorization type.
	ClientParameters *UpdateConnectionOAuthClientRequestParameters

	// The method used to connect to the HTTP endpoint.
	HttpMethod ConnectionOAuthHttpMethod

	// The additional HTTP parameters used for the OAuth authorization request.
	OAuthHttpParameters *ConnectionHttpParameters
}