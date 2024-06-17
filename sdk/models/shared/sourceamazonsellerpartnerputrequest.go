// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceAmazonSellerPartnerPutRequest struct {
	Configuration SourceAmazonSellerPartnerUpdate `json:"configuration"`
	Name          string                          `json:"name"`
	WorkspaceID   string                          `json:"workspaceId"`
}

func (o *SourceAmazonSellerPartnerPutRequest) GetConfiguration() SourceAmazonSellerPartnerUpdate {
	if o == nil {
		return SourceAmazonSellerPartnerUpdate{}
	}
	return o.Configuration
}

func (o *SourceAmazonSellerPartnerPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAmazonSellerPartnerPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
