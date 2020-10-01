// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeInternal_service_exception  ErrorCode = "InternalServiceException"
	ErrorCodeInvalid_parameter_exception ErrorCode = "InvalidParameterException"
)

type GroupByAttribute string

// Enum values for GroupByAttribute
const (
	GroupByAttributeTarget_id     GroupByAttribute = "TARGET_ID"
	GroupByAttributeRegion        GroupByAttribute = "REGION"
	GroupByAttributeResource_type GroupByAttribute = "RESOURCE_TYPE"
)

type TargetIdType string

// Enum values for TargetIdType
const (
	TargetIdTypeAccount TargetIdType = "ACCOUNT"
	TargetIdTypeOu      TargetIdType = "OU"
	TargetIdTypeRoot    TargetIdType = "ROOT"
)