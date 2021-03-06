// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchCreateVariableInput struct {
	_ struct{} `type:"structure"`

	// A collection of key and value pairs.
	Tags []Tag `locationName:"tags" type:"list"`

	// The list of variables for the batch create variable request.
	//
	// VariableEntries is a required field
	VariableEntries []VariableEntry `locationName:"variableEntries" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchCreateVariableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchCreateVariableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchCreateVariableInput"}

	if s.VariableEntries == nil {
		invalidParams.Add(aws.NewErrParamRequired("VariableEntries"))
	}
	if s.VariableEntries != nil && len(s.VariableEntries) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VariableEntries", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchCreateVariableOutput struct {
	_ struct{} `type:"structure"`

	// Provides the errors for the BatchCreateVariable request.
	Errors []BatchCreateVariableError `locationName:"errors" type:"list"`
}

// String returns the string representation
func (s BatchCreateVariableOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchCreateVariable = "BatchCreateVariable"

// BatchCreateVariableRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Creates a batch of variables.
//
//    // Example sending a request using BatchCreateVariableRequest.
//    req := client.BatchCreateVariableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/BatchCreateVariable
func (c *Client) BatchCreateVariableRequest(input *BatchCreateVariableInput) BatchCreateVariableRequest {
	op := &aws.Operation{
		Name:       opBatchCreateVariable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchCreateVariableInput{}
	}

	req := c.newRequest(op, input, &BatchCreateVariableOutput{})

	return BatchCreateVariableRequest{Request: req, Input: input, Copy: c.BatchCreateVariableRequest}
}

// BatchCreateVariableRequest is the request type for the
// BatchCreateVariable API operation.
type BatchCreateVariableRequest struct {
	*aws.Request
	Input *BatchCreateVariableInput
	Copy  func(*BatchCreateVariableInput) BatchCreateVariableRequest
}

// Send marshals and sends the BatchCreateVariable API request.
func (r BatchCreateVariableRequest) Send(ctx context.Context) (*BatchCreateVariableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchCreateVariableResponse{
		BatchCreateVariableOutput: r.Request.Data.(*BatchCreateVariableOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchCreateVariableResponse is the response type for the
// BatchCreateVariable API operation.
type BatchCreateVariableResponse struct {
	*BatchCreateVariableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchCreateVariable request.
func (r *BatchCreateVariableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
