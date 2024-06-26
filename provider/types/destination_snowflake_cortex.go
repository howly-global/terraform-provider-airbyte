// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationSnowflakeCortex struct {
	Embedding   DestinationAstraEmbedding             `tfsdk:"embedding"`
	Indexing    DestinationSnowflakeCortexIndexing    `tfsdk:"indexing"`
	OmitRawText types.Bool                            `tfsdk:"omit_raw_text"`
	Processing  DestinationAstraProcessingConfigModel `tfsdk:"processing"`
}
