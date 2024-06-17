// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationClickhouseCreateRequest struct {
	Configuration DestinationClickhouse `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationClickhouseCreateRequest) GetConfiguration() DestinationClickhouse {
	if o == nil {
		return DestinationClickhouse{}
	}
	return o.Configuration
}

func (o *DestinationClickhouseCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationClickhouseCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationClickhouseCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
