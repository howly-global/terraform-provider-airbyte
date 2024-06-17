// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type Freshdesk string

const (
	FreshdeskFreshdesk Freshdesk = "freshdesk"
)

func (e Freshdesk) ToPointer() *Freshdesk {
	return &e
}
func (e *Freshdesk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "freshdesk":
		*e = Freshdesk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Freshdesk: %v", v)
	}
}

type SourceFreshdesk struct {
	// Freshdesk API Key. See the <a href="https://docs.airbyte.com/integrations/sources/freshdesk">docs</a> for more information on how to obtain this key.
	APIKey string `json:"api_key"`
	// Freshdesk domain
	Domain string `json:"domain"`
	// Number of days for lookback window for the stream Satisfaction Ratings
	LookbackWindowInDays *int64 `default:"14" json:"lookback_window_in_days"`
	// The number of requests per minute that this source allowed to use. There is a rate limit of 50 requests per minute per app per account.
	RequestsPerMinute *int64    `json:"requests_per_minute,omitempty"`
	sourceType        Freshdesk `const:"freshdesk" json:"sourceType"`
	// UTC date and time. Any data created after this date will be replicated. If this parameter is not set, all data will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceFreshdesk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFreshdesk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFreshdesk) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceFreshdesk) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *SourceFreshdesk) GetLookbackWindowInDays() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindowInDays
}

func (o *SourceFreshdesk) GetRequestsPerMinute() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestsPerMinute
}

func (o *SourceFreshdesk) GetSourceType() Freshdesk {
	return FreshdeskFreshdesk
}

func (o *SourceFreshdesk) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
