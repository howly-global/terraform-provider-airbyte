// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationConvexPutRequest struct {
	Configuration DestinationConvexUpdate `json:"configuration"`
	Name          string                  `json:"name"`
	WorkspaceID   string                  `json:"workspaceId"`
}

func (o *DestinationConvexPutRequest) GetConfiguration() DestinationConvexUpdate {
	if o == nil {
		return DestinationConvexUpdate{}
	}
	return o.Configuration
}

func (o *DestinationConvexPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationConvexPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}