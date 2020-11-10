// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the security configurations for the resource policies set on
// individual resources, and also the account-level policy. This operation also
// returns the Data Catalog resource policy. However, if you enabled metadata
// encryption in Data Catalog settings, and you do not have permission on the AWS
// KMS key, the operation can't return the Data Catalog resource policy.
func (c *Client) GetResourcePolicies(ctx context.Context, params *GetResourcePoliciesInput, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) {
	if params == nil {
		params = &GetResourcePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourcePolicies", params, optFns, addOperationGetResourcePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourcePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourcePoliciesInput struct {

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string
}

type GetResourcePoliciesOutput struct {

	// A list of the individual resource policies and the account-level resource
	// policy.
	GetResourcePoliciesResponseList []*types.GluePolicy

	// A continuation token, if the returned list does not contain the last resource
	// policy available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResourcePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResourcePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResourcePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourcePolicies(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetResourcePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetResourcePolicies",
	}
}