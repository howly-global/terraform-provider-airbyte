// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceGithubAuthentication struct {
	OAuth               *OAuth                               `tfsdk:"o_auth" tfPlanOnly:"true"`
	PersonalAccessToken *AuthenticateWithPersonalAccessToken `tfsdk:"personal_access_token" tfPlanOnly:"true"`
}
