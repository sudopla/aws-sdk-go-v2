// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeleteSystemTemplateRequest
type DeleteSystemTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the system to be deleted.
	//
	// The ID should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSystemTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSystemTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSystemTemplateInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeleteSystemTemplateResponse
type DeleteSystemTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSystemTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSystemTemplate = "DeleteSystemTemplate"

// DeleteSystemTemplateRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Deletes a system. New deployments can't contain the system after its deletion.
// Existing deployments that contain the system will continue to work because
// they use a snapshot of the system that is taken when it is deployed.
//
//    // Example sending a request using DeleteSystemTemplateRequest.
//    req := client.DeleteSystemTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeleteSystemTemplate
func (c *Client) DeleteSystemTemplateRequest(input *DeleteSystemTemplateInput) DeleteSystemTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteSystemTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSystemTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteSystemTemplateOutput{})
	return DeleteSystemTemplateRequest{Request: req, Input: input, Copy: c.DeleteSystemTemplateRequest}
}

// DeleteSystemTemplateRequest is the request type for the
// DeleteSystemTemplate API operation.
type DeleteSystemTemplateRequest struct {
	*aws.Request
	Input *DeleteSystemTemplateInput
	Copy  func(*DeleteSystemTemplateInput) DeleteSystemTemplateRequest
}

// Send marshals and sends the DeleteSystemTemplate API request.
func (r DeleteSystemTemplateRequest) Send(ctx context.Context) (*DeleteSystemTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSystemTemplateResponse{
		DeleteSystemTemplateOutput: r.Request.Data.(*DeleteSystemTemplateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSystemTemplateResponse is the response type for the
// DeleteSystemTemplate API operation.
type DeleteSystemTemplateResponse struct {
	*DeleteSystemTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSystemTemplate request.
func (r *DeleteSystemTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}