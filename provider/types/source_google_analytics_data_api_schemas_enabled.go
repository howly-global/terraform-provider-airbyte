// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceGoogleAnalyticsDataAPISchemasEnabled struct {
	CohortReportSettings *SourceGoogleAnalyticsDataAPICohortReportSettings `tfsdk:"cohort_report_settings"`
	Cohorts              []Cohorts                                         `tfsdk:"cohorts"`
	CohortsRange         *SourceGoogleAnalyticsDataAPICohortsRange         `tfsdk:"cohorts_range"`
}
