// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	customTypes "github.com/howly-global/terraform-provider-airbyte/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePrestashopResourceModel) ToSharedSourcePrestashopCreateRequest() *shared.SourcePrestashopCreateRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourcePrestashop{
		AccessKey: accessKey,
		StartDate: startDate,
		URL:       url,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePrestashopCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePrestashopResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourcePrestashopResourceModel) ToSharedSourcePrestashopPutRequest() *shared.SourcePrestashopPutRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourcePrestashopUpdate{
		AccessKey: accessKey,
		StartDate: startDate,
		URL:       url,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePrestashopPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}