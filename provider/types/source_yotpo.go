// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceYotpo struct {
	AccessToken types.String `tfsdk:"access_token"`
	AppKey      types.String `tfsdk:"app_key"`
	Email       types.String `tfsdk:"email"`
	StartDate   types.String `tfsdk:"start_date"`
}
