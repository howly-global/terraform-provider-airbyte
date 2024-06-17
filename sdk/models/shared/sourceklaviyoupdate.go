// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type SourceKlaviyoUpdate struct {
	// Klaviyo API Key. See our <a href="https://docs.airbyte.com/integrations/sources/klaviyo">docs</a> if you need help finding this key.
	APIKey string `json:"api_key"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. This field is optional - if not provided, all data will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceKlaviyoUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceKlaviyoUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceKlaviyoUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceKlaviyoUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}