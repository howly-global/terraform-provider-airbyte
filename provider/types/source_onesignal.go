// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceOnesignal struct {
	Applications []Applications `tfsdk:"applications"`
	OutcomeNames types.String   `tfsdk:"outcome_names"`
	StartDate    types.String   `tfsdk:"start_date"`
	UserAuthKey  types.String   `tfsdk:"user_auth_key"`
}
