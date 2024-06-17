// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutSourceOrbitRequest struct {
	SourceOrbitPutRequest *shared.SourceOrbitPutRequest `request:"mediaType=application/json"`
	SourceID              string                        `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceOrbitRequest) GetSourceOrbitPutRequest() *shared.SourceOrbitPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceOrbitPutRequest
}

func (o *PutSourceOrbitRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceOrbitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceOrbitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceOrbitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceOrbitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
