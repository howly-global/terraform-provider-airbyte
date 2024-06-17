// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationRedisResourceModel) ToSharedDestinationRedisCreateRequest() *shared.DestinationRedisCreateRequest {
	cacheType := new(shared.DestinationRedisCacheType)
	if !r.Configuration.CacheType.IsUnknown() && !r.Configuration.CacheType.IsNull() {
		*cacheType = shared.DestinationRedisCacheType(r.Configuration.CacheType.ValueString())
	} else {
		cacheType = nil
	}
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	ssl := new(bool)
	if !r.Configuration.Ssl.IsUnknown() && !r.Configuration.Ssl.IsNull() {
		*ssl = r.Configuration.Ssl.ValueBool()
	} else {
		ssl = nil
	}
	var sslMode *shared.DestinationRedisSSLModes
	if r.Configuration.SslMode != nil {
		var destinationRedisDisable *shared.DestinationRedisDisable
		if r.Configuration.SslMode.Disable != nil {
			destinationRedisDisable = &shared.DestinationRedisDisable{}
		}
		if destinationRedisDisable != nil {
			sslMode = &shared.DestinationRedisSSLModes{
				DestinationRedisDisable: destinationRedisDisable,
			}
		}
		var destinationRedisVerifyFull *shared.DestinationRedisVerifyFull
		if r.Configuration.SslMode.VerifyFull != nil {
			caCertificate := r.Configuration.SslMode.VerifyFull.CaCertificate.ValueString()
			clientCertificate := r.Configuration.SslMode.VerifyFull.ClientCertificate.ValueString()
			clientKey := r.Configuration.SslMode.VerifyFull.ClientKey.ValueString()
			clientKeyPassword := new(string)
			if !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsNull() {
				*clientKeyPassword = r.Configuration.SslMode.VerifyFull.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword = nil
			}
			destinationRedisVerifyFull = &shared.DestinationRedisVerifyFull{
				CaCertificate:     caCertificate,
				ClientCertificate: clientCertificate,
				ClientKey:         clientKey,
				ClientKeyPassword: clientKeyPassword,
			}
		}
		if destinationRedisVerifyFull != nil {
			sslMode = &shared.DestinationRedisSSLModes{
				DestinationRedisVerifyFull: destinationRedisVerifyFull,
			}
		}
	}
	var tunnelMethod *shared.DestinationRedisSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationRedisNoTunnel *shared.DestinationRedisNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationRedisNoTunnel = &shared.DestinationRedisNoTunnel{}
		}
		if destinationRedisNoTunnel != nil {
			tunnelMethod = &shared.DestinationRedisSSHTunnelMethod{
				DestinationRedisNoTunnel: destinationRedisNoTunnel,
			}
		}
		var destinationRedisSSHKeyAuthentication *shared.DestinationRedisSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationRedisSSHKeyAuthentication = &shared.DestinationRedisSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationRedisSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationRedisSSHTunnelMethod{
				DestinationRedisSSHKeyAuthentication: destinationRedisSSHKeyAuthentication,
			}
		}
		var destinationRedisPasswordAuthentication *shared.DestinationRedisPasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationRedisPasswordAuthentication = &shared.DestinationRedisPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationRedisPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationRedisSSHTunnelMethod{
				DestinationRedisPasswordAuthentication: destinationRedisPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationRedis{
		CacheType:    cacheType,
		Host:         host,
		Password:     password,
		Port:         port,
		Ssl:          ssl,
		SslMode:      sslMode,
		TunnelMethod: tunnelMethod,
		Username:     username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationRedisCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationRedisResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationRedisResourceModel) ToSharedDestinationRedisPutRequest() *shared.DestinationRedisPutRequest {
	cacheType := new(shared.CacheType)
	if !r.Configuration.CacheType.IsUnknown() && !r.Configuration.CacheType.IsNull() {
		*cacheType = shared.CacheType(r.Configuration.CacheType.ValueString())
	} else {
		cacheType = nil
	}
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	ssl := new(bool)
	if !r.Configuration.Ssl.IsUnknown() && !r.Configuration.Ssl.IsNull() {
		*ssl = r.Configuration.Ssl.ValueBool()
	} else {
		ssl = nil
	}
	var sslMode *shared.DestinationRedisUpdateSSLModes
	if r.Configuration.SslMode != nil {
		var destinationRedisUpdateDisable *shared.DestinationRedisUpdateDisable
		if r.Configuration.SslMode.Disable != nil {
			destinationRedisUpdateDisable = &shared.DestinationRedisUpdateDisable{}
		}
		if destinationRedisUpdateDisable != nil {
			sslMode = &shared.DestinationRedisUpdateSSLModes{
				DestinationRedisUpdateDisable: destinationRedisUpdateDisable,
			}
		}
		var destinationRedisUpdateVerifyFull *shared.DestinationRedisUpdateVerifyFull
		if r.Configuration.SslMode.VerifyFull != nil {
			caCertificate := r.Configuration.SslMode.VerifyFull.CaCertificate.ValueString()
			clientCertificate := r.Configuration.SslMode.VerifyFull.ClientCertificate.ValueString()
			clientKey := r.Configuration.SslMode.VerifyFull.ClientKey.ValueString()
			clientKeyPassword := new(string)
			if !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsNull() {
				*clientKeyPassword = r.Configuration.SslMode.VerifyFull.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword = nil
			}
			destinationRedisUpdateVerifyFull = &shared.DestinationRedisUpdateVerifyFull{
				CaCertificate:     caCertificate,
				ClientCertificate: clientCertificate,
				ClientKey:         clientKey,
				ClientKeyPassword: clientKeyPassword,
			}
		}
		if destinationRedisUpdateVerifyFull != nil {
			sslMode = &shared.DestinationRedisUpdateSSLModes{
				DestinationRedisUpdateVerifyFull: destinationRedisUpdateVerifyFull,
			}
		}
	}
	var tunnelMethod *shared.DestinationRedisUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationRedisUpdateNoTunnel *shared.DestinationRedisUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationRedisUpdateNoTunnel = &shared.DestinationRedisUpdateNoTunnel{}
		}
		if destinationRedisUpdateNoTunnel != nil {
			tunnelMethod = &shared.DestinationRedisUpdateSSHTunnelMethod{
				DestinationRedisUpdateNoTunnel: destinationRedisUpdateNoTunnel,
			}
		}
		var destinationRedisUpdateSSHKeyAuthentication *shared.DestinationRedisUpdateSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationRedisUpdateSSHKeyAuthentication = &shared.DestinationRedisUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationRedisUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationRedisUpdateSSHTunnelMethod{
				DestinationRedisUpdateSSHKeyAuthentication: destinationRedisUpdateSSHKeyAuthentication,
			}
		}
		var destinationRedisUpdatePasswordAuthentication *shared.DestinationRedisUpdatePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationRedisUpdatePasswordAuthentication = &shared.DestinationRedisUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationRedisUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationRedisUpdateSSHTunnelMethod{
				DestinationRedisUpdatePasswordAuthentication: destinationRedisUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationRedisUpdate{
		CacheType:    cacheType,
		Host:         host,
		Password:     password,
		Port:         port,
		Ssl:          ssl,
		SslMode:      sslMode,
		TunnelMethod: tunnelMethod,
		Username:     username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationRedisPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
