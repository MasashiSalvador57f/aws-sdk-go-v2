// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetRoomInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The room ID.
	//
	// RoomId is a required field
	RoomId *string `location:"uri" locationName:"roomId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRoomInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.RoomId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRoomInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoomId != nil {
		v := *s.RoomId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roomId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetRoomOutput struct {
	_ struct{} `type:"structure"`

	// The room details.
	Room *Room `type:"structure"`
}

// String returns the string representation
func (s GetRoomOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRoomOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Room != nil {
		v := s.Room

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Room", v, metadata)
	}
	return nil
}

const opGetRoom = "GetRoom"

// GetRoomRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves room details, such as the room name.
//
//    // Example sending a request using GetRoomRequest.
//    req := client.GetRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetRoom
func (c *Client) GetRoomRequest(input *GetRoomInput) GetRoomRequest {
	op := &aws.Operation{
		Name:       opGetRoom,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/rooms/{roomId}",
	}

	if input == nil {
		input = &GetRoomInput{}
	}

	req := c.newRequest(op, input, &GetRoomOutput{})
	return GetRoomRequest{Request: req, Input: input, Copy: c.GetRoomRequest}
}

// GetRoomRequest is the request type for the
// GetRoom API operation.
type GetRoomRequest struct {
	*aws.Request
	Input *GetRoomInput
	Copy  func(*GetRoomInput) GetRoomRequest
}

// Send marshals and sends the GetRoom API request.
func (r GetRoomRequest) Send(ctx context.Context) (*GetRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRoomResponse{
		GetRoomOutput: r.Request.Data.(*GetRoomOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRoomResponse is the response type for the
// GetRoom API operation.
type GetRoomResponse struct {
	*GetRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRoom request.
func (r *GetRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}