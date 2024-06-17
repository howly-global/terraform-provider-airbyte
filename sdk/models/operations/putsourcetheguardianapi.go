// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutSourceTheGuardianAPIRequest struct {
	SourceTheGuardianAPIPutRequest *shared.SourceTheGuardianAPIPutRequest `request:"mediaType=application/json"`
	SourceID                       string                                 `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceTheGuardianAPIRequest) GetSourceTheGuardianAPIPutRequest() *shared.SourceTheGuardianAPIPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceTheGuardianAPIPutRequest
}

func (o *PutSourceTheGuardianAPIRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceTheGuardianAPIResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceTheGuardianAPIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceTheGuardianAPIResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceTheGuardianAPIResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
