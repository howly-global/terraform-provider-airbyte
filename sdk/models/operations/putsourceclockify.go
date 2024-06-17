// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutSourceClockifyRequest struct {
	SourceClockifyPutRequest *shared.SourceClockifyPutRequest `request:"mediaType=application/json"`
	SourceID                 string                           `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceClockifyRequest) GetSourceClockifyPutRequest() *shared.SourceClockifyPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceClockifyPutRequest
}

func (o *PutSourceClockifyRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceClockifyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceClockifyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceClockifyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceClockifyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
