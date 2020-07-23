// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used for requests to update project configuration.
type UpdateProjectInput struct {
	_ struct{} `type:"structure" payload:"Contents"`

	// ZIP or YAML file which contains project configuration to be updated. This
	// should be the contents of the file downloaded from the URL provided in an
	// export project operation.
	Contents []byte `locationName:"contents" type:"blob"`

	// Unique project identifier.
	//
	// ProjectId is a required field
	ProjectId *string `location:"querystring" locationName:"projectId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProjectInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Contents != nil {
		v := s.Contents

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "contents", protocol.BytesStream(v), metadata)
	}
	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure used for requests to updated project configuration.
type UpdateProjectOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the updated AWS Mobile Hub project.
	Details *ProjectDetails `locationName:"details" type:"structure"`
}

// String returns the string representation
func (s UpdateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Details != nil {
		v := s.Details

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "details", v, metadata)
	}
	return nil
}

const opUpdateProject = "UpdateProject"

// UpdateProjectRequest returns a request value for making API operation for
// AWS Mobile.
//
// Update an existing project.
//
//    // Example sending a request using UpdateProjectRequest.
//    req := client.UpdateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/UpdateProject
func (c *Client) UpdateProjectRequest(input *UpdateProjectInput) UpdateProjectRequest {
	op := &aws.Operation{
		Name:       opUpdateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/update",
	}

	if input == nil {
		input = &UpdateProjectInput{}
	}

	req := c.newRequest(op, input, &UpdateProjectOutput{})

	return UpdateProjectRequest{Request: req, Input: input, Copy: c.UpdateProjectRequest}
}

// UpdateProjectRequest is the request type for the
// UpdateProject API operation.
type UpdateProjectRequest struct {
	*aws.Request
	Input *UpdateProjectInput
	Copy  func(*UpdateProjectInput) UpdateProjectRequest
}

// Send marshals and sends the UpdateProject API request.
func (r UpdateProjectRequest) Send(ctx context.Context) (*UpdateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProjectResponse{
		UpdateProjectOutput: r.Request.Data.(*UpdateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProjectResponse is the response type for the
// UpdateProject API operation.
type UpdateProjectResponse struct {
	*UpdateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProject request.
func (r *UpdateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}