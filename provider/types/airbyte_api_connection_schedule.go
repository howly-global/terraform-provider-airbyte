// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type AirbyteAPIConnectionSchedule struct {
	BasicTiming    types.String `tfsdk:"basic_timing"`
	CronExpression types.String `tfsdk:"cron_expression"`
	ScheduleType   types.String `tfsdk:"schedule_type"`
}
