// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceRetentlyAuthenticationMechanism struct {
	AuthenticateViaRetentlyOAuth *AuthenticateViaHarvestOAuth `tfsdk:"authenticate_via_retently_o_auth" tfPlanOnly:"true"`
	AuthenticateWithAPIToken     *AuthenticateWithAPIToken    `tfsdk:"authenticate_with_api_token" tfPlanOnly:"true"`
}
