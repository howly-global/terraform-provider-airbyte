// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceMicrosoftTeamsAuthenticationMechanism struct {
	AuthenticateViaMicrosoft        *AuthenticateViaMicrosoft `tfsdk:"authenticate_via_microsoft" tfPlanOnly:"true"`
	AuthenticateViaMicrosoftOAuth20 *AuthenticateViaOauth2    `tfsdk:"authenticate_via_microsoft_o_auth20" tfPlanOnly:"true"`
}
