// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSftpPutRequest struct {
	Configuration SourceSftpUpdate `json:"configuration"`
	Name          string           `json:"name"`
	WorkspaceID   string           `json:"workspaceId"`
}

func (o *SourceSftpPutRequest) GetConfiguration() SourceSftpUpdate {
	if o == nil {
		return SourceSftpUpdate{}
	}
	return o.Configuration
}

func (o *SourceSftpPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSftpPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
