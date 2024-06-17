// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceCoinmarketcapResourceModel) ToSharedSourceCoinmarketcapCreateRequest() *shared.SourceCoinmarketcapCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dataType := shared.SourceCoinmarketcapDataType(r.Configuration.DataType.ValueString())
	var symbols []string = []string{}
	for _, symbolsItem := range r.Configuration.Symbols {
		symbols = append(symbols, symbolsItem.ValueString())
	}
	configuration := shared.SourceCoinmarketcap{
		APIKey:   apiKey,
		DataType: dataType,
		Symbols:  symbols,
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
	out := shared.SourceCoinmarketcapCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceCoinmarketcapResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceCoinmarketcapResourceModel) ToSharedSourceCoinmarketcapPutRequest() *shared.SourceCoinmarketcapPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dataType := shared.DataType(r.Configuration.DataType.ValueString())
	var symbols []string = []string{}
	for _, symbolsItem := range r.Configuration.Symbols {
		symbols = append(symbols, symbolsItem.ValueString())
	}
	configuration := shared.SourceCoinmarketcapUpdate{
		APIKey:   apiKey,
		DataType: dataType,
		Symbols:  symbols,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceCoinmarketcapPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
