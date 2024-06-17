// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutSourceGoogleAnalyticsV4ServiceAccountOnlyRequest struct {
	SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest `request:"mediaType=application/json"`
	SourceID                                            string                                                      `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceGoogleAnalyticsV4ServiceAccountOnlyRequest) GetSourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest() *shared.SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest
}

func (o *PutSourceGoogleAnalyticsV4ServiceAccountOnlyRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceGoogleAnalyticsV4ServiceAccountOnlyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceGoogleAnalyticsV4ServiceAccountOnlyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceGoogleAnalyticsV4ServiceAccountOnlyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceGoogleAnalyticsV4ServiceAccountOnlyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
