// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PasswordAuthentication struct {
	TunnelHost         types.String `tfsdk:"tunnel_host"`
	TunnelPort         types.Int64  `tfsdk:"tunnel_port"`
	TunnelUser         types.String `tfsdk:"tunnel_user"`
	TunnelUserPassword types.String `tfsdk:"tunnel_user_password"`
}
