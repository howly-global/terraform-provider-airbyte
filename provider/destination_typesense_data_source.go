// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk"
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DestinationTypesenseDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationTypesenseDataSource{}

func NewDestinationTypesenseDataSource() datasource.DataSource {
	return &DestinationTypesenseDataSource{}
}

// DestinationTypesenseDataSource is the data source implementation.
type DestinationTypesenseDataSource struct {
	client *sdk.SDK
}

// DestinationTypesenseDataSourceModel describes the data model.
type DestinationTypesenseDataSourceModel struct {
	Configuration   types.String `tfsdk:"configuration"`
	DestinationID   types.String `tfsdk:"destination_id"`
	DestinationType types.String `tfsdk:"destination_type"`
	Name            types.String `tfsdk:"name"`
	WorkspaceID     types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationTypesenseDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_typesense"
}

// Schema defines the schema for the data source.
func (r *DestinationTypesenseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationTypesense DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.StringAttribute{
				Computed:    true,
				Description: `The values required to configure the destination. Parsed as JSON.`,
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationTypesenseDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationTypesenseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationTypesenseDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationTypesenseRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationTypesense(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
