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
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationSnowflakeCortexResource{}
var _ resource.ResourceWithImportState = &DestinationSnowflakeCortexResource{}

func NewDestinationSnowflakeCortexResource() resource.Resource {
	return &DestinationSnowflakeCortexResource{}
}

// DestinationSnowflakeCortexResource defines the resource implementation.
type DestinationSnowflakeCortexResource struct {
	client *sdk.SDK
}

// DestinationSnowflakeCortexResourceModel describes the resource data model.
type DestinationSnowflakeCortexResourceModel struct {
	Configuration   tfTypes.DestinationSnowflakeCortex `tfsdk:"configuration"`
	DefinitionID    types.String                       `tfsdk:"definition_id"`
	DestinationID   types.String                       `tfsdk:"destination_id"`
	DestinationType types.String                       `tfsdk:"destination_type"`
	Name            types.String                       `tfsdk:"name"`
	WorkspaceID     types.String                       `tfsdk:"workspace_id"`
}

func (r *DestinationSnowflakeCortexResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_snowflake_cortex"
}

func (r *DestinationSnowflakeCortexResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationSnowflakeCortex Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"embedding": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"azure_open_ai": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"api_base": schema.StringAttribute{
										Required:    true,
										Description: `The base URL for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource`,
									},
									"deployment": schema.StringAttribute{
										Required:    true,
										Description: `The deployment for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource`,
									},
									"openai_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `The API key for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource`,
									},
								},
								Description: `Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("cohere"),
										path.MatchRelative().AtParent().AtName("fake"),
										path.MatchRelative().AtParent().AtName("open_ai"),
										path.MatchRelative().AtParent().AtName("open_ai_compatible"),
									}...),
								},
							},
							"cohere": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"cohere_key": schema.StringAttribute{
										Required:  true,
										Sensitive: true,
									},
								},
								Description: `Use the Cohere API to embed text.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("azure_open_ai"),
										path.MatchRelative().AtParent().AtName("fake"),
										path.MatchRelative().AtParent().AtName("open_ai"),
										path.MatchRelative().AtParent().AtName("open_ai_compatible"),
									}...),
								},
							},
							"fake": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("azure_open_ai"),
										path.MatchRelative().AtParent().AtName("cohere"),
										path.MatchRelative().AtParent().AtName("open_ai"),
										path.MatchRelative().AtParent().AtName("open_ai_compatible"),
									}...),
								},
							},
							"open_ai": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"openai_key": schema.StringAttribute{
										Required:  true,
										Sensitive: true,
									},
								},
								Description: `Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("azure_open_ai"),
										path.MatchRelative().AtParent().AtName("cohere"),
										path.MatchRelative().AtParent().AtName("fake"),
										path.MatchRelative().AtParent().AtName("open_ai_compatible"),
									}...),
								},
							},
							"open_ai_compatible": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"api_key": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Sensitive:   true,
										Default:     stringdefault.StaticString(""),
										Description: `Default: ""`,
									},
									"base_url": schema.StringAttribute{
										Required:    true,
										Description: `The base URL for your OpenAI-compatible service`,
									},
									"dimensions": schema.Int64Attribute{
										Required:    true,
										Description: `The number of dimensions the embedding model is generating`,
									},
									"model_name": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Default:     stringdefault.StaticString("text-embedding-ada-002"),
										Description: `The name of the model to use for embedding. Default: "text-embedding-ada-002"`,
									},
								},
								Description: `Use a service that's compatible with the OpenAI API to embed text.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("azure_open_ai"),
										path.MatchRelative().AtParent().AtName("cohere"),
										path.MatchRelative().AtParent().AtName("fake"),
										path.MatchRelative().AtParent().AtName("open_ai"),
									}...),
								},
							},
						},
						Description: `Embedding configuration`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"indexing": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"credentials": schema.SingleNestedAttribute{
								Required: true,
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `Enter the password you want to use to access the database`,
									},
								},
							},
							"database": schema.StringAttribute{
								Required:    true,
								Description: `Enter the name of the database that you want to sync data into`,
							},
							"default_schema": schema.StringAttribute{
								Required:    true,
								Description: `Enter the name of the default schema`,
							},
							"host": schema.StringAttribute{
								Required:    true,
								Description: `Enter the account name you want to use to access the database. This is usually the identifier before .snowflakecomputing.com`,
							},
							"role": schema.StringAttribute{
								Required:    true,
								Description: `Enter the role that you want to use to access Snowflake`,
							},
							"username": schema.StringAttribute{
								Required:    true,
								Description: `Enter the name of the user you want to use to access the database`,
							},
							"warehouse": schema.StringAttribute{
								Required:    true,
								Description: `Enter the name of the warehouse that you want to sync data into`,
							},
						},
						Description: `Snowflake can be used to store vector data and retrieve embeddings.`,
					},
					"omit_raw_text": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
						Description: `Do not store the text that gets embedded along with the vector and the metadata in the destination. If set to true, only the vector and the metadata will be stored - in this case raw text for LLM use cases needs to be retrieved from another source. Default: false`,
					},
					"processing": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"chunk_overlap": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(0),
								Description: `Size of overlap between chunks in tokens to store in vector store to better capture relevant context. Default: 0`,
							},
							"chunk_size": schema.Int64Attribute{
								Required:    true,
								Description: `Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)`,
							},
							"field_name_mappings": schema.ListNestedAttribute{
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from_field": schema.StringAttribute{
											Required:    true,
											Description: `The field name in the source`,
										},
										"to_field": schema.StringAttribute{
											Required:    true,
											Description: `The field name to use in the destination`,
										},
									},
								},
								Description: `List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.`,
							},
							"metadata_fields": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.`,
							},
							"text_fields": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array.`,
							},
							"text_splitter": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"by_markdown_header": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"split_level": schema.Int64Attribute{
												Computed:    true,
												Optional:    true,
												Default:     int64default.StaticInt64(1),
												Description: `Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points. Default: 1`,
											},
										},
										Description: `Split the text by Markdown headers down to the specified header level. If the chunk size fits multiple sections, they will be combined into a single chunk.`,
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(path.Expressions{
												path.MatchRelative().AtParent().AtName("by_programming_language"),
												path.MatchRelative().AtParent().AtName("by_separator"),
											}...),
										},
									},
									"by_programming_language": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"language": schema.StringAttribute{
												Required:    true,
												Description: `Split code in suitable places based on the programming language. must be one of ["cpp", "go", "java", "js", "php", "proto", "python", "rst", "ruby", "rust", "scala", "swift", "markdown", "latex", "html", "sol"]`,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"cpp",
														"go",
														"java",
														"js",
														"php",
														"proto",
														"python",
														"rst",
														"ruby",
														"rust",
														"scala",
														"swift",
														"markdown",
														"latex",
														"html",
														"sol",
													),
												},
											},
										},
										Description: `Split the text by suitable delimiters based on the programming language. This is useful for splitting code into chunks.`,
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(path.Expressions{
												path.MatchRelative().AtParent().AtName("by_markdown_header"),
												path.MatchRelative().AtParent().AtName("by_separator"),
											}...),
										},
									},
									"by_separator": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"keep_separator": schema.BoolAttribute{
												Computed:    true,
												Optional:    true,
												Default:     booldefault.StaticBool(false),
												Description: `Whether to keep the separator in the resulting chunks. Default: false`,
											},
											"separators": schema.ListAttribute{
												Optional:    true,
												ElementType: types.StringType,
												Description: `List of separator strings to split text fields by. The separator itself needs to be wrapped in double quotes, e.g. to split by the dot character, use ".". To split by a newline, use "\n".`,
											},
										},
										Description: `Split the text by the list of separators until the chunk size is reached, using the earlier mentioned separators where possible. This is useful for splitting text fields by paragraphs, sentences, words, etc.`,
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(path.Expressions{
												path.MatchRelative().AtParent().AtName("by_markdown_header"),
												path.MatchRelative().AtParent().AtName("by_programming_language"),
											}...),
										},
									},
								},
								Description: `Split text fields into chunks based on the specified method.`,
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
							},
						},
					},
				},
				MarkdownDescription: `The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,` + "\n" +
					`as well as to provide type safety for the configuration passed to the destination.` + "\n" +
					`` + "\n" +
					`The configuration model is composed of four parts:` + "\n" +
					`* Processing configuration` + "\n" +
					`* Embedding configuration` + "\n" +
					`* Indexing configuration` + "\n" +
					`* Advanced configuration` + "\n" +
					`` + "\n" +
					`Processing, embedding and advanced configuration are provided by this base class, while the indexing configuration is provided by the destination connector in the sub class.`,
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed. `,
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Name of the destination e.g. dev-mysql-instance.`,
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

func (r *DestinationSnowflakeCortexResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationSnowflakeCortexResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationSnowflakeCortexResourceModel
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

	request := data.ToSharedDestinationSnowflakeCortexCreateRequest()
	res, err := r.client.Destinations.CreateDestinationSnowflakeCortex(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	destinationID := data.DestinationID.ValueString()
	request1 := operations.GetDestinationSnowflakeCortexRequest{
		DestinationID: destinationID,
	}
	res1, err := r.client.Destinations.GetDestinationSnowflakeCortex(ctx, request1)
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
	if res1.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationSnowflakeCortexResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationSnowflakeCortexResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationSnowflakeCortexRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationSnowflakeCortex(ctx, request)
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

func (r *DestinationSnowflakeCortexResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationSnowflakeCortexResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationSnowflakeCortexPutRequest := data.ToSharedDestinationSnowflakeCortexPutRequest()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationSnowflakeCortexRequest{
		DestinationSnowflakeCortexPutRequest: destinationSnowflakeCortexPutRequest,
		DestinationID:                        destinationID,
	}
	res, err := r.client.Destinations.PutDestinationSnowflakeCortex(ctx, request)
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
	destinationId1 := data.DestinationID.ValueString()
	request1 := operations.GetDestinationSnowflakeCortexRequest{
		DestinationID: destinationId1,
	}
	res1, err := r.client.Destinations.GetDestinationSnowflakeCortex(ctx, request1)
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
	if res1.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationSnowflakeCortexResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationSnowflakeCortexResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationSnowflakeCortexRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationSnowflakeCortex(ctx, request)
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

func (r *DestinationSnowflakeCortexResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
