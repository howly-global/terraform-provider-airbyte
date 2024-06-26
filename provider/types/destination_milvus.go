// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationMilvus struct {
	Embedding   DestinationAstraEmbedding             `tfsdk:"embedding"`
	Indexing    DestinationMilvusIndexing             `tfsdk:"indexing"`
	OmitRawText types.Bool                            `tfsdk:"omit_raw_text"`
	Processing  DestinationAstraProcessingConfigModel `tfsdk:"processing"`
}
