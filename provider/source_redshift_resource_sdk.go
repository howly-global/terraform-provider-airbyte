// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRedshiftResourceModel) ToSharedSourceRedshiftCreateRequest() *shared.SourceRedshiftCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	var schemas []string = []string{}
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceRedshift{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schemas:       schemas,
		Username:      username,
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
	out := shared.SourceRedshiftCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRedshiftResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceRedshiftResourceModel) ToSharedSourceRedshiftPutRequest() *shared.SourceRedshiftPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	var schemas []string = []string{}
	for _, schemasItem := range r.Configuration.Schemas {
		schemas = append(schemas, schemasItem.ValueString())
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceRedshiftUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schemas:       schemas,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRedshiftPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
