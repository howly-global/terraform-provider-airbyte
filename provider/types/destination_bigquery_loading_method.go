// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type DestinationBigqueryLoadingMethod struct {
	GCSStaging      *GCSStaging `tfsdk:"gcs_staging" tfPlanOnly:"true"`
	StandardInserts *Fake       `tfsdk:"standard_inserts" tfPlanOnly:"true"`
}