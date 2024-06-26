// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type DestinationGcsOutputFormat struct {
	AvroApacheAvro                *AvroApacheAvro                              `tfsdk:"avro_apache_avro" tfPlanOnly:"true"`
	CSVCommaSeparatedValues       *DestinationGcsCSVCommaSeparatedValues       `tfsdk:"csv_comma_separated_values" tfPlanOnly:"true"`
	JSONLinesNewlineDelimitedJSON *DestinationGcsJSONLinesNewlineDelimitedJSON `tfsdk:"json_lines_newline_delimited_json" tfPlanOnly:"true"`
	ParquetColumnarStorage        *DestinationGcsParquetColumnarStorage        `tfsdk:"parquet_columnar_storage" tfPlanOnly:"true"`
}
