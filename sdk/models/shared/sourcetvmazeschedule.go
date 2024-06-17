// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type TvmazeSchedule string

const (
	TvmazeScheduleTvmazeSchedule TvmazeSchedule = "tvmaze-schedule"
)

func (e TvmazeSchedule) ToPointer() *TvmazeSchedule {
	return &e
}
func (e *TvmazeSchedule) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tvmaze-schedule":
		*e = TvmazeSchedule(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TvmazeSchedule: %v", v)
	}
}

type SourceTvmazeSchedule struct {
	// Country code for domestic TV schedule retrieval.
	DomesticScheduleCountryCode string `json:"domestic_schedule_country_code"`
	// End date for TV schedule retrieval. May be in the future. Optional.
	//
	EndDate    *string        `json:"end_date,omitempty"`
	sourceType TvmazeSchedule `const:"tvmaze-schedule" json:"sourceType"`
	// Start date for TV schedule retrieval. May be in the future.
	StartDate string `json:"start_date"`
	// ISO 3166-1 country code for web TV schedule retrieval. Leave blank for
	// all countries plus global web channels (e.g. Netflix). Alternatively,
	// set to 'global' for just global web channels.
	//
	WebScheduleCountryCode *string `json:"web_schedule_country_code,omitempty"`
}

func (s SourceTvmazeSchedule) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceTvmazeSchedule) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceTvmazeSchedule) GetDomesticScheduleCountryCode() string {
	if o == nil {
		return ""
	}
	return o.DomesticScheduleCountryCode
}

func (o *SourceTvmazeSchedule) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceTvmazeSchedule) GetSourceType() TvmazeSchedule {
	return TvmazeScheduleTvmazeSchedule
}

func (o *SourceTvmazeSchedule) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

func (o *SourceTvmazeSchedule) GetWebScheduleCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.WebScheduleCountryCode
}
