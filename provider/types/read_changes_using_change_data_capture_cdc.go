// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ReadChangesUsingChangeDataCaptureCDC struct {
	InitialWaitingSeconds            types.Int64  `tfsdk:"initial_waiting_seconds"`
	InvalidCdcCursorPositionBehavior types.String `tfsdk:"invalid_cdc_cursor_position_behavior"`
	QueueSize                        types.Int64  `tfsdk:"queue_size"`
}