// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type GetDestinationS3Request struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *GetDestinationS3Request) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type GetDestinationS3Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Get a Destination by the id in the path.
	DestinationResponse *shared.DestinationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDestinationS3Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDestinationS3Response) GetDestinationResponse() *shared.DestinationResponse {
	if o == nil {
		return nil
	}
	return o.DestinationResponse
}

func (o *GetDestinationS3Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDestinationS3Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
