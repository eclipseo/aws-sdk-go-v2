// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to add a new Model to an existing RestApi resource.
type CreateModelInput struct {
	_ struct{} `type:"structure"`

	// [Required] The content-type for the model.
	//
	// ContentType is a required field
	ContentType *string `locationName:"contentType" type:"string" required:"true"`

	// The description of the model.
	Description *string `locationName:"description" type:"string"`

	// [Required] The name of the model. Must be alphanumeric.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// [Required] The RestApi identifier under which the Model will be created.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model.
	Schema *string `locationName:"schema" type:"string"`
}

// String returns the string representation
func (s CreateModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateModelInput"}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateModelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents the data structure of a method's request or response payload.
//
// A request model defines the data structure of the client-supplied request
// payload. A response model defines the data structure of the response payload
// returned by the back end. Although not required, models are useful for mapping
// payloads between the front end and back end.
//
// A model is used for generating an API's SDK, validating the input request
// body, and creating a skeletal mapping template.
//
// Method, MethodResponse, Models and Mappings (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html)
type CreateModelOutput struct {
	_ struct{} `type:"structure"`

	// The content-type for the model.
	ContentType *string `locationName:"contentType" type:"string"`

	// The description of the model.
	Description *string `locationName:"description" type:"string"`

	// The identifier for the model resource.
	Id *string `locationName:"id" type:"string"`

	// The name of the model. Must be an alphanumeric string.
	Name *string `locationName:"name" type:"string"`

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model.
	// Do not include "\*/" characters in the description of any properties because
	// such "\*/" characters may be interpreted as the closing marker for comments
	// in some languages, such as Java or JavaScript, causing the installation of
	// your API's SDK generated by API Gateway to fail.
	Schema *string `locationName:"schema" type:"string"`
}

// String returns the string representation
func (s CreateModelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateModelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateModel = "CreateModel"

// CreateModelRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Adds a new Model resource to an existing RestApi resource.
//
//    // Example sending a request using CreateModelRequest.
//    req := client.CreateModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateModelRequest(input *CreateModelInput) CreateModelRequest {
	op := &aws.Operation{
		Name:       opCreateModel,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/models",
	}

	if input == nil {
		input = &CreateModelInput{}
	}

	req := c.newRequest(op, input, &CreateModelOutput{})

	return CreateModelRequest{Request: req, Input: input, Copy: c.CreateModelRequest}
}

// CreateModelRequest is the request type for the
// CreateModel API operation.
type CreateModelRequest struct {
	*aws.Request
	Input *CreateModelInput
	Copy  func(*CreateModelInput) CreateModelRequest
}

// Send marshals and sends the CreateModel API request.
func (r CreateModelRequest) Send(ctx context.Context) (*CreateModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateModelResponse{
		CreateModelOutput: r.Request.Data.(*CreateModelOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateModelResponse is the response type for the
// CreateModel API operation.
type CreateModelResponse struct {
	*CreateModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateModel request.
func (r *CreateModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
