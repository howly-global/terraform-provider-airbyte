// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type Recreation string

const (
	RecreationRecreation Recreation = "recreation"
)

func (e Recreation) ToPointer() *Recreation {
	return &e
}
func (e *Recreation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recreation":
		*e = Recreation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Recreation: %v", v)
	}
}

type SourceRecreation struct {
	// API Key
	Apikey         string     `json:"apikey"`
	QueryCampsites *string    `json:"query_campsites,omitempty"`
	sourceType     Recreation `const:"recreation" json:"sourceType"`
}

func (s SourceRecreation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRecreation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRecreation) GetApikey() string {
	if o == nil {
		return ""
	}
	return o.Apikey
}

func (o *SourceRecreation) GetQueryCampsites() *string {
	if o == nil {
		return nil
	}
	return o.QueryCampsites
}

func (o *SourceRecreation) GetSourceType() Recreation {
	return RecreationRecreation
}
