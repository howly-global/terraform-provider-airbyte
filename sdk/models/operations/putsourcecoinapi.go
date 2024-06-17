// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutSourceCoinAPIRequest struct {
	SourceCoinAPIPutRequest *shared.SourceCoinAPIPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceCoinAPIRequest) GetSourceCoinAPIPutRequest() *shared.SourceCoinAPIPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceCoinAPIPutRequest
}

func (o *PutSourceCoinAPIRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceCoinAPIResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceCoinAPIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceCoinAPIResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceCoinAPIResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
