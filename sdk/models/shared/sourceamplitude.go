// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

// SourceAmplitudeDataRegion - Amplitude data region server
type SourceAmplitudeDataRegion string

const (
	SourceAmplitudeDataRegionStandardServer    SourceAmplitudeDataRegion = "Standard Server"
	SourceAmplitudeDataRegionEuResidencyServer SourceAmplitudeDataRegion = "EU Residency Server"
)

func (e SourceAmplitudeDataRegion) ToPointer() *SourceAmplitudeDataRegion {
	return &e
}
func (e *SourceAmplitudeDataRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard Server":
		fallthrough
	case "EU Residency Server":
		*e = SourceAmplitudeDataRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmplitudeDataRegion: %v", v)
	}
}

type Amplitude string

const (
	AmplitudeAmplitude Amplitude = "amplitude"
)

func (e Amplitude) ToPointer() *Amplitude {
	return &e
}
func (e *Amplitude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "amplitude":
		*e = Amplitude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Amplitude: %v", v)
	}
}

type SourceAmplitude struct {
	// Amplitude API Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
	APIKey string `json:"api_key"`
	// Amplitude data region server
	DataRegion *SourceAmplitudeDataRegion `default:"Standard Server" json:"data_region"`
	// According to <a href="https://www.docs.developers.amplitude.com/analytics/apis/export-api/#considerations">Considerations</a> too big time range in request can cause a timeout error. In this case, set shorter time interval in hours.
	RequestTimeRange *int64 `default:"24" json:"request_time_range"`
	// Amplitude Secret Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
	SecretKey  string    `json:"secret_key"`
	sourceType Amplitude `const:"amplitude" json:"sourceType"`
	// UTC date and time in the format 2021-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceAmplitude) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAmplitude) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAmplitude) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceAmplitude) GetDataRegion() *SourceAmplitudeDataRegion {
	if o == nil {
		return nil
	}
	return o.DataRegion
}

func (o *SourceAmplitude) GetRequestTimeRange() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeRange
}

func (o *SourceAmplitude) GetSecretKey() string {
	if o == nil {
		return ""
	}
	return o.SecretKey
}

func (o *SourceAmplitude) GetSourceType() Amplitude {
	return AmplitudeAmplitude
}

func (o *SourceAmplitude) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
