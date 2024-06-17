// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type SourceLinnworksUpdate struct {
	// Linnworks Application ID
	ApplicationID string `json:"application_id"`
	// Linnworks Application Secret
	ApplicationSecret string `json:"application_secret"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
	Token     string    `json:"token"`
}

func (s SourceLinnworksUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinnworksUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinnworksUpdate) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

func (o *SourceLinnworksUpdate) GetApplicationSecret() string {
	if o == nil {
		return ""
	}
	return o.ApplicationSecret
}

func (o *SourceLinnworksUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceLinnworksUpdate) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}