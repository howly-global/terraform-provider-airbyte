// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type Marketo string

const (
	MarketoMarketo Marketo = "marketo"
)

func (e Marketo) ToPointer() *Marketo {
	return &e
}
func (e *Marketo) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketo":
		*e = Marketo(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Marketo: %v", v)
	}
}

type SourceMarketo struct {
	// The Client ID of your Marketo developer application. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
	ClientID string `json:"client_id"`
	// The Client Secret of your Marketo developer application. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
	ClientSecret string `json:"client_secret"`
	// Your Marketo Base URL. See <a href="https://docs.airbyte.com/integrations/sources/marketo"> the docs </a> for info on how to obtain this.
	DomainURL  string  `json:"domain_url"`
	sourceType Marketo `const:"marketo" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceMarketo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMarketo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMarketo) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceMarketo) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceMarketo) GetDomainURL() string {
	if o == nil {
		return ""
	}
	return o.DomainURL
}

func (o *SourceMarketo) GetSourceType() Marketo {
	return MarketoMarketo
}

func (o *SourceMarketo) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
