// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceGoogleSearchConsoleAuthenticationType struct {
	OAuth                           *AuthenticateViaGoogleOauth                               `tfsdk:"o_auth" tfPlanOnly:"true"`
	ServiceAccountKeyAuthentication *SourceGoogleSearchConsoleServiceAccountKeyAuthentication `tfsdk:"service_account_key_authentication" tfPlanOnly:"true"`
}
