// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type Xkcd string

const (
	XkcdXkcd Xkcd = "xkcd"
)

func (e Xkcd) ToPointer() *Xkcd {
	return &e
}
func (e *Xkcd) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "xkcd":
		*e = Xkcd(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Xkcd: %v", v)
	}
}

type SourceXkcd struct {
	sourceType *Xkcd `const:"xkcd" json:"sourceType,omitempty"`
}

func (s SourceXkcd) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceXkcd) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceXkcd) GetSourceType() *Xkcd {
	return XkcdXkcd.ToPointer()
}