// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type SourceZoomUpdate struct {
	// The account ID for your Zoom account. You can find this in the Zoom Marketplace under the "Manage" tab for your app.
	AccountID             string  `json:"account_id"`
	AuthorizationEndpoint *string `default:"https://zoom.us/oauth/token" json:"authorization_endpoint"`
	// The client ID for your Zoom app. You can find this in the Zoom Marketplace under the "Manage" tab for your app.
	ClientID string `json:"client_id"`
	// The client secret for your Zoom app. You can find this in the Zoom Marketplace under the "Manage" tab for your app.
	ClientSecret string `json:"client_secret"`
}

func (s SourceZoomUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZoomUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZoomUpdate) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *SourceZoomUpdate) GetAuthorizationEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationEndpoint
}

func (o *SourceZoomUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceZoomUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}