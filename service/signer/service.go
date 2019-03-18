// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Signer provides the API operation methods for making requests to
// AWS Signer. See this package's package overview docs
// for details on the service.
//
// Signer methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Signer struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Signer)

// Used for custom request initialization logic
var initRequest func(*Signer, *aws.Request)

// Service information constants
const (
	ServiceName = "signer"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Signer client with a config.
//
// Example:
//     // Create a Signer client from just a config.
//     svc := signer.New(myConfig)
func New(config aws.Config) *Signer {
	var signingName string
	signingName = "signer"
	signingRegion := config.Region

	svc := &Signer{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-08-25",
				JSONVersion:   "1.1",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Signer operation and runs any
// custom request initialization.
func (c *Signer) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}