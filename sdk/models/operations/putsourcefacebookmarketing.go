// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutSourceFacebookMarketingRequest struct {
	SourceFacebookMarketingPutRequest *shared.SourceFacebookMarketingPutRequest `request:"mediaType=application/json"`
	SourceID                          string                                    `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceFacebookMarketingRequest) GetSourceFacebookMarketingPutRequest() *shared.SourceFacebookMarketingPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceFacebookMarketingPutRequest
}

func (o *PutSourceFacebookMarketingRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceFacebookMarketingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceFacebookMarketingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceFacebookMarketingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceFacebookMarketingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
