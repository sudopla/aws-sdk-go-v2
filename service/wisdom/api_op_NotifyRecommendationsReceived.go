// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified recommendations from the specified assistant's queue of
// newly available recommendations. You can use this API in conjunction with
// GetRecommendations
// (https://docs.aws.amazon.com/wisdom/latest/APIReference/API_GetRecommendations.html)
// and a waitTimeSeconds input for long-polling behavior and avoiding duplicate
// recommendations.
func (c *Client) NotifyRecommendationsReceived(ctx context.Context, params *NotifyRecommendationsReceivedInput, optFns ...func(*Options)) (*NotifyRecommendationsReceivedOutput, error) {
	if params == nil {
		params = &NotifyRecommendationsReceivedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyRecommendationsReceived", params, optFns, c.addOperationNotifyRecommendationsReceivedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyRecommendationsReceivedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyRecommendationsReceivedInput struct {

	// The identifier of the Wisdom assistant. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The identifiers of the recommendations.
	//
	// This member is required.
	RecommendationIds []string

	// The identifier of the session. Can be either the ID or the ARN. URLs cannot
	// contain the ARN.
	//
	// This member is required.
	SessionId *string

	noSmithyDocumentSerde
}

type NotifyRecommendationsReceivedOutput struct {

	// The identifiers of recommendations that are causing errors.
	Errors []types.NotifyRecommendationsReceivedError

	// The identifiers of the recommendations.
	RecommendationIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyRecommendationsReceivedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpNotifyRecommendationsReceived{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpNotifyRecommendationsReceived{}, middleware.After)
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
	if err = addOpNotifyRecommendationsReceivedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyRecommendationsReceived(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNotifyRecommendationsReceived(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "NotifyRecommendationsReceived",
	}
}
