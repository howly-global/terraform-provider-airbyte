// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSmartengageUpdate struct {
	// API Key
	APIKey string `json:"api_key"`
}

func (o *SourceSmartengageUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}