// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMetabaseResourceModel) ToSharedSourceMetabaseCreateRequest() *shared.SourceMetabaseCreateRequest {
	instanceAPIURL := r.Configuration.InstanceAPIURL.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	sessionToken := new(string)
	if !r.Configuration.SessionToken.IsUnknown() && !r.Configuration.SessionToken.IsNull() {
		*sessionToken = r.Configuration.SessionToken.ValueString()
	} else {
		sessionToken = nil
	}
	username := new(string)
	if !r.Configuration.Username.IsUnknown() && !r.Configuration.Username.IsNull() {
		*username = r.Configuration.Username.ValueString()
	} else {
		username = nil
	}
	configuration := shared.SourceMetabase{
		InstanceAPIURL: instanceAPIURL,
		Password:       password,
		SessionToken:   sessionToken,
		Username:       username,
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
	out := shared.SourceMetabaseCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMetabaseResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceMetabaseResourceModel) ToSharedSourceMetabasePutRequest() *shared.SourceMetabasePutRequest {
	instanceAPIURL := r.Configuration.InstanceAPIURL.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	sessionToken := new(string)
	if !r.Configuration.SessionToken.IsUnknown() && !r.Configuration.SessionToken.IsNull() {
		*sessionToken = r.Configuration.SessionToken.ValueString()
	} else {
		sessionToken = nil
	}
	username := new(string)
	if !r.Configuration.Username.IsUnknown() && !r.Configuration.Username.IsNull() {
		*username = r.Configuration.Username.ValueString()
	} else {
		username = nil
	}
	configuration := shared.SourceMetabaseUpdate{
		InstanceAPIURL: instanceAPIURL,
		Password:       password,
		SessionToken:   sessionToken,
		Username:       username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMetabasePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}