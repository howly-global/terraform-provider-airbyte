// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceGoogleDriveAuthentication struct {
	AuthenticateViaGoogleOAuth      *DestinationGoogleSheetsAuthenticationViaGoogleOAuth `tfsdk:"authenticate_via_google_o_auth" tfPlanOnly:"true"`
	ServiceAccountKeyAuthentication *SourceGoogleDriveServiceAccountKeyAuthentication    `tfsdk:"service_account_key_authentication" tfPlanOnly:"true"`
}
