// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to delete a specified traffic policy version.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteTrafficPolicyRequest
type DeleteTrafficPolicyInput struct {
	_ struct{} `type:"structure"`

	// The ID of the traffic policy that you want to delete.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`

	// The version number of the traffic policy that you want to delete.
	//
	// Version is a required field
	Version *int64 `location:"uri" locationName:"Version" min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s DeleteTrafficPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrafficPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrafficPolicyInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && *s.Version < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTrafficPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Version", protocol.Int64Value(v), metadata)
	}
	return nil
}

// An empty element.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteTrafficPolicyResponse
type DeleteTrafficPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTrafficPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTrafficPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteTrafficPolicy = "DeleteTrafficPolicy"

// DeleteTrafficPolicyRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Deletes a traffic policy.
//
//    // Example sending a request using DeleteTrafficPolicyRequest.
//    req := client.DeleteTrafficPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteTrafficPolicy
func (c *Client) DeleteTrafficPolicyRequest(input *DeleteTrafficPolicyInput) DeleteTrafficPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficPolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2013-04-01/trafficpolicy/{Id}/{Version}",
	}

	if input == nil {
		input = &DeleteTrafficPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteTrafficPolicyOutput{})
	return DeleteTrafficPolicyRequest{Request: req, Input: input, Copy: c.DeleteTrafficPolicyRequest}
}

// DeleteTrafficPolicyRequest is the request type for the
// DeleteTrafficPolicy API operation.
type DeleteTrafficPolicyRequest struct {
	*aws.Request
	Input *DeleteTrafficPolicyInput
	Copy  func(*DeleteTrafficPolicyInput) DeleteTrafficPolicyRequest
}

// Send marshals and sends the DeleteTrafficPolicy API request.
func (r DeleteTrafficPolicyRequest) Send(ctx context.Context) (*DeleteTrafficPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficPolicyResponse{
		DeleteTrafficPolicyOutput: r.Request.Data.(*DeleteTrafficPolicyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficPolicyResponse is the response type for the
// DeleteTrafficPolicy API operation.
type DeleteTrafficPolicyResponse struct {
	*DeleteTrafficPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficPolicy request.
func (r *DeleteTrafficPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}