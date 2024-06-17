// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceAmplitudeResourceModel) ToSharedSourceAmplitudeCreateRequest() *shared.SourceAmplitudeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dataRegion := new(shared.SourceAmplitudeDataRegion)
	if !r.Configuration.DataRegion.IsUnknown() && !r.Configuration.DataRegion.IsNull() {
		*dataRegion = shared.SourceAmplitudeDataRegion(r.Configuration.DataRegion.ValueString())
	} else {
		dataRegion = nil
	}
	requestTimeRange := new(int64)
	if !r.Configuration.RequestTimeRange.IsUnknown() && !r.Configuration.RequestTimeRange.IsNull() {
		*requestTimeRange = r.Configuration.RequestTimeRange.ValueInt64()
	} else {
		requestTimeRange = nil
	}
	secretKey := r.Configuration.SecretKey.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceAmplitude{
		APIKey:           apiKey,
		DataRegion:       dataRegion,
		RequestTimeRange: requestTimeRange,
		SecretKey:        secretKey,
		StartDate:        startDate,
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
	out := shared.SourceAmplitudeCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmplitudeResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceAmplitudeResourceModel) ToSharedSourceAmplitudePutRequest() *shared.SourceAmplitudePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dataRegion := new(shared.DataRegion)
	if !r.Configuration.DataRegion.IsUnknown() && !r.Configuration.DataRegion.IsNull() {
		*dataRegion = shared.DataRegion(r.Configuration.DataRegion.ValueString())
	} else {
		dataRegion = nil
	}
	requestTimeRange := new(int64)
	if !r.Configuration.RequestTimeRange.IsUnknown() && !r.Configuration.RequestTimeRange.IsNull() {
		*requestTimeRange = r.Configuration.RequestTimeRange.ValueInt64()
	} else {
		requestTimeRange = nil
	}
	secretKey := r.Configuration.SecretKey.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceAmplitudeUpdate{
		APIKey:           apiKey,
		DataRegion:       dataRegion,
		RequestTimeRange: requestTimeRange,
		SecretKey:        secretKey,
		StartDate:        startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmplitudePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
