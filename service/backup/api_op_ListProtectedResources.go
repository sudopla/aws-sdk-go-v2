// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListProtectedResourcesInput
type ListProtectedResourcesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListProtectedResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProtectedResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProtectedResourcesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProtectedResourcesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListProtectedResourcesOutput
type ListProtectedResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`

	// An array of resources successfully backed up by AWS Backup including the
	// time the resource was saved, an Amazon Resource Name (ARN) of the resource,
	// and a resource type.
	Results []ProtectedResource `type:"list"`
}

// String returns the string representation
func (s ListProtectedResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProtectedResourcesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Results != nil {
		v := s.Results

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Results", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListProtectedResources = "ListProtectedResources"

// ListProtectedResourcesRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns an array of resources successfully backed up by AWS Backup, including
// the time the resource was saved, an Amazon Resource Name (ARN) of the resource,
// and a resource type.
//
//    // Example sending a request using ListProtectedResourcesRequest.
//    req := client.ListProtectedResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListProtectedResources
func (c *Client) ListProtectedResourcesRequest(input *ListProtectedResourcesInput) ListProtectedResourcesRequest {
	op := &aws.Operation{
		Name:       opListProtectedResources,
		HTTPMethod: "GET",
		HTTPPath:   "/resources/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProtectedResourcesInput{}
	}

	req := c.newRequest(op, input, &ListProtectedResourcesOutput{})
	return ListProtectedResourcesRequest{Request: req, Input: input, Copy: c.ListProtectedResourcesRequest}
}

// ListProtectedResourcesRequest is the request type for the
// ListProtectedResources API operation.
type ListProtectedResourcesRequest struct {
	*aws.Request
	Input *ListProtectedResourcesInput
	Copy  func(*ListProtectedResourcesInput) ListProtectedResourcesRequest
}

// Send marshals and sends the ListProtectedResources API request.
func (r ListProtectedResourcesRequest) Send(ctx context.Context) (*ListProtectedResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProtectedResourcesResponse{
		ListProtectedResourcesOutput: r.Request.Data.(*ListProtectedResourcesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProtectedResourcesRequestPaginator returns a paginator for ListProtectedResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProtectedResourcesRequest(input)
//   p := backup.NewListProtectedResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProtectedResourcesPaginator(req ListProtectedResourcesRequest) ListProtectedResourcesPaginator {
	return ListProtectedResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProtectedResourcesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListProtectedResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProtectedResourcesPaginator struct {
	aws.Pager
}

func (p *ListProtectedResourcesPaginator) CurrentPage() *ListProtectedResourcesOutput {
	return p.Pager.CurrentPage().(*ListProtectedResourcesOutput)
}

// ListProtectedResourcesResponse is the response type for the
// ListProtectedResources API operation.
type ListProtectedResourcesResponse struct {
	*ListProtectedResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProtectedResources request.
func (r *ListProtectedResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}