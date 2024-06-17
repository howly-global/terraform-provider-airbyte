// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

// ConnectionResponse - Provides details of a single connection.
type ConnectionResponse struct {
	// A list of configured stream options for a connection.
	Configurations StreamConfigurations `json:"configurations"`
	ConnectionID   string               `json:"connectionId"`
	DataResidency  *GeographyEnum       `default:"auto" json:"dataResidency"`
	DestinationID  string               `json:"destinationId"`
	Name           string               `json:"name"`
	// Define the location where the data will be stored in the destination
	NamespaceDefinition *NamespaceDefinitionEnum `default:"destination" json:"namespaceDefinition"`
	NamespaceFormat     *string                  `json:"namespaceFormat,omitempty"`
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior *NonBreakingSchemaUpdatesBehaviorEnum `default:"ignore" json:"nonBreakingSchemaUpdatesBehavior"`
	Prefix                           *string                               `json:"prefix,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	Schedule    ConnectionScheduleResponse `json:"schedule"`
	SourceID    string                     `json:"sourceId"`
	Status      ConnectionStatusEnum       `json:"status"`
	WorkspaceID string                     `json:"workspaceId"`
}

func (c ConnectionResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionResponse) GetConfigurations() StreamConfigurations {
	if o == nil {
		return StreamConfigurations{}
	}
	return o.Configurations
}

func (o *ConnectionResponse) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ConnectionResponse) GetDataResidency() *GeographyEnum {
	if o == nil {
		return nil
	}
	return o.DataResidency
}

func (o *ConnectionResponse) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *ConnectionResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionResponse) GetNamespaceDefinition() *NamespaceDefinitionEnum {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionResponse) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionResponse) GetNonBreakingSchemaUpdatesBehavior() *NonBreakingSchemaUpdatesBehaviorEnum {
	if o == nil {
		return nil
	}
	return o.NonBreakingSchemaUpdatesBehavior
}

func (o *ConnectionResponse) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionResponse) GetSchedule() ConnectionScheduleResponse {
	if o == nil {
		return ConnectionScheduleResponse{}
	}
	return o.Schedule
}

func (o *ConnectionResponse) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *ConnectionResponse) GetStatus() ConnectionStatusEnum {
	if o == nil {
		return ConnectionStatusEnum("")
	}
	return o.Status
}

func (o *ConnectionResponse) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
