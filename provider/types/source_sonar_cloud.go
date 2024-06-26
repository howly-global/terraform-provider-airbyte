// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSonarCloud struct {
	ComponentKeys []types.String `tfsdk:"component_keys"`
	EndDate       types.String   `tfsdk:"end_date"`
	Organization  types.String   `tfsdk:"organization"`
	StartDate     types.String   `tfsdk:"start_date"`
	UserToken     types.String   `tfsdk:"user_token"`
}
