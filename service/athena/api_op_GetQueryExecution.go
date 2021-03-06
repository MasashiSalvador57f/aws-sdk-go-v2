// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetQueryExecutionInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query execution.
	//
	// QueryExecutionId is a required field
	QueryExecutionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetQueryExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueryExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueryExecutionInput"}

	if s.QueryExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryExecutionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetQueryExecutionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the query execution.
	QueryExecution *QueryExecution `type:"structure"`
}

// String returns the string representation
func (s GetQueryExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetQueryExecution = "GetQueryExecution"

// GetQueryExecutionRequest returns a request value for making API operation for
// Amazon Athena.
//
// Returns information about a single execution of a query if you have access
// to the workgroup in which the query ran. Each time a query executes, information
// about the query execution is saved with a unique ID.
//
//    // Example sending a request using GetQueryExecutionRequest.
//    req := client.GetQueryExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/GetQueryExecution
func (c *Client) GetQueryExecutionRequest(input *GetQueryExecutionInput) GetQueryExecutionRequest {
	op := &aws.Operation{
		Name:       opGetQueryExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQueryExecutionInput{}
	}

	req := c.newRequest(op, input, &GetQueryExecutionOutput{})

	return GetQueryExecutionRequest{Request: req, Input: input, Copy: c.GetQueryExecutionRequest}
}

// GetQueryExecutionRequest is the request type for the
// GetQueryExecution API operation.
type GetQueryExecutionRequest struct {
	*aws.Request
	Input *GetQueryExecutionInput
	Copy  func(*GetQueryExecutionInput) GetQueryExecutionRequest
}

// Send marshals and sends the GetQueryExecution API request.
func (r GetQueryExecutionRequest) Send(ctx context.Context) (*GetQueryExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetQueryExecutionResponse{
		GetQueryExecutionOutput: r.Request.Data.(*GetQueryExecutionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetQueryExecutionResponse is the response type for the
// GetQueryExecution API operation.
type GetQueryExecutionResponse struct {
	*GetQueryExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetQueryExecution request.
func (r *GetQueryExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
