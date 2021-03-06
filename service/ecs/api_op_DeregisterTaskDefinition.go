// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterTaskDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition to deregister. You must specify a revision.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterTaskDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterTaskDefinitionInput"}

	if s.TaskDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterTaskDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the deregistered task.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`
}

// String returns the string representation
func (s DeregisterTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterTaskDefinition = "DeregisterTaskDefinition"

// DeregisterTaskDefinitionRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Deregisters the specified task definition by family and revision. Upon deregistration,
// the task definition is marked as INACTIVE. Existing tasks and services that
// reference an INACTIVE task definition continue to run without disruption.
// Existing services that reference an INACTIVE task definition can still scale
// up or down by modifying the service's desired count.
//
// You cannot use an INACTIVE task definition to run new tasks or create new
// services, and you cannot update an existing service to reference an INACTIVE
// task definition. However, there may be up to a 10-minute window following
// deregistration where these restrictions have not yet taken effect.
//
// At this time, INACTIVE task definitions remain discoverable in your account
// indefinitely. However, this behavior is subject to change in the future,
// so you should not rely on INACTIVE task definitions persisting beyond the
// lifecycle of any associated tasks and services.
//
//    // Example sending a request using DeregisterTaskDefinitionRequest.
//    req := client.DeregisterTaskDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeregisterTaskDefinition
func (c *Client) DeregisterTaskDefinitionRequest(input *DeregisterTaskDefinitionInput) DeregisterTaskDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeregisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterTaskDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeregisterTaskDefinitionOutput{})

	return DeregisterTaskDefinitionRequest{Request: req, Input: input, Copy: c.DeregisterTaskDefinitionRequest}
}

// DeregisterTaskDefinitionRequest is the request type for the
// DeregisterTaskDefinition API operation.
type DeregisterTaskDefinitionRequest struct {
	*aws.Request
	Input *DeregisterTaskDefinitionInput
	Copy  func(*DeregisterTaskDefinitionInput) DeregisterTaskDefinitionRequest
}

// Send marshals and sends the DeregisterTaskDefinition API request.
func (r DeregisterTaskDefinitionRequest) Send(ctx context.Context) (*DeregisterTaskDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTaskDefinitionResponse{
		DeregisterTaskDefinitionOutput: r.Request.Data.(*DeregisterTaskDefinitionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTaskDefinitionResponse is the response type for the
// DeregisterTaskDefinition API operation.
type DeregisterTaskDefinitionResponse struct {
	*DeregisterTaskDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTaskDefinition request.
func (r *DeregisterTaskDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
