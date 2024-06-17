// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationSnowflakeCortexCreateRequest struct {
	// The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
	// as well as to provide type safety for the configuration passed to the destination.
	//
	// The configuration model is composed of four parts:
	// * Processing configuration
	// * Embedding configuration
	// * Indexing configuration
	// * Advanced configuration
	//
	// Processing, embedding and advanced configuration are provided by this base class, while the indexing configuration is provided by the destination connector in the sub class.
	Configuration DestinationSnowflakeCortex `json:"configuration"`
	// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the destination e.g. dev-mysql-instance.
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationSnowflakeCortexCreateRequest) GetConfiguration() DestinationSnowflakeCortex {
	if o == nil {
		return DestinationSnowflakeCortex{}
	}
	return o.Configuration
}

func (o *DestinationSnowflakeCortexCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *DestinationSnowflakeCortexCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationSnowflakeCortexCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
