// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"github.com/howly-global/terraform-provider-airbyte/sdk/types"
)

// ActionReportTime - Specifies the principle for conversion reporting.
type ActionReportTime string

const (
	ActionReportTimeConversion ActionReportTime = "conversion"
	ActionReportTimeImpression ActionReportTime = "impression"
)

func (e ActionReportTime) ToPointer() *ActionReportTime {
	return &e
}
func (e *ActionReportTime) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "conversion":
		fallthrough
	case "impression":
		*e = ActionReportTime(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActionReportTime: %v", v)
	}
}

// SwipeUpAttributionWindow - Attribution window for swipe ups.
type SwipeUpAttributionWindow string

const (
	SwipeUpAttributionWindowOneDay         SwipeUpAttributionWindow = "1_DAY"
	SwipeUpAttributionWindowSevenDay       SwipeUpAttributionWindow = "7_DAY"
	SwipeUpAttributionWindowTwentyEightDay SwipeUpAttributionWindow = "28_DAY"
)

func (e SwipeUpAttributionWindow) ToPointer() *SwipeUpAttributionWindow {
	return &e
}
func (e *SwipeUpAttributionWindow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1_DAY":
		fallthrough
	case "7_DAY":
		fallthrough
	case "28_DAY":
		*e = SwipeUpAttributionWindow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SwipeUpAttributionWindow: %v", v)
	}
}

// ViewAttributionWindow - Attribution window for views.
type ViewAttributionWindow string

const (
	ViewAttributionWindowOneHour   ViewAttributionWindow = "1_HOUR"
	ViewAttributionWindowThreeHour ViewAttributionWindow = "3_HOUR"
	ViewAttributionWindowSixHour   ViewAttributionWindow = "6_HOUR"
	ViewAttributionWindowOneDay    ViewAttributionWindow = "1_DAY"
	ViewAttributionWindowSevenDay  ViewAttributionWindow = "7_DAY"
)

func (e ViewAttributionWindow) ToPointer() *ViewAttributionWindow {
	return &e
}
func (e *ViewAttributionWindow) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1_HOUR":
		fallthrough
	case "3_HOUR":
		fallthrough
	case "6_HOUR":
		fallthrough
	case "1_DAY":
		fallthrough
	case "7_DAY":
		*e = ViewAttributionWindow(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ViewAttributionWindow: %v", v)
	}
}

type SourceSnapchatMarketingUpdate struct {
	// Specifies the principle for conversion reporting.
	ActionReportTime *ActionReportTime `default:"conversion" json:"action_report_time"`
	// The Client ID of your Snapchat developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Snapchat developer application.
	ClientSecret string `json:"client_secret"`
	// Date in the format 2017-01-25. Any data after this date will not be replicated.
	EndDate *types.Date `json:"end_date,omitempty"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// Date in the format 2022-01-01. Any data before this date will not be replicated.
	StartDate *types.Date `default:"2022-01-01" json:"start_date"`
	// Attribution window for swipe ups.
	SwipeUpAttributionWindow *SwipeUpAttributionWindow `default:"28_DAY" json:"swipe_up_attribution_window"`
	// Attribution window for views.
	ViewAttributionWindow *ViewAttributionWindow `default:"1_DAY" json:"view_attribution_window"`
}

func (s SourceSnapchatMarketingUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSnapchatMarketingUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSnapchatMarketingUpdate) GetActionReportTime() *ActionReportTime {
	if o == nil {
		return nil
	}
	return o.ActionReportTime
}

func (o *SourceSnapchatMarketingUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSnapchatMarketingUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSnapchatMarketingUpdate) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceSnapchatMarketingUpdate) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceSnapchatMarketingUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceSnapchatMarketingUpdate) GetSwipeUpAttributionWindow() *SwipeUpAttributionWindow {
	if o == nil {
		return nil
	}
	return o.SwipeUpAttributionWindow
}

func (o *SourceSnapchatMarketingUpdate) GetViewAttributionWindow() *ViewAttributionWindow {
	if o == nil {
		return nil
	}
	return o.ViewAttributionWindow
}
