// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Jsonl struct {
	BlockSize               types.Int64  `tfsdk:"block_size"`
	NewlinesInValues        types.Bool   `tfsdk:"newlines_in_values"`
	UnexpectedFieldBehavior types.String `tfsdk:"unexpected_field_behavior"`
}
