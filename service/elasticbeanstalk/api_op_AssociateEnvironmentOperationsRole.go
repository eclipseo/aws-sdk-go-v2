// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Request to add or change the operations role used by an environment.
type AssociateEnvironmentOperationsRoleInput struct {
	_ struct{} `type:"structure"`

	// The name of the environment to which to set the operations role.
	//
	// EnvironmentName is a required field
	EnvironmentName *string `min:"4" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the
	// environment's operations role.
	//
	// OperationsRole is a required field
	OperationsRole *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateEnvironmentOperationsRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateEnvironmentOperationsRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateEnvironmentOperationsRoleInput"}

	if s.EnvironmentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentName"))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}

	if s.OperationsRole == nil {
		invalidParams.Add(aws.NewErrParamRequired("OperationsRole"))
	}
	if s.OperationsRole != nil && len(*s.OperationsRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OperationsRole", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateEnvironmentOperationsRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateEnvironmentOperationsRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateEnvironmentOperationsRole = "AssociateEnvironmentOperationsRole"

// AssociateEnvironmentOperationsRoleRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Add or change the operations role used by an environment. After this call
// is made, Elastic Beanstalk uses the associated operations role for permissions
// to downstream services during subsequent calls acting on this environment.
// For more information, see Operations roles (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html)
// in the AWS Elastic Beanstalk Developer Guide.
//
//    // Example sending a request using AssociateEnvironmentOperationsRoleRequest.
//    req := client.AssociateEnvironmentOperationsRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/AssociateEnvironmentOperationsRole
func (c *Client) AssociateEnvironmentOperationsRoleRequest(input *AssociateEnvironmentOperationsRoleInput) AssociateEnvironmentOperationsRoleRequest {
	op := &aws.Operation{
		Name:       opAssociateEnvironmentOperationsRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateEnvironmentOperationsRoleInput{}
	}

	req := c.newRequest(op, input, &AssociateEnvironmentOperationsRoleOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AssociateEnvironmentOperationsRoleRequest{Request: req, Input: input, Copy: c.AssociateEnvironmentOperationsRoleRequest}
}

// AssociateEnvironmentOperationsRoleRequest is the request type for the
// AssociateEnvironmentOperationsRole API operation.
type AssociateEnvironmentOperationsRoleRequest struct {
	*aws.Request
	Input *AssociateEnvironmentOperationsRoleInput
	Copy  func(*AssociateEnvironmentOperationsRoleInput) AssociateEnvironmentOperationsRoleRequest
}

// Send marshals and sends the AssociateEnvironmentOperationsRole API request.
func (r AssociateEnvironmentOperationsRoleRequest) Send(ctx context.Context) (*AssociateEnvironmentOperationsRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateEnvironmentOperationsRoleResponse{
		AssociateEnvironmentOperationsRoleOutput: r.Request.Data.(*AssociateEnvironmentOperationsRoleOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateEnvironmentOperationsRoleResponse is the response type for the
// AssociateEnvironmentOperationsRole API operation.
type AssociateEnvironmentOperationsRoleResponse struct {
	*AssociateEnvironmentOperationsRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateEnvironmentOperationsRole request.
func (r *AssociateEnvironmentOperationsRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
