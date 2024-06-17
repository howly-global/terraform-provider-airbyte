// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSftpBulkResourceModel) ToSharedSourceSftpBulkCreateRequest() *shared.SourceSftpBulkCreateRequest {
	var credentials shared.SourceSftpBulkAuthentication
	var sourceSftpBulkAuthenticateViaPassword *shared.SourceSftpBulkAuthenticateViaPassword
	if r.Configuration.Credentials.AuthenticateViaPassword != nil {
		password := r.Configuration.Credentials.AuthenticateViaPassword.Password.ValueString()
		sourceSftpBulkAuthenticateViaPassword = &shared.SourceSftpBulkAuthenticateViaPassword{
			Password: password,
		}
	}
	if sourceSftpBulkAuthenticateViaPassword != nil {
		credentials = shared.SourceSftpBulkAuthentication{
			SourceSftpBulkAuthenticateViaPassword: sourceSftpBulkAuthenticateViaPassword,
		}
	}
	var sourceSftpBulkAuthenticateViaPrivateKey *shared.SourceSftpBulkAuthenticateViaPrivateKey
	if r.Configuration.Credentials.AuthenticateViaPrivateKey != nil {
		privateKey := r.Configuration.Credentials.AuthenticateViaPrivateKey.PrivateKey.ValueString()
		sourceSftpBulkAuthenticateViaPrivateKey = &shared.SourceSftpBulkAuthenticateViaPrivateKey{
			PrivateKey: privateKey,
		}
	}
	if sourceSftpBulkAuthenticateViaPrivateKey != nil {
		credentials = shared.SourceSftpBulkAuthentication{
			SourceSftpBulkAuthenticateViaPrivateKey: sourceSftpBulkAuthenticateViaPrivateKey,
		}
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	host := r.Configuration.Host.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceSftpBulkFileBasedStreamConfig = []shared.SourceSftpBulkFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceSftpBulkFormat
		var sourceSftpBulkAvroFormat *shared.SourceSftpBulkAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceSftpBulkAvroFormat = &shared.SourceSftpBulkAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceSftpBulkAvroFormat != nil {
			format = shared.SourceSftpBulkFormat{
				SourceSftpBulkAvroFormat: sourceSftpBulkAvroFormat,
			}
		}
		var sourceSftpBulkCSVFormat *shared.SourceSftpBulkCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = []string{}
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceSftpBulkCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceSftpBulkFromCSV *shared.SourceSftpBulkFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceSftpBulkFromCSV = &shared.SourceSftpBulkFromCSV{}
				}
				if sourceSftpBulkFromCSV != nil {
					headerDefinition = &shared.SourceSftpBulkCSVHeaderDefinition{
						SourceSftpBulkFromCSV: sourceSftpBulkFromCSV,
					}
				}
				var sourceSftpBulkAutogenerated *shared.SourceSftpBulkAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceSftpBulkAutogenerated = &shared.SourceSftpBulkAutogenerated{}
				}
				if sourceSftpBulkAutogenerated != nil {
					headerDefinition = &shared.SourceSftpBulkCSVHeaderDefinition{
						SourceSftpBulkAutogenerated: sourceSftpBulkAutogenerated,
					}
				}
				var sourceSftpBulkUserProvided *shared.SourceSftpBulkUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceSftpBulkUserProvided = &shared.SourceSftpBulkUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceSftpBulkUserProvided != nil {
					headerDefinition = &shared.SourceSftpBulkCSVHeaderDefinition{
						SourceSftpBulkUserProvided: sourceSftpBulkUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
			}
			inferenceType := new(shared.SourceSftpBulkInferenceType)
			if !streamsItem.Format.CSVFormat.InferenceType.IsUnknown() && !streamsItem.Format.CSVFormat.InferenceType.IsNull() {
				*inferenceType = shared.SourceSftpBulkInferenceType(streamsItem.Format.CSVFormat.InferenceType.ValueString())
			} else {
				inferenceType = nil
			}
			var nullValues []string = []string{}
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = []string{}
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceSftpBulkCSVFormat = &shared.SourceSftpBulkCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				InferenceType:                inferenceType,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceSftpBulkCSVFormat != nil {
			format = shared.SourceSftpBulkFormat{
				SourceSftpBulkCSVFormat: sourceSftpBulkCSVFormat,
			}
		}
		var sourceSftpBulkJsonlFormat *shared.SourceSftpBulkJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceSftpBulkJsonlFormat = &shared.SourceSftpBulkJsonlFormat{}
		}
		if sourceSftpBulkJsonlFormat != nil {
			format = shared.SourceSftpBulkFormat{
				SourceSftpBulkJsonlFormat: sourceSftpBulkJsonlFormat,
			}
		}
		var sourceSftpBulkParquetFormat *shared.SourceSftpBulkParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceSftpBulkParquetFormat = &shared.SourceSftpBulkParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceSftpBulkParquetFormat != nil {
			format = shared.SourceSftpBulkFormat{
				SourceSftpBulkParquetFormat: sourceSftpBulkParquetFormat,
			}
		}
		var sourceSftpBulkDocumentFileTypeFormatExperimental *shared.SourceSftpBulkDocumentFileTypeFormatExperimental
		if streamsItem.Format.DocumentFileTypeFormatExperimental != nil {
			var processing *shared.SourceSftpBulkProcessing
			if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing != nil {
				var sourceSftpBulkLocal *shared.SourceSftpBulkLocal
				if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.Local != nil {
					sourceSftpBulkLocal = &shared.SourceSftpBulkLocal{}
				}
				if sourceSftpBulkLocal != nil {
					processing = &shared.SourceSftpBulkProcessing{
						SourceSftpBulkLocal: sourceSftpBulkLocal,
					}
				}
				var sourceSftpBulkViaAPI *shared.SourceSftpBulkViaAPI
				if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI != nil {
					apiKey := new(string)
					if !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIKey.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIKey.IsNull() {
						*apiKey = streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIKey.ValueString()
					} else {
						apiKey = nil
					}
					apiURL := new(string)
					if !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIURL.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIURL.IsNull() {
						*apiURL = streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIURL.ValueString()
					} else {
						apiURL = nil
					}
					var parameters []shared.SourceSftpBulkAPIParameterConfigModel = []shared.SourceSftpBulkAPIParameterConfigModel{}
					for _, parametersItem := range streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.Parameters {
						name := parametersItem.Name.ValueString()
						value := parametersItem.Value.ValueString()
						parameters = append(parameters, shared.SourceSftpBulkAPIParameterConfigModel{
							Name:  name,
							Value: value,
						})
					}
					sourceSftpBulkViaAPI = &shared.SourceSftpBulkViaAPI{
						APIKey:     apiKey,
						APIURL:     apiURL,
						Parameters: parameters,
					}
				}
				if sourceSftpBulkViaAPI != nil {
					processing = &shared.SourceSftpBulkProcessing{
						SourceSftpBulkViaAPI: sourceSftpBulkViaAPI,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceSftpBulkParsingStrategy)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsNull() {
				*strategy = shared.SourceSftpBulkParsingStrategy(streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceSftpBulkDocumentFileTypeFormatExperimental = &shared.SourceSftpBulkDocumentFileTypeFormatExperimental{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceSftpBulkDocumentFileTypeFormatExperimental != nil {
			format = shared.SourceSftpBulkFormat{
				SourceSftpBulkDocumentFileTypeFormatExperimental: sourceSftpBulkDocumentFileTypeFormatExperimental,
			}
		}
		var globs []string = []string{}
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		legacyPrefix := new(string)
		if !streamsItem.LegacyPrefix.IsUnknown() && !streamsItem.LegacyPrefix.IsNull() {
			*legacyPrefix = streamsItem.LegacyPrefix.ValueString()
		} else {
			legacyPrefix = nil
		}
		name1 := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceSftpBulkValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceSftpBulkValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceSftpBulkFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			LegacyPrefix:              legacyPrefix,
			Name:                      name1,
			PrimaryKey:                primaryKey,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceSftpBulk{
		Credentials: credentials,
		FolderPath:  folderPath,
		Host:        host,
		Port:        port,
		StartDate:   startDate,
		Streams:     streams,
		Username:    username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name2 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSftpBulkCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name2,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSftpBulkResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceSftpBulkResourceModel) ToSharedSourceSftpBulkPutRequest() *shared.SourceSftpBulkPutRequest {
	var credentials shared.SourceSftpBulkUpdateAuthentication
	var authenticateViaPassword *shared.AuthenticateViaPassword
	if r.Configuration.Credentials.AuthenticateViaPassword != nil {
		password := r.Configuration.Credentials.AuthenticateViaPassword.Password.ValueString()
		authenticateViaPassword = &shared.AuthenticateViaPassword{
			Password: password,
		}
	}
	if authenticateViaPassword != nil {
		credentials = shared.SourceSftpBulkUpdateAuthentication{
			AuthenticateViaPassword: authenticateViaPassword,
		}
	}
	var authenticateViaPrivateKey *shared.AuthenticateViaPrivateKey
	if r.Configuration.Credentials.AuthenticateViaPrivateKey != nil {
		privateKey := r.Configuration.Credentials.AuthenticateViaPrivateKey.PrivateKey.ValueString()
		authenticateViaPrivateKey = &shared.AuthenticateViaPrivateKey{
			PrivateKey: privateKey,
		}
	}
	if authenticateViaPrivateKey != nil {
		credentials = shared.SourceSftpBulkUpdateAuthentication{
			AuthenticateViaPrivateKey: authenticateViaPrivateKey,
		}
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	host := r.Configuration.Host.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceSftpBulkUpdateFileBasedStreamConfig = []shared.SourceSftpBulkUpdateFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceSftpBulkUpdateFormat
		var sourceSftpBulkUpdateAvroFormat *shared.SourceSftpBulkUpdateAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceSftpBulkUpdateAvroFormat = &shared.SourceSftpBulkUpdateAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceSftpBulkUpdateAvroFormat != nil {
			format = shared.SourceSftpBulkUpdateFormat{
				SourceSftpBulkUpdateAvroFormat: sourceSftpBulkUpdateAvroFormat,
			}
		}
		var sourceSftpBulkUpdateCSVFormat *shared.SourceSftpBulkUpdateCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = []string{}
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceSftpBulkUpdateCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceSftpBulkUpdateFromCSV *shared.SourceSftpBulkUpdateFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceSftpBulkUpdateFromCSV = &shared.SourceSftpBulkUpdateFromCSV{}
				}
				if sourceSftpBulkUpdateFromCSV != nil {
					headerDefinition = &shared.SourceSftpBulkUpdateCSVHeaderDefinition{
						SourceSftpBulkUpdateFromCSV: sourceSftpBulkUpdateFromCSV,
					}
				}
				var sourceSftpBulkUpdateAutogenerated *shared.SourceSftpBulkUpdateAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceSftpBulkUpdateAutogenerated = &shared.SourceSftpBulkUpdateAutogenerated{}
				}
				if sourceSftpBulkUpdateAutogenerated != nil {
					headerDefinition = &shared.SourceSftpBulkUpdateCSVHeaderDefinition{
						SourceSftpBulkUpdateAutogenerated: sourceSftpBulkUpdateAutogenerated,
					}
				}
				var sourceSftpBulkUpdateUserProvided *shared.SourceSftpBulkUpdateUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceSftpBulkUpdateUserProvided = &shared.SourceSftpBulkUpdateUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceSftpBulkUpdateUserProvided != nil {
					headerDefinition = &shared.SourceSftpBulkUpdateCSVHeaderDefinition{
						SourceSftpBulkUpdateUserProvided: sourceSftpBulkUpdateUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
			}
			inferenceType := new(shared.SourceSftpBulkUpdateInferenceType)
			if !streamsItem.Format.CSVFormat.InferenceType.IsUnknown() && !streamsItem.Format.CSVFormat.InferenceType.IsNull() {
				*inferenceType = shared.SourceSftpBulkUpdateInferenceType(streamsItem.Format.CSVFormat.InferenceType.ValueString())
			} else {
				inferenceType = nil
			}
			var nullValues []string = []string{}
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = []string{}
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceSftpBulkUpdateCSVFormat = &shared.SourceSftpBulkUpdateCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				InferenceType:                inferenceType,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceSftpBulkUpdateCSVFormat != nil {
			format = shared.SourceSftpBulkUpdateFormat{
				SourceSftpBulkUpdateCSVFormat: sourceSftpBulkUpdateCSVFormat,
			}
		}
		var sourceSftpBulkUpdateJsonlFormat *shared.SourceSftpBulkUpdateJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceSftpBulkUpdateJsonlFormat = &shared.SourceSftpBulkUpdateJsonlFormat{}
		}
		if sourceSftpBulkUpdateJsonlFormat != nil {
			format = shared.SourceSftpBulkUpdateFormat{
				SourceSftpBulkUpdateJsonlFormat: sourceSftpBulkUpdateJsonlFormat,
			}
		}
		var sourceSftpBulkUpdateParquetFormat *shared.SourceSftpBulkUpdateParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceSftpBulkUpdateParquetFormat = &shared.SourceSftpBulkUpdateParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceSftpBulkUpdateParquetFormat != nil {
			format = shared.SourceSftpBulkUpdateFormat{
				SourceSftpBulkUpdateParquetFormat: sourceSftpBulkUpdateParquetFormat,
			}
		}
		var sourceSftpBulkUpdateDocumentFileTypeFormatExperimental *shared.SourceSftpBulkUpdateDocumentFileTypeFormatExperimental
		if streamsItem.Format.DocumentFileTypeFormatExperimental != nil {
			var processing *shared.SourceSftpBulkUpdateProcessing
			if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing != nil {
				var sourceSftpBulkUpdateLocal *shared.SourceSftpBulkUpdateLocal
				if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.Local != nil {
					sourceSftpBulkUpdateLocal = &shared.SourceSftpBulkUpdateLocal{}
				}
				if sourceSftpBulkUpdateLocal != nil {
					processing = &shared.SourceSftpBulkUpdateProcessing{
						SourceSftpBulkUpdateLocal: sourceSftpBulkUpdateLocal,
					}
				}
				var viaAPI *shared.ViaAPI
				if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI != nil {
					apiKey := new(string)
					if !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIKey.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIKey.IsNull() {
						*apiKey = streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIKey.ValueString()
					} else {
						apiKey = nil
					}
					apiURL := new(string)
					if !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIURL.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIURL.IsNull() {
						*apiURL = streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.APIURL.ValueString()
					} else {
						apiURL = nil
					}
					var parameters []shared.APIParameterConfigModel = []shared.APIParameterConfigModel{}
					for _, parametersItem := range streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.ViaAPI.Parameters {
						name := parametersItem.Name.ValueString()
						value := parametersItem.Value.ValueString()
						parameters = append(parameters, shared.APIParameterConfigModel{
							Name:  name,
							Value: value,
						})
					}
					viaAPI = &shared.ViaAPI{
						APIKey:     apiKey,
						APIURL:     apiURL,
						Parameters: parameters,
					}
				}
				if viaAPI != nil {
					processing = &shared.SourceSftpBulkUpdateProcessing{
						ViaAPI: viaAPI,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceSftpBulkUpdateParsingStrategy)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsNull() {
				*strategy = shared.SourceSftpBulkUpdateParsingStrategy(streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceSftpBulkUpdateDocumentFileTypeFormatExperimental = &shared.SourceSftpBulkUpdateDocumentFileTypeFormatExperimental{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceSftpBulkUpdateDocumentFileTypeFormatExperimental != nil {
			format = shared.SourceSftpBulkUpdateFormat{
				SourceSftpBulkUpdateDocumentFileTypeFormatExperimental: sourceSftpBulkUpdateDocumentFileTypeFormatExperimental,
			}
		}
		var globs []string = []string{}
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		legacyPrefix := new(string)
		if !streamsItem.LegacyPrefix.IsUnknown() && !streamsItem.LegacyPrefix.IsNull() {
			*legacyPrefix = streamsItem.LegacyPrefix.ValueString()
		} else {
			legacyPrefix = nil
		}
		name1 := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceSftpBulkUpdateValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceSftpBulkUpdateValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceSftpBulkUpdateFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			LegacyPrefix:              legacyPrefix,
			Name:                      name1,
			PrimaryKey:                primaryKey,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceSftpBulkUpdate{
		Credentials: credentials,
		FolderPath:  folderPath,
		Host:        host,
		Port:        port,
		StartDate:   startDate,
		Streams:     streams,
		Username:    username,
	}
	name2 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSftpBulkPutRequest{
		Configuration: configuration,
		Name:          name2,
		WorkspaceID:   workspaceID,
	}
	return &out
}