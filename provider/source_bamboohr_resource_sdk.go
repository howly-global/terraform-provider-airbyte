// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceBambooHrResourceModel) ToSharedSourceBambooHrCreateRequest() *shared.SourceBambooHrCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	customReportsFields := new(string)
	if !r.Configuration.CustomReportsFields.IsUnknown() && !r.Configuration.CustomReportsFields.IsNull() {
		*customReportsFields = r.Configuration.CustomReportsFields.ValueString()
	} else {
		customReportsFields = nil
	}
	customReportsIncludeDefaultFields := new(bool)
	if !r.Configuration.CustomReportsIncludeDefaultFields.IsUnknown() && !r.Configuration.CustomReportsIncludeDefaultFields.IsNull() {
		*customReportsIncludeDefaultFields = r.Configuration.CustomReportsIncludeDefaultFields.ValueBool()
	} else {
		customReportsIncludeDefaultFields = nil
	}
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceBambooHr{
		APIKey:                            apiKey,
		CustomReportsFields:               customReportsFields,
		CustomReportsIncludeDefaultFields: customReportsIncludeDefaultFields,
		Subdomain:                         subdomain,
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
	out := shared.SourceBambooHrCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBambooHrResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceBambooHrResourceModel) ToSharedSourceBambooHrPutRequest() *shared.SourceBambooHrPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	customReportsFields := new(string)
	if !r.Configuration.CustomReportsFields.IsUnknown() && !r.Configuration.CustomReportsFields.IsNull() {
		*customReportsFields = r.Configuration.CustomReportsFields.ValueString()
	} else {
		customReportsFields = nil
	}
	customReportsIncludeDefaultFields := new(bool)
	if !r.Configuration.CustomReportsIncludeDefaultFields.IsUnknown() && !r.Configuration.CustomReportsIncludeDefaultFields.IsNull() {
		*customReportsIncludeDefaultFields = r.Configuration.CustomReportsIncludeDefaultFields.ValueBool()
	} else {
		customReportsIncludeDefaultFields = nil
	}
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceBambooHrUpdate{
		APIKey:                            apiKey,
		CustomReportsFields:               customReportsFields,
		CustomReportsIncludeDefaultFields: customReportsIncludeDefaultFields,
		Subdomain:                         subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBambooHrPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
