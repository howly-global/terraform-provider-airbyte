// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"net/http"
)

type PutDestinationVectaraRequest struct {
	DestinationVectaraPutRequest *shared.DestinationVectaraPutRequest `request:"mediaType=application/json"`
	DestinationID                string                               `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationVectaraRequest) GetDestinationVectaraPutRequest() *shared.DestinationVectaraPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationVectaraPutRequest
}

func (o *PutDestinationVectaraRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationVectaraResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationVectaraResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationVectaraResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationVectaraResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
