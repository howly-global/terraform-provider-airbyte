// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceZohoCrmCreateRequest struct {
	Configuration SourceZohoCrm `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceZohoCrmCreateRequest) GetConfiguration() SourceZohoCrm {
	if o == nil {
		return SourceZohoCrm{}
	}
	return o.Configuration
}

func (o *SourceZohoCrmCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *SourceZohoCrmCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceZohoCrmCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceZohoCrmCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}