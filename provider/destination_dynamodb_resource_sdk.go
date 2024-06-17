// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDynamodbResourceModel) ToSharedDestinationDynamodbCreateRequest() *shared.DestinationDynamodbCreateRequest {
	accessKeyID := r.Configuration.AccessKeyID.ValueString()
	dynamodbEndpoint := new(string)
	if !r.Configuration.DynamodbEndpoint.IsUnknown() && !r.Configuration.DynamodbEndpoint.IsNull() {
		*dynamodbEndpoint = r.Configuration.DynamodbEndpoint.ValueString()
	} else {
		dynamodbEndpoint = nil
	}
	dynamodbRegion := new(shared.DestinationDynamodbDynamoDBRegion)
	if !r.Configuration.DynamodbRegion.IsUnknown() && !r.Configuration.DynamodbRegion.IsNull() {
		*dynamodbRegion = shared.DestinationDynamodbDynamoDBRegion(r.Configuration.DynamodbRegion.ValueString())
	} else {
		dynamodbRegion = nil
	}
	dynamodbTableNamePrefix := r.Configuration.DynamodbTableNamePrefix.ValueString()
	secretAccessKey := r.Configuration.SecretAccessKey.ValueString()
	configuration := shared.DestinationDynamodb{
		AccessKeyID:             accessKeyID,
		DynamodbEndpoint:        dynamodbEndpoint,
		DynamodbRegion:          dynamodbRegion,
		DynamodbTableNamePrefix: dynamodbTableNamePrefix,
		SecretAccessKey:         secretAccessKey,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDynamodbCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDynamodbResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationDynamodbResourceModel) ToSharedDestinationDynamodbPutRequest() *shared.DestinationDynamodbPutRequest {
	accessKeyID := r.Configuration.AccessKeyID.ValueString()
	dynamodbEndpoint := new(string)
	if !r.Configuration.DynamodbEndpoint.IsUnknown() && !r.Configuration.DynamodbEndpoint.IsNull() {
		*dynamodbEndpoint = r.Configuration.DynamodbEndpoint.ValueString()
	} else {
		dynamodbEndpoint = nil
	}
	dynamodbRegion := new(shared.DynamoDBRegion)
	if !r.Configuration.DynamodbRegion.IsUnknown() && !r.Configuration.DynamodbRegion.IsNull() {
		*dynamodbRegion = shared.DynamoDBRegion(r.Configuration.DynamodbRegion.ValueString())
	} else {
		dynamodbRegion = nil
	}
	dynamodbTableNamePrefix := r.Configuration.DynamodbTableNamePrefix.ValueString()
	secretAccessKey := r.Configuration.SecretAccessKey.ValueString()
	configuration := shared.DestinationDynamodbUpdate{
		AccessKeyID:             accessKeyID,
		DynamodbEndpoint:        dynamodbEndpoint,
		DynamodbRegion:          dynamodbRegion,
		DynamodbTableNamePrefix: dynamodbTableNamePrefix,
		SecretAccessKey:         secretAccessKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDynamodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
