// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOrbResourceModel) ToSharedSourceOrbCreateRequest() *shared.SourceOrbCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	var numericEventPropertiesKeys []string = []string{}
	for _, numericEventPropertiesKeysItem := range r.Configuration.NumericEventPropertiesKeys {
		numericEventPropertiesKeys = append(numericEventPropertiesKeys, numericEventPropertiesKeysItem.ValueString())
	}
	planID := new(string)
	if !r.Configuration.PlanID.IsUnknown() && !r.Configuration.PlanID.IsNull() {
		*planID = r.Configuration.PlanID.ValueString()
	} else {
		planID = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	var stringEventPropertiesKeys []string = []string{}
	for _, stringEventPropertiesKeysItem := range r.Configuration.StringEventPropertiesKeys {
		stringEventPropertiesKeys = append(stringEventPropertiesKeys, stringEventPropertiesKeysItem.ValueString())
	}
	subscriptionUsageGroupingKey := new(string)
	if !r.Configuration.SubscriptionUsageGroupingKey.IsUnknown() && !r.Configuration.SubscriptionUsageGroupingKey.IsNull() {
		*subscriptionUsageGroupingKey = r.Configuration.SubscriptionUsageGroupingKey.ValueString()
	} else {
		subscriptionUsageGroupingKey = nil
	}
	configuration := shared.SourceOrb{
		APIKey:                       apiKey,
		EndDate:                      endDate,
		LookbackWindowDays:           lookbackWindowDays,
		NumericEventPropertiesKeys:   numericEventPropertiesKeys,
		PlanID:                       planID,
		StartDate:                    startDate,
		StringEventPropertiesKeys:    stringEventPropertiesKeys,
		SubscriptionUsageGroupingKey: subscriptionUsageGroupingKey,
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
	out := shared.SourceOrbCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceOrbResourceModel) ToSharedSourceOrbPutRequest() *shared.SourceOrbPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	var numericEventPropertiesKeys []string = []string{}
	for _, numericEventPropertiesKeysItem := range r.Configuration.NumericEventPropertiesKeys {
		numericEventPropertiesKeys = append(numericEventPropertiesKeys, numericEventPropertiesKeysItem.ValueString())
	}
	planID := new(string)
	if !r.Configuration.PlanID.IsUnknown() && !r.Configuration.PlanID.IsNull() {
		*planID = r.Configuration.PlanID.ValueString()
	} else {
		planID = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	var stringEventPropertiesKeys []string = []string{}
	for _, stringEventPropertiesKeysItem := range r.Configuration.StringEventPropertiesKeys {
		stringEventPropertiesKeys = append(stringEventPropertiesKeys, stringEventPropertiesKeysItem.ValueString())
	}
	subscriptionUsageGroupingKey := new(string)
	if !r.Configuration.SubscriptionUsageGroupingKey.IsUnknown() && !r.Configuration.SubscriptionUsageGroupingKey.IsNull() {
		*subscriptionUsageGroupingKey = r.Configuration.SubscriptionUsageGroupingKey.ValueString()
	} else {
		subscriptionUsageGroupingKey = nil
	}
	configuration := shared.SourceOrbUpdate{
		APIKey:                       apiKey,
		EndDate:                      endDate,
		LookbackWindowDays:           lookbackWindowDays,
		NumericEventPropertiesKeys:   numericEventPropertiesKeys,
		PlanID:                       planID,
		StartDate:                    startDate,
		StringEventPropertiesKeys:    stringEventPropertiesKeys,
		SubscriptionUsageGroupingKey: subscriptionUsageGroupingKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOrbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
