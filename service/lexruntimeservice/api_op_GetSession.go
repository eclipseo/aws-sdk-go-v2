// Code generated by smithy-go-codegen DO NOT EDIT.
package lexruntimeservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns session information for a specified bot, alias, and user ID.
func (c *Client) GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) {
	stack := middleware.NewStack("GetSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSession{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSession{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)
	stack.Finalize.Add(&v4.ComputePayloadSHA256Middleware{}, middleware.Before)
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(options.HTTPSigner), middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "GetSession",
			Err:           err,
		}
	}
	out := result.(*GetSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSessionInput struct {
	// The alias in use for the bot that contains the session data.
	BotAlias *string
	// The name of the bot that contains the session data.
	BotName *string
	// A string used to filter the intents returned in the recentIntentSummaryView
	// structure. When you specify a filter, only intents with their checkpointLabel
	// field set to that string are returned.
	CheckpointLabelFilter *string
	// The ID of the client application user. Amazon Lex uses this to identify a user's
	// conversation with your bot.
	UserId *string
}

type GetSessionOutput struct {
	// Describes the current state of the bot.
	DialogAction *types.DialogAction
	// An array of information about the intents used in the session. The array can
	// contain a maximum of three summaries. If more than three intents are used in the
	// session, the recentIntentSummaryView operation contains information about the
	// last three intents used. If you set the checkpointLabelFilter parameter in the
	// request, the array contains only the intents with the specified label.
	RecentIntentSummaryView []*types.IntentSummary
	// Map of key/value pairs representing the session-specific context information. It
	// contains application information passed between Amazon Lex and a client
	// application.
	SessionAttributes map[string]*string
	// A unique identifier for the session.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}