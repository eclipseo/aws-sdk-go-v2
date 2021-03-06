// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListOrganizationAdminAccountsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListOrganizationAdminAccountsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOrganizationAdminAccountsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOrganizationAdminAccountsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListOrganizationAdminAccountsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the accounts that are designated as delegated
// administrators of Amazon Macie for an AWS organization.
type ListOrganizationAdminAccountsOutput struct {
	_ struct{} `type:"structure"`

	AdminAccounts []AdminAccount `locationName:"adminAccounts" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListOrganizationAdminAccountsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListOrganizationAdminAccountsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdminAccounts != nil {
		v := s.AdminAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "adminAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListOrganizationAdminAccounts = "ListOrganizationAdminAccounts"

// ListOrganizationAdminAccountsRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves information about the account that's designated as the delegated
// administrator of Amazon Macie for an AWS organization.
//
//    // Example sending a request using ListOrganizationAdminAccountsRequest.
//    req := client.ListOrganizationAdminAccountsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/ListOrganizationAdminAccounts
func (c *Client) ListOrganizationAdminAccountsRequest(input *ListOrganizationAdminAccountsInput) ListOrganizationAdminAccountsRequest {
	op := &aws.Operation{
		Name:       opListOrganizationAdminAccounts,
		HTTPMethod: "GET",
		HTTPPath:   "/admin",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListOrganizationAdminAccountsInput{}
	}

	req := c.newRequest(op, input, &ListOrganizationAdminAccountsOutput{})

	return ListOrganizationAdminAccountsRequest{Request: req, Input: input, Copy: c.ListOrganizationAdminAccountsRequest}
}

// ListOrganizationAdminAccountsRequest is the request type for the
// ListOrganizationAdminAccounts API operation.
type ListOrganizationAdminAccountsRequest struct {
	*aws.Request
	Input *ListOrganizationAdminAccountsInput
	Copy  func(*ListOrganizationAdminAccountsInput) ListOrganizationAdminAccountsRequest
}

// Send marshals and sends the ListOrganizationAdminAccounts API request.
func (r ListOrganizationAdminAccountsRequest) Send(ctx context.Context) (*ListOrganizationAdminAccountsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOrganizationAdminAccountsResponse{
		ListOrganizationAdminAccountsOutput: r.Request.Data.(*ListOrganizationAdminAccountsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListOrganizationAdminAccountsRequestPaginator returns a paginator for ListOrganizationAdminAccounts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListOrganizationAdminAccountsRequest(input)
//   p := macie2.NewListOrganizationAdminAccountsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListOrganizationAdminAccountsPaginator(req ListOrganizationAdminAccountsRequest) ListOrganizationAdminAccountsPaginator {
	return ListOrganizationAdminAccountsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListOrganizationAdminAccountsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListOrganizationAdminAccountsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListOrganizationAdminAccountsPaginator struct {
	aws.Pager
}

func (p *ListOrganizationAdminAccountsPaginator) CurrentPage() *ListOrganizationAdminAccountsOutput {
	return p.Pager.CurrentPage().(*ListOrganizationAdminAccountsOutput)
}

// ListOrganizationAdminAccountsResponse is the response type for the
// ListOrganizationAdminAccounts API operation.
type ListOrganizationAdminAccountsResponse struct {
	*ListOrganizationAdminAccountsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOrganizationAdminAccounts request.
func (r *ListOrganizationAdminAccountsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
