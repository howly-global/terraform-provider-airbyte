// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSentry struct {
	AuthToken      types.String   `tfsdk:"auth_token"`
	DiscoverFields []types.String `tfsdk:"discover_fields"`
	Hostname       types.String   `tfsdk:"hostname"`
	Organization   types.String   `tfsdk:"organization"`
	Project        types.String   `tfsdk:"project"`
}
