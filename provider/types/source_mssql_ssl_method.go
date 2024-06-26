// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceMssqlSSLMethod struct {
	EncryptedTrustServerCertificate *Fake                                  `tfsdk:"encrypted_trust_server_certificate" tfPlanOnly:"true"`
	EncryptedVerifyCertificate      *SourceMssqlEncryptedVerifyCertificate `tfsdk:"encrypted_verify_certificate" tfPlanOnly:"true"`
	Unencrypted                     *Fake                                  `tfsdk:"unencrypted" tfPlanOnly:"true"`
}
