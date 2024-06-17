// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutDestinationPineconeRequest struct {
	DestinationPineconePutRequest *shared.DestinationPineconePutRequest `request:"mediaType=application/json"`
	DestinationID                 string                                `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationPineconeRequest) GetDestinationPineconePutRequest() *shared.DestinationPineconePutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationPineconePutRequest
}

func (o *PutDestinationPineconeRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationPineconeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationPineconeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationPineconeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationPineconeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
