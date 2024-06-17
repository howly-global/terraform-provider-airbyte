// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type CreateDestinationAstraResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	DestinationResponse *shared.DestinationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateDestinationAstraResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDestinationAstraResponse) GetDestinationResponse() *shared.DestinationResponse {
	if o == nil {
		return nil
	}
	return o.DestinationResponse
}

func (o *CreateDestinationAstraResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDestinationAstraResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
