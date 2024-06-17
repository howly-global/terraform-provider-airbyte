// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type SourceOnesignalApplications struct {
	AppAPIKey string  `json:"app_api_key"`
	AppID     string  `json:"app_id"`
	AppName   *string `json:"app_name,omitempty"`
}

func (o *SourceOnesignalApplications) GetAppAPIKey() string {
	if o == nil {
		return ""
	}
	return o.AppAPIKey
}

func (o *SourceOnesignalApplications) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *SourceOnesignalApplications) GetAppName() *string {
	if o == nil {
		return nil
	}
	return o.AppName
}

type Onesignal string

const (
	OnesignalOnesignal Onesignal = "onesignal"
)

func (e Onesignal) ToPointer() *Onesignal {
	return &e
}
func (e *Onesignal) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "onesignal":
		*e = Onesignal(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Onesignal: %v", v)
	}
}

type SourceOnesignal struct {
	// Applications keys, see the <a href="https://documentation.onesignal.com/docs/accounts-and-keys">docs</a> for more information on how to obtain this data
	Applications []SourceOnesignalApplications `json:"applications"`
	// Comma-separated list of names and the value (sum/count) for the returned outcome data. See the <a href="https://documentation.onesignal.com/reference/view-outcomes">docs</a> for more details
	OutcomeNames string    `json:"outcome_names"`
	sourceType   Onesignal `const:"onesignal" json:"sourceType"`
	// The date from which you'd like to replicate data for OneSignal API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// OneSignal User Auth Key, see the <a href="https://documentation.onesignal.com/docs/accounts-and-keys#user-auth-key">docs</a> for more information on how to obtain this key.
	UserAuthKey string `json:"user_auth_key"`
}

func (s SourceOnesignal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOnesignal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOnesignal) GetApplications() []SourceOnesignalApplications {
	if o == nil {
		return []SourceOnesignalApplications{}
	}
	return o.Applications
}

func (o *SourceOnesignal) GetOutcomeNames() string {
	if o == nil {
		return ""
	}
	return o.OutcomeNames
}

func (o *SourceOnesignal) GetSourceType() Onesignal {
	return OnesignalOnesignal
}

func (o *SourceOnesignal) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceOnesignal) GetUserAuthKey() string {
	if o == nil {
		return ""
	}
	return o.UserAuthKey
}
