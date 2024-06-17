// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDevNullResourceModel) ToSharedDestinationDevNullCreateRequest() *shared.DestinationDevNullCreateRequest {
	var testDestination shared.DestinationDevNullTestDestination
	var destinationDevNullSilent *shared.DestinationDevNullSilent
	if r.Configuration.TestDestination.Silent != nil {
		destinationDevNullSilent = &shared.DestinationDevNullSilent{}
	}
	if destinationDevNullSilent != nil {
		testDestination = shared.DestinationDevNullTestDestination{
			DestinationDevNullSilent: destinationDevNullSilent,
		}
	}
	configuration := shared.DestinationDevNull{
		TestDestination: testDestination,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDevNullCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDevNullResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationDevNullResourceModel) ToSharedDestinationDevNullPutRequest() *shared.DestinationDevNullPutRequest {
	var testDestination shared.TestDestination
	var silent *shared.Silent
	if r.Configuration.TestDestination.Silent != nil {
		silent = &shared.Silent{}
	}
	if silent != nil {
		testDestination = shared.TestDestination{
			Silent: silent,
		}
	}
	configuration := shared.DestinationDevNullUpdate{
		TestDestination: testDestination,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDevNullPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}