// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Fetches the execution history of the flow.
func (c *Client) DescribeFlowExecutionRecords(ctx context.Context, params *DescribeFlowExecutionRecordsInput, optFns ...func(*Options)) (*DescribeFlowExecutionRecordsOutput, error) {
	if params == nil {
		params = &DescribeFlowExecutionRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFlowExecutionRecords", params, optFns, addOperationDescribeFlowExecutionRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFlowExecutionRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowExecutionRecordsInput struct {

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string

	// Specifies the maximum number of items that should be returned in the result set.
	// The default for maxResults is 20 (for all paginated API operations).
	MaxResults *int32

	// The pagination token for the next page of data.
	NextToken *string
}

type DescribeFlowExecutionRecordsOutput struct {

	// Returns a list of all instances when this flow was run.
	FlowExecutions []*types.ExecutionRecord

	// The pagination token for the next page of data.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFlowExecutionRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFlowExecutionRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFlowExecutionRecords{}, middleware.After)
	if err != nil {
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpDescribeFlowExecutionRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlowExecutionRecords(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeFlowExecutionRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeFlowExecutionRecords",
	}
}