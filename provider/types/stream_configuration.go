// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type StreamConfiguration struct {
	CursorField []types.String   `tfsdk:"cursor_field"`
	Name        types.String     `tfsdk:"name"`
	PrimaryKey  [][]types.String `tfsdk:"primary_key"`
	SyncMode    types.String     `tfsdk:"sync_mode"`
}
