// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceOmnisendPutRequest struct {
	Configuration SourceOmnisendUpdate `json:"configuration"`
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
}

func (o *SourceOmnisendPutRequest) GetConfiguration() SourceOmnisendUpdate {
	if o == nil {
		return SourceOmnisendUpdate{}
	}
	return o.Configuration
}

func (o *SourceOmnisendPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceOmnisendPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}