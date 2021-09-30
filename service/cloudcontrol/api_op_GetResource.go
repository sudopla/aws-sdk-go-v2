// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the current state of the specified resource. For
// details, see Reading a resource's current state
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-read.html).
// You can use this action to return information about an existing resource in your
// account and Amazon Web Services Region, whether or not those resources were
// provisioned using Cloud Control API.
func (c *Client) GetResource(ctx context.Context, params *GetResourceInput, optFns ...func(*Options)) (*GetResourceOutput, error) {
	if params == nil {
		params = &GetResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResource", params, optFns, c.addOperationGetResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceInput struct {

	// The identifier for the resource. You can specify the primary identifier, or any
	// secondary identifier defined for the resource type in its resource schema. You
	// can only specify one identifier. Primary identifiers can be specified as a
	// string or JSON; secondary identifiers must be specified as JSON. For compound
	// primary identifiers (that is, one that consists of multiple resource properties
	// strung together), to specify the primary identifier as a string, list the
	// property values in the order they are specified in the primary identifier
	// definition, separated by |. For more information, see Identifying resources
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html)
	// in the Amazon Web Services Cloud Control API User Guide.
	//
	// This member is required.
	Identifier *string

	// The name of the resource type.
	//
	// This member is required.
	TypeName *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) for
	// Cloud Control API to use when performing this resource operation. The role
	// specified must have the permissions required for this operation. The necessary
	// permissions for each event handler are defined in the handlers
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-handlers)
	// section of the resource type definition schema
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html).
	// If you do not specify a role, Cloud Control API uses a temporary session created
	// using your Amazon Web Services user credentials. For more information, see
	// Specifying credentials
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-permissions)
	// in the Amazon Web Services Cloud Control API User Guide.
	RoleArn *string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version.
	TypeVersionId *string

	noSmithyDocumentSerde
}

type GetResourceOutput struct {

	// Represents information about a provisioned resource.
	ResourceDescription *types.ResourceDescription

	// The name of the resource type.
	TypeName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetResource{}, middleware.After)
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
	if err = addOpGetResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudcontrolapi",
		OperationName: "GetResource",
	}
}
