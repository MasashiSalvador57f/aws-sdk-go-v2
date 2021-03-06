// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopTriggerInput struct {
	_ struct{} `type:"structure"`

	// The name of the trigger to stop.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopTriggerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopTriggerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopTriggerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopTriggerOutput struct {
	_ struct{} `type:"structure"`

	// The name of the trigger that was stopped.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StopTriggerOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopTrigger = "StopTrigger"

// StopTriggerRequest returns a request value for making API operation for
// AWS Glue.
//
// Stops a specified trigger.
//
//    // Example sending a request using StopTriggerRequest.
//    req := client.StopTriggerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StopTrigger
func (c *Client) StopTriggerRequest(input *StopTriggerInput) StopTriggerRequest {
	op := &aws.Operation{
		Name:       opStopTrigger,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopTriggerInput{}
	}

	req := c.newRequest(op, input, &StopTriggerOutput{})

	return StopTriggerRequest{Request: req, Input: input, Copy: c.StopTriggerRequest}
}

// StopTriggerRequest is the request type for the
// StopTrigger API operation.
type StopTriggerRequest struct {
	*aws.Request
	Input *StopTriggerInput
	Copy  func(*StopTriggerInput) StopTriggerRequest
}

// Send marshals and sends the StopTrigger API request.
func (r StopTriggerRequest) Send(ctx context.Context) (*StopTriggerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopTriggerResponse{
		StopTriggerOutput: r.Request.Data.(*StopTriggerOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopTriggerResponse is the response type for the
// StopTrigger API operation.
type StopTriggerResponse struct {
	*StopTriggerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopTrigger request.
func (r *StopTriggerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
