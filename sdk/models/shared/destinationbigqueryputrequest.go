// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationBigqueryPutRequest struct {
	Configuration DestinationBigqueryUpdate `json:"configuration"`
	Name          string                    `json:"name"`
	WorkspaceID   string                    `json:"workspaceId"`
}

func (o *DestinationBigqueryPutRequest) GetConfiguration() DestinationBigqueryUpdate {
	if o == nil {
		return DestinationBigqueryUpdate{}
	}
	return o.Configuration
}

func (o *DestinationBigqueryPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationBigqueryPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}