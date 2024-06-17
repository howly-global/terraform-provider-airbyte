// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type SourceMailgunUpdate struct {
	// Domain region code. 'EU' or 'US' are possible values. The default is 'US'.
	DomainRegion *string `default:"US" json:"domain_region"`
	// Primary account API key to access your Mailgun data.
	PrivateKey string `json:"private_key"`
	// UTC date and time in the format 2020-10-01 00:00:00. Any data before this date will not be replicated. If omitted, defaults to 3 days ago.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceMailgunUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailgunUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailgunUpdate) GetDomainRegion() *string {
	if o == nil {
		return nil
	}
	return o.DomainRegion
}

func (o *SourceMailgunUpdate) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *SourceMailgunUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
