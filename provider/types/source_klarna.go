// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceKlarna struct {
	Password   types.String `tfsdk:"password"`
	Playground types.Bool   `tfsdk:"playground"`
	Region     types.String `tfsdk:"region"`
	Username   types.String `tfsdk:"username"`
}
