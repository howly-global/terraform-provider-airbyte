// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceSlackAuthenticationMechanism struct {
	APIToken            *SourceK6Cloud       `tfsdk:"api_token" tfPlanOnly:"true"`
	SignInViaSlackOAuth *SourceNotionOAuth20 `tfsdk:"sign_in_via_slack_o_auth" tfPlanOnly:"true"`
}
