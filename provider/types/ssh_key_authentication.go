// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SSHKeyAuthentication struct {
	SSHKey     types.String `tfsdk:"ssh_key"`
	TunnelHost types.String `tfsdk:"tunnel_host"`
	TunnelPort types.Int64  `tfsdk:"tunnel_port"`
	TunnelUser types.String `tfsdk:"tunnel_user"`
}
