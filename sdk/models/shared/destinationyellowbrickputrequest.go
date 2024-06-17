// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationYellowbrickPutRequest struct {
	Configuration DestinationYellowbrickUpdate `json:"configuration"`
	Name          string                       `json:"name"`
	WorkspaceID   string                       `json:"workspaceId"`
}

func (o *DestinationYellowbrickPutRequest) GetConfiguration() DestinationYellowbrickUpdate {
	if o == nil {
		return DestinationYellowbrickUpdate{}
	}
	return o.Configuration
}

func (o *DestinationYellowbrickPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationYellowbrickPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
