// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the state for a routing control. A routing control is a simple on/off switch
// that you can use to route traffic to cells. When the state is On, traffic flows
// to a cell. When it's off, traffic does not flow. Before you can create a routing
// control, you first must create a cluster to host the control. For more
// information, see CreateCluster
// (https://docs.aws.amazon.com/recovery-cluster/latest/api/cluster.html). Access
// one of the endpoints for the cluster to get or update the routing control state
// to redirect traffic. For more information about working with routing controls,
// see Routing control
// (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html) in the
// Route 53 Application Recovery Controller Developer Guide.
func (c *Client) GetRoutingControlState(ctx context.Context, params *GetRoutingControlStateInput, optFns ...func(*Options)) (*GetRoutingControlStateOutput, error) {
	if params == nil {
		params = &GetRoutingControlStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRoutingControlState", params, optFns, c.addOperationGetRoutingControlStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRoutingControlStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoutingControlStateInput struct {

	// The Amazon Resource Number (ARN) for the routing control that you want to get
	// the state for.
	//
	// This member is required.
	RoutingControlArn *string

	noSmithyDocumentSerde
}

type GetRoutingControlStateOutput struct {

	// The Amazon Resource Number (ARN) of the response.
	//
	// This member is required.
	RoutingControlArn *string

	// The state of the routing control.
	//
	// This member is required.
	RoutingControlState types.RoutingControlState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRoutingControlStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRoutingControlState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRoutingControlState{}, middleware.After)
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
	if err = addOpGetRoutingControlStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoutingControlState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRoutingControlState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-cluster",
		OperationName: "GetRoutingControlState",
	}
}