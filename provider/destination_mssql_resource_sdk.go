// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMssqlResourceModel) ToSharedDestinationMssqlCreateRequest() *shared.DestinationMssqlCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
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
	rawDataSchema := new(string)
	if !r.Configuration.RawDataSchema.IsUnknown() && !r.Configuration.RawDataSchema.IsNull() {
		*rawDataSchema = r.Configuration.RawDataSchema.ValueString()
	} else {
		rawDataSchema = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	var sslMethod *shared.DestinationMssqlSSLMethod
	if r.Configuration.SslMethod != nil {
		var destinationMssqlEncryptedTrustServerCertificate *shared.DestinationMssqlEncryptedTrustServerCertificate
		if r.Configuration.SslMethod.EncryptedTrustServerCertificate != nil {
			destinationMssqlEncryptedTrustServerCertificate = &shared.DestinationMssqlEncryptedTrustServerCertificate{}
		}
		if destinationMssqlEncryptedTrustServerCertificate != nil {
			sslMethod = &shared.DestinationMssqlSSLMethod{
				DestinationMssqlEncryptedTrustServerCertificate: destinationMssqlEncryptedTrustServerCertificate,
			}
		}
		var destinationMssqlEncryptedVerifyCertificate *shared.DestinationMssqlEncryptedVerifyCertificate
		if r.Configuration.SslMethod.EncryptedVerifyCertificate != nil {
			hostNameInCertificate := new(string)
			if !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsUnknown() && !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsNull() {
				*hostNameInCertificate = r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.ValueString()
			} else {
				hostNameInCertificate = nil
			}
			destinationMssqlEncryptedVerifyCertificate = &shared.DestinationMssqlEncryptedVerifyCertificate{
				HostNameInCertificate: hostNameInCertificate,
			}
		}
		if destinationMssqlEncryptedVerifyCertificate != nil {
			sslMethod = &shared.DestinationMssqlSSLMethod{
				DestinationMssqlEncryptedVerifyCertificate: destinationMssqlEncryptedVerifyCertificate,
			}
		}
	}
	var tunnelMethod *shared.DestinationMssqlSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMssqlNoTunnel *shared.DestinationMssqlNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationMssqlNoTunnel = &shared.DestinationMssqlNoTunnel{}
		}
		if destinationMssqlNoTunnel != nil {
			tunnelMethod = &shared.DestinationMssqlSSHTunnelMethod{
				DestinationMssqlNoTunnel: destinationMssqlNoTunnel,
			}
		}
		var destinationMssqlSSHKeyAuthentication *shared.DestinationMssqlSSHKeyAuthentication
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
			destinationMssqlSSHKeyAuthentication = &shared.DestinationMssqlSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationMssqlSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMssqlSSHTunnelMethod{
				DestinationMssqlSSHKeyAuthentication: destinationMssqlSSHKeyAuthentication,
			}
		}
		var destinationMssqlPasswordAuthentication *shared.DestinationMssqlPasswordAuthentication
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
			destinationMssqlPasswordAuthentication = &shared.DestinationMssqlPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMssqlPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMssqlSSHTunnelMethod{
				DestinationMssqlPasswordAuthentication: destinationMssqlPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMssql{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		RawDataSchema: rawDataSchema,
		Schema:        schema,
		SslMethod:     sslMethod,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMssqlCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMssqlResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationMssqlResourceModel) ToSharedDestinationMssqlPutRequest() *shared.DestinationMssqlPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
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
	rawDataSchema := new(string)
	if !r.Configuration.RawDataSchema.IsUnknown() && !r.Configuration.RawDataSchema.IsNull() {
		*rawDataSchema = r.Configuration.RawDataSchema.ValueString()
	} else {
		rawDataSchema = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	var sslMethod *shared.SSLMethod
	if r.Configuration.SslMethod != nil {
		var encryptedTrustServerCertificate *shared.EncryptedTrustServerCertificate
		if r.Configuration.SslMethod.EncryptedTrustServerCertificate != nil {
			encryptedTrustServerCertificate = &shared.EncryptedTrustServerCertificate{}
		}
		if encryptedTrustServerCertificate != nil {
			sslMethod = &shared.SSLMethod{
				EncryptedTrustServerCertificate: encryptedTrustServerCertificate,
			}
		}
		var encryptedVerifyCertificate *shared.EncryptedVerifyCertificate
		if r.Configuration.SslMethod.EncryptedVerifyCertificate != nil {
			hostNameInCertificate := new(string)
			if !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsUnknown() && !r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.IsNull() {
				*hostNameInCertificate = r.Configuration.SslMethod.EncryptedVerifyCertificate.HostNameInCertificate.ValueString()
			} else {
				hostNameInCertificate = nil
			}
			encryptedVerifyCertificate = &shared.EncryptedVerifyCertificate{
				HostNameInCertificate: hostNameInCertificate,
			}
		}
		if encryptedVerifyCertificate != nil {
			sslMethod = &shared.SSLMethod{
				EncryptedVerifyCertificate: encryptedVerifyCertificate,
			}
		}
	}
	var tunnelMethod *shared.DestinationMssqlUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMssqlUpdateNoTunnel *shared.DestinationMssqlUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationMssqlUpdateNoTunnel = &shared.DestinationMssqlUpdateNoTunnel{}
		}
		if destinationMssqlUpdateNoTunnel != nil {
			tunnelMethod = &shared.DestinationMssqlUpdateSSHTunnelMethod{
				DestinationMssqlUpdateNoTunnel: destinationMssqlUpdateNoTunnel,
			}
		}
		var destinationMssqlUpdateSSHKeyAuthentication *shared.DestinationMssqlUpdateSSHKeyAuthentication
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
			destinationMssqlUpdateSSHKeyAuthentication = &shared.DestinationMssqlUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationMssqlUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMssqlUpdateSSHTunnelMethod{
				DestinationMssqlUpdateSSHKeyAuthentication: destinationMssqlUpdateSSHKeyAuthentication,
			}
		}
		var destinationMssqlUpdatePasswordAuthentication *shared.DestinationMssqlUpdatePasswordAuthentication
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
			destinationMssqlUpdatePasswordAuthentication = &shared.DestinationMssqlUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMssqlUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMssqlUpdateSSHTunnelMethod{
				DestinationMssqlUpdatePasswordAuthentication: destinationMssqlUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMssqlUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		RawDataSchema: rawDataSchema,
		Schema:        schema,
		SslMethod:     sslMethod,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMssqlPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
