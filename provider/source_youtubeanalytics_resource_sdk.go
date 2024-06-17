// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceYoutubeAnalyticsResourceModel) ToSharedSourceYoutubeAnalyticsCreateRequest() *shared.SourceYoutubeAnalyticsCreateRequest {
	var additionalProperties interface{}
	if !r.Configuration.Credentials.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AdditionalProperties.IsNull() {
		_ = json.Unmarshal([]byte(r.Configuration.Credentials.AdditionalProperties.ValueString()), &additionalProperties)
	}
	clientID := r.Configuration.Credentials.ClientID.ValueString()
	clientSecret := r.Configuration.Credentials.ClientSecret.ValueString()
	refreshToken := r.Configuration.Credentials.RefreshToken.ValueString()
	credentials := shared.SourceYoutubeAnalyticsAuthenticateViaOAuth20{
		AdditionalProperties: additionalProperties,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		RefreshToken:         refreshToken,
	}
	configuration := shared.SourceYoutubeAnalytics{
		Credentials: credentials,
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
	out := shared.SourceYoutubeAnalyticsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceYoutubeAnalyticsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceYoutubeAnalyticsResourceModel) ToSharedSourceYoutubeAnalyticsPutRequest() *shared.SourceYoutubeAnalyticsPutRequest {
	var additionalProperties interface{}
	if !r.Configuration.Credentials.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AdditionalProperties.IsNull() {
		_ = json.Unmarshal([]byte(r.Configuration.Credentials.AdditionalProperties.ValueString()), &additionalProperties)
	}
	clientID := r.Configuration.Credentials.ClientID.ValueString()
	clientSecret := r.Configuration.Credentials.ClientSecret.ValueString()
	refreshToken := r.Configuration.Credentials.RefreshToken.ValueString()
	credentials := shared.AuthenticateViaOAuth20{
		AdditionalProperties: additionalProperties,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		RefreshToken:         refreshToken,
	}
	configuration := shared.SourceYoutubeAnalyticsUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceYoutubeAnalyticsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
