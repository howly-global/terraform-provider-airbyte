// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceIterablePutRequest struct {
	Configuration SourceIterableUpdate `json:"configuration"`
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
}

func (o *SourceIterablePutRequest) GetConfiguration() SourceIterableUpdate {
	if o == nil {
		return SourceIterableUpdate{}
	}
	return o.Configuration
}

func (o *SourceIterablePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceIterablePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
