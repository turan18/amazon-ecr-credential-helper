// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets detailed information for an image. Images are specified with either an
// imageTag or imageDigest . When an image is pulled, the BatchGetImage API is
// called once to retrieve the image manifest.
func (c *Client) BatchGetImage(ctx context.Context, params *BatchGetImageInput, optFns ...func(*Options)) (*BatchGetImageOutput, error) {
	if params == nil {
		params = &BatchGetImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetImage", params, optFns, c.addOperationBatchGetImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetImageInput struct {

	// A list of image ID references that correspond to images to describe. The format
	// of the imageIds reference is imageTag=tag or imageDigest=digest .
	//
	// This member is required.
	ImageIds []types.ImageIdentifier

	// The repository that contains the images to describe.
	//
	// This member is required.
	RepositoryName *string

	// The accepted media types for the request. Valid values:
	// application/vnd.docker.distribution.manifest.v1+json |
	// application/vnd.docker.distribution.manifest.v2+json |
	// application/vnd.oci.image.manifest.v1+json
	AcceptedMediaTypes []string

	// The Amazon Web Services account ID associated with the registry that contains
	// the images to describe. If you do not specify a registry, the default registry
	// is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type BatchGetImageOutput struct {

	// Any failures associated with the call.
	Failures []types.ImageFailure

	// A list of image objects corresponding to the image references in the request.
	Images []types.Image

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetImage{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchGetImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetImage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "BatchGetImage",
	}
}
