// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type Zenloop string

const (
	ZenloopZenloop Zenloop = "zenloop"
)

func (e Zenloop) ToPointer() *Zenloop {
	return &e
}
func (e *Zenloop) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zenloop":
		*e = Zenloop(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Zenloop: %v", v)
	}
}

type SourceZenloop struct {
	// Zenloop API Token. You can get the API token in settings page <a href="https://app.zenloop.com/settings/api">here</a>
	APIToken string `json:"api_token"`
	// Zenloop date_from. Format: 2021-10-24T03:30:30Z or 2021-10-24. Leave empty if only data from current data should be synced
	DateFrom   *string `json:"date_from,omitempty"`
	sourceType Zenloop `const:"zenloop" json:"sourceType"`
	// Zenloop Survey Group ID. Can be found by pulling All Survey Groups via SurveyGroups stream. Leave empty to pull answers from all survey groups
	SurveyGroupID *string `json:"survey_group_id,omitempty"`
	// Zenloop Survey ID. Can be found <a href="https://app.zenloop.com/settings/api">here</a>. Leave empty to pull answers from all surveys
	SurveyID *string `json:"survey_id,omitempty"`
}

func (s SourceZenloop) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZenloop) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZenloop) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceZenloop) GetDateFrom() *string {
	if o == nil {
		return nil
	}
	return o.DateFrom
}

func (o *SourceZenloop) GetSourceType() Zenloop {
	return ZenloopZenloop
}

func (o *SourceZenloop) GetSurveyGroupID() *string {
	if o == nil {
		return nil
	}
	return o.SurveyGroupID
}

func (o *SourceZenloop) GetSurveyID() *string {
	if o == nil {
		return nil
	}
	return o.SurveyID
}
