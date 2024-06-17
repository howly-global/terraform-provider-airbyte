// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceNotionPutRequest struct {
	Configuration SourceNotionUpdate `json:"configuration"`
	Name          string             `json:"name"`
	WorkspaceID   string             `json:"workspaceId"`
}

func (o *SourceNotionPutRequest) GetConfiguration() SourceNotionUpdate {
	if o == nil {
		return SourceNotionUpdate{}
	}
	return o.Configuration
}

func (o *SourceNotionPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceNotionPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
