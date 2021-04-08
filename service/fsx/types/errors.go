// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An Active Directory error.
type ActiveDirectoryError struct {
	Message *string

	ActiveDirectoryId *string
	Type              ActiveDirectoryErrorType
}

func (e *ActiveDirectoryError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ActiveDirectoryError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ActiveDirectoryError) ErrorCode() string             { return "ActiveDirectoryError" }
func (e *ActiveDirectoryError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Another backup is already under way. Wait for completion before initiating
// additional backups of this file system.
type BackupInProgress struct {
	Message *string
}

func (e *BackupInProgress) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupInProgress) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupInProgress) ErrorCode() string             { return "BackupInProgress" }
func (e *BackupInProgress) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No Amazon FSx backups were found based upon the supplied parameters.
type BackupNotFound struct {
	Message *string
}

func (e *BackupNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupNotFound) ErrorCode() string             { return "BackupNotFound" }
func (e *BackupNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can't delete a backup while it's being used to restore a file system.
type BackupRestoring struct {
	Message *string

	FileSystemId *string
}

func (e *BackupRestoring) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupRestoring) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupRestoring) ErrorCode() string             { return "BackupRestoring" }
func (e *BackupRestoring) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A generic error indicating a failure with a client request.
type BadRequest struct {
	Message *string
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequest) ErrorCode() string             { return "BadRequest" }
func (e *BadRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The data repository task could not be canceled because the task has already
// ended.
type DataRepositoryTaskEnded struct {
	Message *string
}

func (e *DataRepositoryTaskEnded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryTaskEnded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryTaskEnded) ErrorCode() string             { return "DataRepositoryTaskEnded" }
func (e *DataRepositoryTaskEnded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An existing data repository task is currently executing on the file system. Wait
// until the existing task has completed, then create the new task.
type DataRepositoryTaskExecuting struct {
	Message *string
}

func (e *DataRepositoryTaskExecuting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryTaskExecuting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryTaskExecuting) ErrorCode() string             { return "DataRepositoryTaskExecuting" }
func (e *DataRepositoryTaskExecuting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The data repository task or tasks you specified could not be found.
type DataRepositoryTaskNotFound struct {
	Message *string
}

func (e *DataRepositoryTaskNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataRepositoryTaskNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataRepositoryTaskNotFound) ErrorCode() string             { return "DataRepositoryTaskNotFound" }
func (e *DataRepositoryTaskNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No Amazon FSx file systems were found based upon supplied parameters.
type FileSystemNotFound struct {
	Message *string
}

func (e *FileSystemNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FileSystemNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FileSystemNotFound) ErrorCode() string             { return "FileSystemNotFound" }
func (e *FileSystemNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The error returned when a second request is received with the same client
// request token but different parameters settings. A client request token should
// always uniquely identify a single request.
type IncompatibleParameterError struct {
	Message *string

	Parameter *string
}

func (e *IncompatibleParameterError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleParameterError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleParameterError) ErrorCode() string             { return "IncompatibleParameterError" }
func (e *IncompatibleParameterError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A generic error indicating a server-side failure.
type InternalServerError struct {
	Message *string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The path provided for data repository export isn't valid.
type InvalidExportPath struct {
	Message *string
}

func (e *InvalidExportPath) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidExportPath) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidExportPath) ErrorCode() string             { return "InvalidExportPath" }
func (e *InvalidExportPath) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The path provided for data repository import isn't valid.
type InvalidImportPath struct {
	Message *string
}

func (e *InvalidImportPath) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidImportPath) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidImportPath) ErrorCode() string             { return "InvalidImportPath" }
func (e *InvalidImportPath) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more network settings specified in the request are invalid. InvalidVpcId
// means that the ID passed for the virtual private cloud (VPC) is invalid.
// InvalidSubnetIds returns the list of IDs for subnets that are either invalid or
// not part of the VPC specified. InvalidSecurityGroupIds returns the list of IDs
// for security groups that are either invalid or not part of the VPC specified.
type InvalidNetworkSettings struct {
	Message *string

	InvalidSubnetId        *string
	InvalidSecurityGroupId *string
}

func (e *InvalidNetworkSettings) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNetworkSettings) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNetworkSettings) ErrorCode() string             { return "InvalidNetworkSettings" }
func (e *InvalidNetworkSettings) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An invalid value for PerUnitStorageThroughput was provided. Please create your
// file system again, using a valid value.
type InvalidPerUnitStorageThroughput struct {
	Message *string
}

func (e *InvalidPerUnitStorageThroughput) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPerUnitStorageThroughput) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPerUnitStorageThroughput) ErrorCode() string {
	return "InvalidPerUnitStorageThroughput"
}
func (e *InvalidPerUnitStorageThroughput) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A file system configuration is required for this operation.
type MissingFileSystemConfiguration struct {
	Message *string
}

func (e *MissingFileSystemConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingFileSystemConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingFileSystemConfiguration) ErrorCode() string             { return "MissingFileSystemConfiguration" }
func (e *MissingFileSystemConfiguration) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified for the tagging operation is not a resource type owned by
// Amazon FSx. Use the API of the relevant service to perform the operation.
type NotServiceResourceError struct {
	Message *string

	ResourceARN *string
}

func (e *NotServiceResourceError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotServiceResourceError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotServiceResourceError) ErrorCode() string             { return "NotServiceResourceError" }
func (e *NotServiceResourceError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified does not support tagging.
type ResourceDoesNotSupportTagging struct {
	Message *string

	ResourceARN *string
}

func (e *ResourceDoesNotSupportTagging) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceDoesNotSupportTagging) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceDoesNotSupportTagging) ErrorCode() string             { return "ResourceDoesNotSupportTagging" }
func (e *ResourceDoesNotSupportTagging) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified by the Amazon Resource Name (ARN) can't be found.
type ResourceNotFound struct {
	Message *string

	ResourceARN *string
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string             { return "ResourceNotFound" }
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error indicating that a particular service limit was exceeded. You can
// increase some service limits by contacting AWS Support.
type ServiceLimitExceeded struct {
	Message *string

	Limit ServiceLimit
}

func (e *ServiceLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLimitExceeded) ErrorCode() string             { return "ServiceLimitExceeded" }
func (e *ServiceLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested operation is not supported for this resource or API.
type UnsupportedOperation struct {
	Message *string
}

func (e *UnsupportedOperation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperation) ErrorCode() string             { return "UnsupportedOperation" }
func (e *UnsupportedOperation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }