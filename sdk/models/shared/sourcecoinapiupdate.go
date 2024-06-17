// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

// Environment - The environment to use. Either sandbox or production.
type Environment string

const (
	EnvironmentSandbox    Environment = "sandbox"
	EnvironmentProduction Environment = "production"
)

func (e Environment) ToPointer() *Environment {
	return &e
}
func (e *Environment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sandbox":
		fallthrough
	case "production":
		*e = Environment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Environment: %v", v)
	}
}

type SourceCoinAPIUpdate struct {
	// API Key
	APIKey string `json:"api_key"`
	// The end date in ISO 8601 format. If not supplied, data will be returned
	// from the start date to the current time, or when the count of result
	// elements reaches its limit.
	//
	EndDate *string `json:"end_date,omitempty"`
	// The environment to use. Either sandbox or production.
	//
	Environment *Environment `default:"sandbox" json:"environment"`
	// The maximum number of elements to return. If not supplied, the default
	// is 100. For numbers larger than 100, each 100 items is counted as one
	// request for pricing purposes. Maximum value is 100000.
	//
	Limit *int64 `default:"100" json:"limit"`
	// The period to use. See the documentation for a list. https://docs.coinapi.io/#list-all-periods-get
	Period string `json:"period"`
	// The start date in ISO 8601 format.
	StartDate string `json:"start_date"`
	// The symbol ID to use. See the documentation for a list.
	// https://docs.coinapi.io/#list-all-symbols-get
	//
	SymbolID string `json:"symbol_id"`
}

func (s SourceCoinAPIUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceCoinAPIUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceCoinAPIUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceCoinAPIUpdate) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceCoinAPIUpdate) GetEnvironment() *Environment {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *SourceCoinAPIUpdate) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *SourceCoinAPIUpdate) GetPeriod() string {
	if o == nil {
		return ""
	}
	return o.Period
}

func (o *SourceCoinAPIUpdate) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

func (o *SourceCoinAPIUpdate) GetSymbolID() string {
	if o == nil {
		return ""
	}
	return o.SymbolID
}