// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutDestinationMongodbRequest struct {
	DestinationMongodbPutRequest *shared.DestinationMongodbPutRequest `request:"mediaType=application/json"`
	DestinationID                string                               `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationMongodbRequest) GetDestinationMongodbPutRequest() *shared.DestinationMongodbPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationMongodbPutRequest
}

func (o *PutDestinationMongodbRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationMongodbResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationMongodbResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationMongodbResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationMongodbResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
