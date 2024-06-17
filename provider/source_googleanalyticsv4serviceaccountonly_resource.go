// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_objectplanmodifier "github.com/howly-global/terraform-provider-airbyte/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/howly-global/terraform-provider-airbyte/planmodifiers/stringplanmodifier"
	tfTypes "github.com/howly-global/terraform-provider-airbyte/provider/types"
	"github.com/howly-global/terraform-provider-airbyte/sdk"
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/operations"
	"github.com/howly-global/terraform-provider-airbyte/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceGoogleAnalyticsV4ServiceAccountOnlyResource{}
var _ resource.ResourceWithImportState = &SourceGoogleAnalyticsV4ServiceAccountOnlyResource{}

func NewSourceGoogleAnalyticsV4ServiceAccountOnlyResource() resource.Resource {
	return &SourceGoogleAnalyticsV4ServiceAccountOnlyResource{}
}

// SourceGoogleAnalyticsV4ServiceAccountOnlyResource defines the resource implementation.
type SourceGoogleAnalyticsV4ServiceAccountOnlyResource struct {
	client *sdk.SDK
}

// SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel describes the resource data model.
type SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel struct {
	Configuration tfTypes.SourceGoogleAnalyticsV4ServiceAccountOnly `tfsdk:"configuration"`
	DefinitionID  types.String                                      `tfsdk:"definition_id"`
	Name          types.String                                      `tfsdk:"name"`
	SecretID      types.String                                      `tfsdk:"secret_id"`
	SourceID      types.String                                      `tfsdk:"source_id"`
	SourceType    types.String                                      `tfsdk:"source_type"`
	WorkspaceID   types.String                                      `tfsdk:"workspace_id"`
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_analytics_v4_service_account_only"
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleAnalyticsV4ServiceAccountOnly Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"service_account_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"credentials_json": schema.StringAttribute{
										Required:    true,
										Description: `The JSON key of the service account to use for authorization`,
									},
								},
							},
						},
						Description: `Credentials for the service`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"custom_reports": schema.StringAttribute{
						Optional:    true,
						Description: `A JSON array describing the custom reports you want to sync from Google Analytics. See <a href="https://docs.airbyte.com/integrations/sources/google-analytics-v4#data-processing-latency">the docs</a> for more information about the exact format you can use to fill out this field.`,
					},
					"end_date": schema.StringAttribute{
						Optional:    true,
						Description: `The date in the format YYYY-MM-DD. Any data after this date will not be replicated.`,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
					},
					"start_date": schema.StringAttribute{
						Required:    true,
						Description: `The date in the format YYYY-MM-DD. Any data before this date will not be replicated.`,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
					},
					"view_id": schema.StringAttribute{
						Required:    true,
						Description: `The ID for the Google Analytics View you want to fetch data from. This can be found from the <a href="https://ga-dev-tools.appspot.com/account-explorer/">Google Analytics Account Explorer</a>.`,
					},
					"window_in_days": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(1),
						Description: `The time increment used by the connector when requesting data from the Google Analytics API. More information is available in the <a href="https://docs.airbyte.com/integrations/sources/google-analytics-v4/#sampling-in-reports">the docs</a>. The bigger this value is, the faster the sync will be, but the more likely that sampling will be applied to your data, potentially causing inaccuracies in the returned results. We recommend setting this to 1 unless you have a hard requirement to make the sync faster at the expense of accuracy. The minimum allowed value for this field is 1, and the maximum is 364. . Default: 1`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed. `,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Name of the source e.g. dev-mysql-instance.`,
			},
			"secret_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed. `,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedSourceGoogleAnalyticsV4ServiceAccountOnlyCreateRequest()
	res, err := r.client.Sources.CreateSourceGoogleAnalyticsV4ServiceAccountOnly(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	sourceID := data.SourceID.ValueString()
	request1 := operations.GetSourceGoogleAnalyticsV4ServiceAccountOnlyRequest{
		SourceID: sourceID,
	}
	res1, err := r.client.Sources.GetSourceGoogleAnalyticsV4ServiceAccountOnly(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceGoogleAnalyticsV4ServiceAccountOnlyRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleAnalyticsV4ServiceAccountOnly(ctx, request)
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
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest := data.ToSharedSourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceGoogleAnalyticsV4ServiceAccountOnlyRequest{
		SourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest: sourceGoogleAnalyticsV4ServiceAccountOnlyPutRequest,
		SourceID: sourceID,
	}
	res, err := r.client.Sources.PutSourceGoogleAnalyticsV4ServiceAccountOnly(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	sourceId1 := data.SourceID.ValueString()
	request1 := operations.GetSourceGoogleAnalyticsV4ServiceAccountOnlyRequest{
		SourceID: sourceId1,
	}
	res1, err := r.client.Sources.GetSourceGoogleAnalyticsV4ServiceAccountOnly(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceGoogleAnalyticsV4ServiceAccountOnlyResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourceGoogleAnalyticsV4ServiceAccountOnlyRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceGoogleAnalyticsV4ServiceAccountOnly(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceGoogleAnalyticsV4ServiceAccountOnlyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
