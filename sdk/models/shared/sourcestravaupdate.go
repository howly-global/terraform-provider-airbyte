// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type SourceStravaUpdateAuthType string

const (
	SourceStravaUpdateAuthTypeClient SourceStravaUpdateAuthType = "Client"
)

func (e SourceStravaUpdateAuthType) ToPointer() *SourceStravaUpdateAuthType {
	return &e
}
func (e *SourceStravaUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceStravaUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceStravaUpdateAuthType: %v", v)
	}
}

type SourceStravaUpdate struct {
	// The Athlete ID of your Strava developer application.
	AthleteID int64                       `json:"athlete_id"`
	authType  *SourceStravaUpdateAuthType `const:"Client" json:"auth_type"`
	// The Client ID of your Strava developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Strava developer application.
	ClientSecret string `json:"client_secret"`
	// The Refresh Token with the activity: read_all permissions.
	RefreshToken string `json:"refresh_token"`
	// UTC date and time. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceStravaUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceStravaUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceStravaUpdate) GetAthleteID() int64 {
	if o == nil {
		return 0
	}
	return o.AthleteID
}

func (o *SourceStravaUpdate) GetAuthType() *SourceStravaUpdateAuthType {
	return SourceStravaUpdateAuthTypeClient.ToPointer()
}

func (o *SourceStravaUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceStravaUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceStravaUpdate) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceStravaUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
