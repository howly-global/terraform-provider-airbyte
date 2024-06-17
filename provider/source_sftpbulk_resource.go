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
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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
var _ resource.Resource = &SourceSftpBulkResource{}
var _ resource.ResourceWithImportState = &SourceSftpBulkResource{}

func NewSourceSftpBulkResource() resource.Resource {
	return &SourceSftpBulkResource{}
}

// SourceSftpBulkResource defines the resource implementation.
type SourceSftpBulkResource struct {
	client *sdk.SDK
}

// SourceSftpBulkResourceModel describes the resource data model.
type SourceSftpBulkResourceModel struct {
	Configuration tfTypes.SourceSftpBulk `tfsdk:"configuration"`
	DefinitionID  types.String           `tfsdk:"definition_id"`
	Name          types.String           `tfsdk:"name"`
	SecretID      types.String           `tfsdk:"secret_id"`
	SourceID      types.String           `tfsdk:"source_id"`
	SourceType    types.String           `tfsdk:"source_type"`
	WorkspaceID   types.String           `tfsdk:"workspace_id"`
}

func (r *SourceSftpBulkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_sftp_bulk"
}

func (r *SourceSftpBulkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceSftpBulk Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"authenticate_via_password": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `Password`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("authenticate_via_private_key"),
									}...),
								},
							},
							"authenticate_via_private_key": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"private_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `The Private key`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("authenticate_via_password"),
									}...),
								},
							},
						},
						Description: `Credentials for connecting to the SFTP Server`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"folder_path": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("/"),
						Description: `The directory to search files for sync. Default: "/"`,
					},
					"host": schema.StringAttribute{
						Required:    true,
						Description: `The server host address`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(22),
						Description: `The server port. Default: 22`,
					},
					"start_date": schema.StringAttribute{
						Optional:    true,
						Description: `UTC date and time in the format 2017-01-25T00:00:00.000000Z. Any file modified before this date will not be replicated.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"streams": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"days_to_sync_if_history_is_full": schema.Int64Attribute{
									Computed:    true,
									Optional:    true,
									Default:     int64default.StaticInt64(3),
									Description: `When the state history of the file store is full, syncs will only read files that were last modified in the provided day range. Default: 3`,
								},
								"format": schema.SingleNestedAttribute{
									Required: true,
									Attributes: map[string]schema.Attribute{
										"avro_format": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"double_as_string": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(false),
													Description: `Whether to convert double fields to strings. This is recommended if you have decimal numbers with a high degree of precision because there can be a loss precision when handling floating point numbers. Default: false`,
												},
											},
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("csv_format"),
													path.MatchRelative().AtParent().AtName("document_file_type_format_experimental"),
													path.MatchRelative().AtParent().AtName("jsonl_format"),
													path.MatchRelative().AtParent().AtName("parquet_format"),
												}...),
											},
										},
										"csv_format": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"delimiter": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Default:     stringdefault.StaticString(","),
													Description: `The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'. Default: ","`,
												},
												"double_quote": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(true),
													Description: `Whether two quotes in a quoted CSV value denote a single quote in the data. Default: true`,
												},
												"encoding": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Default:     stringdefault.StaticString("utf8"),
													Description: `The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options. Default: "utf8"`,
												},
												"escape_char": schema.StringAttribute{
													Optional:    true,
													Description: `The character used for escaping special characters. To disallow escaping, leave this field blank.`,
												},
												"false_values": schema.ListAttribute{
													Optional:    true,
													ElementType: types.StringType,
													Description: `A set of case-sensitive strings that should be interpreted as false values.`,
													Validators: []validator.List{
														listvalidator.UniqueValues(),
													},
												},
												"header_definition": schema.SingleNestedAttribute{
													Optional: true,
													Attributes: map[string]schema.Attribute{
														"autogenerated": schema.SingleNestedAttribute{
															Optional:   true,
															Attributes: map[string]schema.Attribute{},
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.Expressions{
																	path.MatchRelative().AtParent().AtName("from_csv"),
																	path.MatchRelative().AtParent().AtName("user_provided"),
																}...),
															},
														},
														"from_csv": schema.SingleNestedAttribute{
															Optional:   true,
															Attributes: map[string]schema.Attribute{},
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.Expressions{
																	path.MatchRelative().AtParent().AtName("autogenerated"),
																	path.MatchRelative().AtParent().AtName("user_provided"),
																}...),
															},
														},
														"user_provided": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"column_names": schema.ListAttribute{
																	Required:    true,
																	ElementType: types.StringType,
																	Description: `The column names that will be used while emitting the CSV records`,
																},
															},
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.Expressions{
																	path.MatchRelative().AtParent().AtName("autogenerated"),
																	path.MatchRelative().AtParent().AtName("from_csv"),
																}...),
															},
														},
													},
													Description: `How headers will be defined. ` + "`" + `User Provided` + "`" + ` assumes the CSV does not have a header row and uses the headers provided and ` + "`" + `Autogenerated` + "`" + ` assumes the CSV does not have a header row and the CDK will generate headers using for ` + "`" + `f{i}` + "`" + ` where ` + "`" + `i` + "`" + ` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.`,
													Validators: []validator.Object{
														validators.ExactlyOneChild(),
													},
												},
												"ignore_errors_on_fields_mismatch": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(false),
													Description: `Whether to ignore errors that occur when the number of fields in the CSV does not match the number of columns in the schema. Default: false`,
												},
												"inference_type": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Default:     stringdefault.StaticString("None"),
													Description: `How to infer the types of the columns. If none, inference default to strings. must be one of ["None", "Primitive Types Only"]; Default: "None"`,
													Validators: []validator.String{
														stringvalidator.OneOf(
															"None",
															"Primitive Types Only",
														),
													},
												},
												"null_values": schema.ListAttribute{
													Optional:    true,
													ElementType: types.StringType,
													Description: `A set of case-sensitive strings that should be interpreted as null values. For example, if the value 'NA' should be interpreted as null, enter 'NA' in this field.`,
													Validators: []validator.List{
														listvalidator.UniqueValues(),
													},
												},
												"quote_char": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Default:     stringdefault.StaticString("\""),
													Description: `The character used for quoting CSV values. To disallow quoting, make this field blank. Default: "\""`,
												},
												"skip_rows_after_header": schema.Int64Attribute{
													Computed:    true,
													Optional:    true,
													Default:     int64default.StaticInt64(0),
													Description: `The number of rows to skip after the header row. Default: 0`,
												},
												"skip_rows_before_header": schema.Int64Attribute{
													Computed:    true,
													Optional:    true,
													Default:     int64default.StaticInt64(0),
													Description: `The number of rows to skip before the header row. For example, if the header row is on the 3rd row, enter 2 in this field. Default: 0`,
												},
												"strings_can_be_null": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(true),
													Description: `Whether strings can be interpreted as null values. If true, strings that match the null_values set will be interpreted as null. If false, strings that match the null_values set will be interpreted as the string itself. Default: true`,
												},
												"true_values": schema.ListAttribute{
													Optional:    true,
													ElementType: types.StringType,
													Description: `A set of case-sensitive strings that should be interpreted as true values.`,
													Validators: []validator.List{
														listvalidator.UniqueValues(),
													},
												},
											},
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("avro_format"),
													path.MatchRelative().AtParent().AtName("document_file_type_format_experimental"),
													path.MatchRelative().AtParent().AtName("jsonl_format"),
													path.MatchRelative().AtParent().AtName("parquet_format"),
												}...),
											},
										},
										"document_file_type_format_experimental": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"processing": schema.SingleNestedAttribute{
													Optional: true,
													Attributes: map[string]schema.Attribute{
														"local": schema.SingleNestedAttribute{
															Optional:    true,
															Attributes:  map[string]schema.Attribute{},
															Description: `Process files locally, supporting ` + "`" + `fast` + "`" + ` and ` + "`" + `ocr` + "`" + ` modes. This is the default option.`,
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.Expressions{
																	path.MatchRelative().AtParent().AtName("via_api"),
																}...),
															},
														},
														"via_api": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"api_key": schema.StringAttribute{
																	Computed:    true,
																	Optional:    true,
																	Sensitive:   true,
																	Default:     stringdefault.StaticString(""),
																	Description: `The API key to use matching the environment. Default: ""`,
																},
																"api_url": schema.StringAttribute{
																	Computed:    true,
																	Optional:    true,
																	Default:     stringdefault.StaticString("https://api.unstructured.io"),
																	Description: `The URL of the unstructured API to use. Default: "https://api.unstructured.io"`,
																},
																"parameters": schema.ListNestedAttribute{
																	Optional: true,
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Required:    true,
																				Description: `The name of the unstructured API parameter to use`,
																			},
																			"value": schema.StringAttribute{
																				Required:    true,
																				Description: `The value of the parameter`,
																			},
																		},
																	},
																	Description: `List of parameters send to the API`,
																},
															},
															Description: `Process files via an API, using the ` + "`" + `hi_res` + "`" + ` mode. This option is useful for increased performance and accuracy, but requires an API key and a hosted instance of unstructured.`,
															Validators: []validator.Object{
																objectvalidator.ConflictsWith(path.Expressions{
																	path.MatchRelative().AtParent().AtName("local"),
																}...),
															},
														},
													},
													Description: `Processing configuration`,
													Validators: []validator.Object{
														validators.ExactlyOneChild(),
													},
												},
												"skip_unprocessable_files": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(true),
													Description: `If true, skip files that cannot be parsed and pass the error message along as the _ab_source_file_parse_error field. If false, fail the sync. Default: true`,
												},
												"strategy": schema.StringAttribute{
													Computed:    true,
													Optional:    true,
													Default:     stringdefault.StaticString("auto"),
													Description: `The strategy used to parse documents. ` + "`" + `fast` + "`" + ` extracts text directly from the document which doesn't work for all files. ` + "`" + `ocr_only` + "`" + ` is more reliable, but slower. ` + "`" + `hi_res` + "`" + ` is the most reliable, but requires an API key and a hosted instance of unstructured and can't be used with local mode. See the unstructured.io documentation for more details: https://unstructured-io.github.io/unstructured/core/partition.html#partition-pdf. must be one of ["auto", "fast", "ocr_only", "hi_res"]; Default: "auto"`,
													Validators: []validator.String{
														stringvalidator.OneOf(
															"auto",
															"fast",
															"ocr_only",
															"hi_res",
														),
													},
												},
											},
											Description: `Extract text from document formats (.pdf, .docx, .md, .pptx) and emit as one record per file.`,
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("avro_format"),
													path.MatchRelative().AtParent().AtName("csv_format"),
													path.MatchRelative().AtParent().AtName("jsonl_format"),
													path.MatchRelative().AtParent().AtName("parquet_format"),
												}...),
											},
										},
										"jsonl_format": schema.SingleNestedAttribute{
											Optional:   true,
											Attributes: map[string]schema.Attribute{},
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("avro_format"),
													path.MatchRelative().AtParent().AtName("csv_format"),
													path.MatchRelative().AtParent().AtName("document_file_type_format_experimental"),
													path.MatchRelative().AtParent().AtName("parquet_format"),
												}...),
											},
										},
										"parquet_format": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"decimal_as_float": schema.BoolAttribute{
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(false),
													Description: `Whether to convert decimal fields to floats. There is a loss of precision when converting decimals to floats, so this is not recommended. Default: false`,
												},
											},
											Validators: []validator.Object{
												objectvalidator.ConflictsWith(path.Expressions{
													path.MatchRelative().AtParent().AtName("avro_format"),
													path.MatchRelative().AtParent().AtName("csv_format"),
													path.MatchRelative().AtParent().AtName("document_file_type_format_experimental"),
													path.MatchRelative().AtParent().AtName("jsonl_format"),
												}...),
											},
										},
									},
									Description: `The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.`,
									Validators: []validator.Object{
										validators.ExactlyOneChild(),
									},
								},
								"globs": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
									Description: `The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look <a href="https://en.wikipedia.org/wiki/Glob_(programming)">here</a>.`,
								},
								"input_schema": schema.StringAttribute{
									Optional:    true,
									Description: `The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.`,
								},
								"legacy_prefix": schema.StringAttribute{
									Optional:    true,
									Description: `The path prefix configured in v3 versions of the S3 connector. This option is deprecated in favor of a single glob.`,
								},
								"name": schema.StringAttribute{
									Required:    true,
									Description: `The name of the stream.`,
								},
								"primary_key": schema.StringAttribute{
									Optional:    true,
									Sensitive:   true,
									Description: `The column or columns (for a composite key) that serves as the unique identifier of a record. If empty, the primary key will default to the parser's default primary key.`,
								},
								"schemaless": schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Default:     booldefault.StaticBool(false),
									Description: `When enabled, syncs will not validate or structure records against the stream's schema. Default: false`,
								},
								"validation_policy": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Default:     stringdefault.StaticString("Emit Record"),
									Description: `The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema. must be one of ["Emit Record", "Skip Record", "Wait for Discover"]; Default: "Emit Record"`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"Emit Record",
											"Skip Record",
											"Wait for Discover",
										),
									},
								},
							},
						},
						Description: `Each instance of this configuration defines a <a href="https://docs.airbyte.com/cloud/core-concepts#stream">stream</a>. Use this to define which files belong in the stream, their format, and how they should be parsed and validated. When sending data to warehouse destination such as Snowflake or BigQuery, each stream is a separate table.`,
					},
					"username": schema.StringAttribute{
						Required:    true,
						Description: `The server user`,
					},
				},
				MarkdownDescription: `Used during spec; allows the developer to configure the cloud provider specific options` + "\n" +
					`that are needed when users configure a file-based source.`,
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

func (r *SourceSftpBulkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceSftpBulkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceSftpBulkResourceModel
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

	request := data.ToSharedSourceSftpBulkCreateRequest()
	res, err := r.client.Sources.CreateSourceSftpBulk(ctx, request)
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
	request1 := operations.GetSourceSftpBulkRequest{
		SourceID: sourceID,
	}
	res1, err := r.client.Sources.GetSourceSftpBulk(ctx, request1)
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

func (r *SourceSftpBulkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceSftpBulkResourceModel
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
	request := operations.GetSourceSftpBulkRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceSftpBulk(ctx, request)
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

func (r *SourceSftpBulkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceSftpBulkResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceSftpBulkPutRequest := data.ToSharedSourceSftpBulkPutRequest()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceSftpBulkRequest{
		SourceSftpBulkPutRequest: sourceSftpBulkPutRequest,
		SourceID:                 sourceID,
	}
	res, err := r.client.Sources.PutSourceSftpBulk(ctx, request)
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
	request1 := operations.GetSourceSftpBulkRequest{
		SourceID: sourceId1,
	}
	res1, err := r.client.Sources.GetSourceSftpBulk(ctx, request1)
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

func (r *SourceSftpBulkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceSftpBulkResourceModel
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
	request := operations.DeleteSourceSftpBulkRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceSftpBulk(ctx, request)
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

func (r *SourceSftpBulkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
