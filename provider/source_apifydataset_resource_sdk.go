// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceApifyDatasetResourceModel) ToSharedSourceApifyDatasetCreateRequest() *shared.SourceApifyDatasetCreateRequest {
	datasetID := r.Configuration.DatasetID.ValueString()
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceApifyDataset{
		DatasetID: datasetID,
		Token:     token,
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
	out := shared.SourceApifyDatasetCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceApifyDatasetResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceApifyDatasetResourceModel) ToSharedSourceApifyDatasetPutRequest() *shared.SourceApifyDatasetPutRequest {
	datasetID := r.Configuration.DatasetID.ValueString()
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceApifyDatasetUpdate{
		DatasetID: datasetID,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceApifyDatasetPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}