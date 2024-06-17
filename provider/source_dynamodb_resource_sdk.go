// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDynamodbResourceModel) ToSharedSourceDynamodbCreateRequest() *shared.SourceDynamodbCreateRequest {
	var credentials *shared.SourceDynamodbCredentials
	if r.Configuration.Credentials != nil {
		var sourceDynamodbAuthenticateViaAccessKeys *shared.SourceDynamodbAuthenticateViaAccessKeys
		if r.Configuration.Credentials.AuthenticateViaAccessKeys != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.AuthenticateViaAccessKeys.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaAccessKeys.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.AuthenticateViaAccessKeys.AdditionalProperties.ValueString()), &additionalProperties)
			}
			accessKeyID := r.Configuration.Credentials.AuthenticateViaAccessKeys.AccessKeyID.ValueString()
			secretAccessKey := r.Configuration.Credentials.AuthenticateViaAccessKeys.SecretAccessKey.ValueString()
			sourceDynamodbAuthenticateViaAccessKeys = &shared.SourceDynamodbAuthenticateViaAccessKeys{
				AdditionalProperties: additionalProperties,
				AccessKeyID:          accessKeyID,
				SecretAccessKey:      secretAccessKey,
			}
		}
		if sourceDynamodbAuthenticateViaAccessKeys != nil {
			credentials = &shared.SourceDynamodbCredentials{
				SourceDynamodbAuthenticateViaAccessKeys: sourceDynamodbAuthenticateViaAccessKeys,
			}
		}
		var sourceDynamodbRoleBasedAuthentication *shared.SourceDynamodbRoleBasedAuthentication
		if r.Configuration.Credentials.RoleBasedAuthentication != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.RoleBasedAuthentication.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.RoleBasedAuthentication.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.RoleBasedAuthentication.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceDynamodbRoleBasedAuthentication = &shared.SourceDynamodbRoleBasedAuthentication{
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceDynamodbRoleBasedAuthentication != nil {
			credentials = &shared.SourceDynamodbCredentials{
				SourceDynamodbRoleBasedAuthentication: sourceDynamodbRoleBasedAuthentication,
			}
		}
	}
	endpoint := new(string)
	if !r.Configuration.Endpoint.IsUnknown() && !r.Configuration.Endpoint.IsNull() {
		*endpoint = r.Configuration.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	ignoreMissingReadPermissionsTables := new(bool)
	if !r.Configuration.IgnoreMissingReadPermissionsTables.IsUnknown() && !r.Configuration.IgnoreMissingReadPermissionsTables.IsNull() {
		*ignoreMissingReadPermissionsTables = r.Configuration.IgnoreMissingReadPermissionsTables.ValueBool()
	} else {
		ignoreMissingReadPermissionsTables = nil
	}
	region := new(shared.SourceDynamodbDynamodbRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceDynamodbDynamodbRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	reservedAttributeNames := new(string)
	if !r.Configuration.ReservedAttributeNames.IsUnknown() && !r.Configuration.ReservedAttributeNames.IsNull() {
		*reservedAttributeNames = r.Configuration.ReservedAttributeNames.ValueString()
	} else {
		reservedAttributeNames = nil
	}
	configuration := shared.SourceDynamodb{
		Credentials:                        credentials,
		Endpoint:                           endpoint,
		IgnoreMissingReadPermissionsTables: ignoreMissingReadPermissionsTables,
		Region:                             region,
		ReservedAttributeNames:             reservedAttributeNames,
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
	out := shared.SourceDynamodbCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDynamodbResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceDynamodbResourceModel) ToSharedSourceDynamodbPutRequest() *shared.SourceDynamodbPutRequest {
	var credentials *shared.Credentials
	if r.Configuration.Credentials != nil {
		var authenticateViaAccessKeys *shared.AuthenticateViaAccessKeys
		if r.Configuration.Credentials.AuthenticateViaAccessKeys != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.AuthenticateViaAccessKeys.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaAccessKeys.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.AuthenticateViaAccessKeys.AdditionalProperties.ValueString()), &additionalProperties)
			}
			accessKeyID := r.Configuration.Credentials.AuthenticateViaAccessKeys.AccessKeyID.ValueString()
			secretAccessKey := r.Configuration.Credentials.AuthenticateViaAccessKeys.SecretAccessKey.ValueString()
			authenticateViaAccessKeys = &shared.AuthenticateViaAccessKeys{
				AdditionalProperties: additionalProperties,
				AccessKeyID:          accessKeyID,
				SecretAccessKey:      secretAccessKey,
			}
		}
		if authenticateViaAccessKeys != nil {
			credentials = &shared.Credentials{
				AuthenticateViaAccessKeys: authenticateViaAccessKeys,
			}
		}
		var roleBasedAuthentication *shared.RoleBasedAuthentication
		if r.Configuration.Credentials.RoleBasedAuthentication != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.RoleBasedAuthentication.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.RoleBasedAuthentication.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.RoleBasedAuthentication.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			roleBasedAuthentication = &shared.RoleBasedAuthentication{
				AdditionalProperties: additionalProperties1,
			}
		}
		if roleBasedAuthentication != nil {
			credentials = &shared.Credentials{
				RoleBasedAuthentication: roleBasedAuthentication,
			}
		}
	}
	endpoint := new(string)
	if !r.Configuration.Endpoint.IsUnknown() && !r.Configuration.Endpoint.IsNull() {
		*endpoint = r.Configuration.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	ignoreMissingReadPermissionsTables := new(bool)
	if !r.Configuration.IgnoreMissingReadPermissionsTables.IsUnknown() && !r.Configuration.IgnoreMissingReadPermissionsTables.IsNull() {
		*ignoreMissingReadPermissionsTables = r.Configuration.IgnoreMissingReadPermissionsTables.ValueBool()
	} else {
		ignoreMissingReadPermissionsTables = nil
	}
	region := new(shared.SourceDynamodbUpdateDynamodbRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceDynamodbUpdateDynamodbRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	reservedAttributeNames := new(string)
	if !r.Configuration.ReservedAttributeNames.IsUnknown() && !r.Configuration.ReservedAttributeNames.IsNull() {
		*reservedAttributeNames = r.Configuration.ReservedAttributeNames.ValueString()
	} else {
		reservedAttributeNames = nil
	}
	configuration := shared.SourceDynamodbUpdate{
		Credentials:                        credentials,
		Endpoint:                           endpoint,
		IgnoreMissingReadPermissionsTables: ignoreMissingReadPermissionsTables,
		Region:                             region,
		ReservedAttributeNames:             reservedAttributeNames,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDynamodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
