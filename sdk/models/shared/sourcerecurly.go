// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type Recurly string

const (
	RecurlyRecurly Recurly = "recurly"
)

func (e Recurly) ToPointer() *Recurly {
	return &e
}
func (e *Recurly) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recurly":
		*e = Recurly(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Recurly: %v", v)
	}
}

type SourceRecurly struct {
	// Recurly API Key. See the  <a href="https://docs.airbyte.com/integrations/sources/recurly">docs</a> for more information on how to generate this key.
	APIKey string `json:"api_key"`
	// ISO8601 timestamp from which the replication from Recurly API will start from.
	BeginTime *string `json:"begin_time,omitempty"`
	// ISO8601 timestamp to which the replication from Recurly API will stop. Records after that date won't be imported.
	EndTime    *string `json:"end_time,omitempty"`
	sourceType Recurly `const:"recurly" json:"sourceType"`
}

func (s SourceRecurly) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRecurly) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRecurly) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceRecurly) GetBeginTime() *string {
	if o == nil {
		return nil
	}
	return o.BeginTime
}

func (o *SourceRecurly) GetEndTime() *string {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *SourceRecurly) GetSourceType() Recurly {
	return RecurlyRecurly
}
