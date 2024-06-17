// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type DevNull string

const (
	DevNullDevNull DevNull = "dev-null"
)

func (e DevNull) ToPointer() *DevNull {
	return &e
}
func (e *DevNull) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dev-null":
		*e = DevNull(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DevNull: %v", v)
	}
}

type DestinationDevNullTestDestinationType string

const (
	DestinationDevNullTestDestinationTypeSilent DestinationDevNullTestDestinationType = "SILENT"
)

func (e DestinationDevNullTestDestinationType) ToPointer() *DestinationDevNullTestDestinationType {
	return &e
}
func (e *DestinationDevNullTestDestinationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SILENT":
		*e = DestinationDevNullTestDestinationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullTestDestinationType: %v", v)
	}
}

type DestinationDevNullSilent struct {
	testDestinationType *DestinationDevNullTestDestinationType `const:"SILENT" json:"test_destination_type"`
}

func (d DestinationDevNullSilent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDevNullSilent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDevNullSilent) GetTestDestinationType() *DestinationDevNullTestDestinationType {
	return DestinationDevNullTestDestinationTypeSilent.ToPointer()
}

type DestinationDevNullTestDestinationUnionType string

const (
	DestinationDevNullTestDestinationUnionTypeDestinationDevNullSilent DestinationDevNullTestDestinationUnionType = "destination-dev-null_Silent"
)

// DestinationDevNullTestDestination - The type of destination to be used
type DestinationDevNullTestDestination struct {
	DestinationDevNullSilent *DestinationDevNullSilent

	Type DestinationDevNullTestDestinationUnionType
}

func CreateDestinationDevNullTestDestinationDestinationDevNullSilent(destinationDevNullSilent DestinationDevNullSilent) DestinationDevNullTestDestination {
	typ := DestinationDevNullTestDestinationUnionTypeDestinationDevNullSilent

	return DestinationDevNullTestDestination{
		DestinationDevNullSilent: &destinationDevNullSilent,
		Type:                     typ,
	}
}

func (u *DestinationDevNullTestDestination) UnmarshalJSON(data []byte) error {

	var destinationDevNullSilent DestinationDevNullSilent = DestinationDevNullSilent{}
	if err := utils.UnmarshalJSON(data, &destinationDevNullSilent, "", true, true); err == nil {
		u.DestinationDevNullSilent = &destinationDevNullSilent
		u.Type = DestinationDevNullTestDestinationUnionTypeDestinationDevNullSilent
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationDevNullTestDestination) MarshalJSON() ([]byte, error) {
	if u.DestinationDevNullSilent != nil {
		return utils.MarshalJSON(u.DestinationDevNullSilent, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationDevNull struct {
	destinationType DevNull `const:"dev-null" json:"destinationType"`
	// The type of destination to be used
	TestDestination DestinationDevNullTestDestination `json:"test_destination"`
}

func (d DestinationDevNull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDevNull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDevNull) GetDestinationType() DevNull {
	return DevNullDevNull
}

func (o *DestinationDevNull) GetTestDestination() DestinationDevNullTestDestination {
	if o == nil {
		return DestinationDevNullTestDestination{}
	}
	return o.TestDestination
}
