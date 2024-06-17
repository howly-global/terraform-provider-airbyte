// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSenseforce struct {
	AccessToken types.String `tfsdk:"access_token"`
	BackendURL  types.String `tfsdk:"backend_url"`
	DatasetID   types.String `tfsdk:"dataset_id"`
	SliceRange  types.Int64  `tfsdk:"slice_range"`
	StartDate   types.String `tfsdk:"start_date"`
}
