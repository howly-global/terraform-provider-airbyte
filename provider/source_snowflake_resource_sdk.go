// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSnowflakeResourceModel) ToSharedSourceSnowflakeCreateRequest() *shared.SourceSnowflakeCreateRequest {
	var credentials *shared.SourceSnowflakeAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceSnowflakeOAuth20 *shared.SourceSnowflakeOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			refreshToken := new(string)
			if !r.Configuration.Credentials.OAuth20.RefreshToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.RefreshToken.IsNull() {
				*refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			} else {
				refreshToken = nil
			}
			sourceSnowflakeOAuth20 = &shared.SourceSnowflakeOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceSnowflakeOAuth20 != nil {
			credentials = &shared.SourceSnowflakeAuthorizationMethod{
				SourceSnowflakeOAuth20: sourceSnowflakeOAuth20,
			}
		}
		var sourceSnowflakeUsernameAndPassword *shared.SourceSnowflakeUsernameAndPassword
		if r.Configuration.Credentials.UsernameAndPassword != nil {
			password := r.Configuration.Credentials.UsernameAndPassword.Password.ValueString()
			username := r.Configuration.Credentials.UsernameAndPassword.Username.ValueString()
			sourceSnowflakeUsernameAndPassword = &shared.SourceSnowflakeUsernameAndPassword{
				Password: password,
				Username: username,
			}
		}
		if sourceSnowflakeUsernameAndPassword != nil {
			credentials = &shared.SourceSnowflakeAuthorizationMethod{
				SourceSnowflakeUsernameAndPassword: sourceSnowflakeUsernameAndPassword,
			}
		}
	}
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	role := r.Configuration.Role.ValueString()
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.SourceSnowflake{
		Credentials:   credentials,
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Role:          role,
		Schema:        schema,
		Warehouse:     warehouse,
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
	out := shared.SourceSnowflakeCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSnowflakeResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSnowflakeResourceModel) ToSharedSourceSnowflakePutRequest() *shared.SourceSnowflakePutRequest {
	var credentials *shared.SourceSnowflakeUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceSnowflakeUpdateOAuth20 *shared.SourceSnowflakeUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			accessToken := new(string)
			if !r.Configuration.Credentials.OAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.AccessToken.IsNull() {
				*accessToken = r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			} else {
				accessToken = nil
			}
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			refreshToken := new(string)
			if !r.Configuration.Credentials.OAuth20.RefreshToken.IsUnknown() && !r.Configuration.Credentials.OAuth20.RefreshToken.IsNull() {
				*refreshToken = r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			} else {
				refreshToken = nil
			}
			sourceSnowflakeUpdateOAuth20 = &shared.SourceSnowflakeUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceSnowflakeUpdateOAuth20 != nil {
			credentials = &shared.SourceSnowflakeUpdateAuthorizationMethod{
				SourceSnowflakeUpdateOAuth20: sourceSnowflakeUpdateOAuth20,
			}
		}
		var sourceSnowflakeUpdateUsernameAndPassword *shared.SourceSnowflakeUpdateUsernameAndPassword
		if r.Configuration.Credentials.UsernameAndPassword != nil {
			password := r.Configuration.Credentials.UsernameAndPassword.Password.ValueString()
			username := r.Configuration.Credentials.UsernameAndPassword.Username.ValueString()
			sourceSnowflakeUpdateUsernameAndPassword = &shared.SourceSnowflakeUpdateUsernameAndPassword{
				Password: password,
				Username: username,
			}
		}
		if sourceSnowflakeUpdateUsernameAndPassword != nil {
			credentials = &shared.SourceSnowflakeUpdateAuthorizationMethod{
				SourceSnowflakeUpdateUsernameAndPassword: sourceSnowflakeUpdateUsernameAndPassword,
			}
		}
	}
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	role := r.Configuration.Role.ValueString()
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.SourceSnowflakeUpdate{
		Credentials:   credentials,
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Role:          role,
		Schema:        schema,
		Warehouse:     warehouse,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSnowflakePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
