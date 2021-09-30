// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified property values in the resource. You specify your resource
// property updates as a list of patch operations contained in a JSON patch
// document that adheres to the  RFC 6902 - JavaScript Object Notation (JSON) Patch
// (https://datatracker.ietf.org/doc/html/rfc6902) standard. For details on how
// Cloud Control API performs resource update operations, see Updating a resource
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html)
// in the Amazon Web Services Cloud Control API User Guide. After you have
// initiated a resource update request, you can monitor the progress of your
// request by calling GetResourceRequestStatus
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
// using the RequestToken of the ProgressEvent returned by UpdateResource. For more
// information about the properties of a specific resource, refer to the related
// topic for the resource in the Resource and property types reference
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
// in the Amazon Web Services CloudFormation Users Guide.
func (c *Client) UpdateResource(ctx context.Context, params *UpdateResourceInput, optFns ...func(*Options)) (*UpdateResourceOutput, error) {
	if params == nil {
		params = &UpdateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResource", params, optFns, c.addOperationUpdateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceInput struct {

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

	// A JavaScript Object Notation (JSON) document listing the patch operations that
	// represent the updates to apply to the current resource properties. For details,
	// see Composing the patch document
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html#resource-operations-update-patch)
	// in the Amazon Web Services Cloud Control API User Guide.
	//
	// This member is required.
	PatchDocument *string

	// The name of the resource type.
	//
	// This member is required.
	TypeName *string

	// A unique identifier to ensure the idempotency of the resource request. As a best
	// practice, specify this token to ensure idempotency, so that Amazon Web Services
	// Cloud Control API can accurately distinguish between request retries and new
	// resource requests. You might retry a resource request to ensure that it was
	// successfully received. A client token is valid for 36 hours once used. After
	// that, a resource request with the same client token is treated as a new request.
	// If you do not specify a client token, one is generated for inclusion in the
	// request. For more information, see Ensuring resource operation requests are
	// unique
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-idempotency)
	// in the Amazon Web Services Cloud Control API User Guide.
	ClientToken *string

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

type UpdateResourceOutput struct {

	// Represents the current status of the resource update request. Use the
	// RequestToken of the ProgressEvent with GetResourceRequestStatus
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// to return the current status of a resource operation request.
	ProgressEvent *types.ProgressEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateResource{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateResourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateResource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateResource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateResourceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateResourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateResource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudcontrolapi",
		OperationName: "UpdateResource",
	}
}
