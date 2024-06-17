// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceKyve struct {
	MaxPages types.Int64  `tfsdk:"max_pages"`
	PageSize types.Int64  `tfsdk:"page_size"`
	PoolIds  types.String `tfsdk:"pool_ids"`
	StartIds types.String `tfsdk:"start_ids"`
	URLBase  types.String `tfsdk:"url_base"`
}
