// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePipedrivePutRequest struct {
	Configuration SourcePipedriveUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourcePipedrivePutRequest) GetConfiguration() SourcePipedriveUpdate {
	if o == nil {
		return SourcePipedriveUpdate{}
	}
	return o.Configuration
}

func (o *SourcePipedrivePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePipedrivePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
