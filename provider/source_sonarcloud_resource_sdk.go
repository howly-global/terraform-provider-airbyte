// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	customTypes "github.com/howly-global/terraform-provider-airbyte/sdk/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSonarCloudResourceModel) ToSharedSourceSonarCloudCreateRequest() *shared.SourceSonarCloudCreateRequest {
	var componentKeys []interface{} = []interface{}{}
	for _, componentKeysItem := range r.Configuration.ComponentKeys {
		var componentKeysTmp interface{}
		_ = json.Unmarshal([]byte(componentKeysItem.ValueString()), &componentKeysTmp)
		componentKeys = append(componentKeys, componentKeysTmp)
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	organization := r.Configuration.Organization.ValueString()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	userToken := r.Configuration.UserToken.ValueString()
	configuration := shared.SourceSonarCloud{
		ComponentKeys: componentKeys,
		EndDate:       endDate,
		Organization:  organization,
		StartDate:     startDate,
		UserToken:     userToken,
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
	out := shared.SourceSonarCloudCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSonarCloudResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSonarCloudResourceModel) ToSharedSourceSonarCloudPutRequest() *shared.SourceSonarCloudPutRequest {
	var componentKeys []interface{} = []interface{}{}
	for _, componentKeysItem := range r.Configuration.ComponentKeys {
		var componentKeysTmp interface{}
		_ = json.Unmarshal([]byte(componentKeysItem.ValueString()), &componentKeysTmp)
		componentKeys = append(componentKeys, componentKeysTmp)
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	organization := r.Configuration.Organization.ValueString()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	userToken := r.Configuration.UserToken.ValueString()
	configuration := shared.SourceSonarCloudUpdate{
		ComponentKeys: componentKeys,
		EndDate:       endDate,
		Organization:  organization,
		StartDate:     startDate,
		UserToken:     userToken,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSonarCloudPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
