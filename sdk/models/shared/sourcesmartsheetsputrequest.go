// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSmartsheetsPutRequest struct {
	Configuration SourceSmartsheetsUpdate `json:"configuration"`
	Name          string                  `json:"name"`
	WorkspaceID   string                  `json:"workspaceId"`
}

func (o *SourceSmartsheetsPutRequest) GetConfiguration() SourceSmartsheetsUpdate {
	if o == nil {
		return SourceSmartsheetsUpdate{}
	}
	return o.Configuration
}

func (o *SourceSmartsheetsPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSmartsheetsPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
