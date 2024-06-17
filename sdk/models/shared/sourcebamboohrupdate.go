// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type SourceBambooHrUpdate struct {
	// Api key of bamboo hr
	APIKey string `json:"api_key"`
	// Comma-separated list of fields to include in custom reports.
	CustomReportsFields *string `default:"" json:"custom_reports_fields"`
	// If true, the custom reports endpoint will include the default fields defined here: https://documentation.bamboohr.com/docs/list-of-field-names.
	CustomReportsIncludeDefaultFields *bool `default:"true" json:"custom_reports_include_default_fields"`
	// Sub Domain of bamboo hr
	Subdomain string `json:"subdomain"`
}

func (s SourceBambooHrUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceBambooHrUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceBambooHrUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceBambooHrUpdate) GetCustomReportsFields() *string {
	if o == nil {
		return nil
	}
	return o.CustomReportsFields
}

func (o *SourceBambooHrUpdate) GetCustomReportsIncludeDefaultFields() *bool {
	if o == nil {
		return nil
	}
	return o.CustomReportsIncludeDefaultFields
}

func (o *SourceBambooHrUpdate) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
