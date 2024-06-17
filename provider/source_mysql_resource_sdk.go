// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMysqlResourceModel) ToSharedSourceMysqlCreateRequest() *shared.SourceMysqlCreateRequest {
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
	var replicationMethod shared.SourceMysqlUpdateMethod
	var sourceMysqlReadChangesUsingBinaryLogCDC *shared.SourceMysqlReadChangesUsingBinaryLogCDC
	if r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC != nil {
		initialWaitingSeconds := new(int64)
		if !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InitialWaitingSeconds.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InitialWaitingSeconds.IsNull() {
			*initialWaitingSeconds = r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InitialWaitingSeconds.ValueInt64()
		} else {
			initialWaitingSeconds = nil
		}
		invalidCdcCursorPositionBehavior := new(shared.SourceMysqlInvalidCDCPositionBehaviorAdvanced)
		if !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InvalidCdcCursorPositionBehavior.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InvalidCdcCursorPositionBehavior.IsNull() {
			*invalidCdcCursorPositionBehavior = shared.SourceMysqlInvalidCDCPositionBehaviorAdvanced(r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InvalidCdcCursorPositionBehavior.ValueString())
		} else {
			invalidCdcCursorPositionBehavior = nil
		}
		serverTimeZone := new(string)
		if !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.ServerTimeZone.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.ServerTimeZone.IsNull() {
			*serverTimeZone = r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.ServerTimeZone.ValueString()
		} else {
			serverTimeZone = nil
		}
		sourceMysqlReadChangesUsingBinaryLogCDC = &shared.SourceMysqlReadChangesUsingBinaryLogCDC{
			InitialWaitingSeconds:            initialWaitingSeconds,
			InvalidCdcCursorPositionBehavior: invalidCdcCursorPositionBehavior,
			ServerTimeZone:                   serverTimeZone,
		}
	}
	if sourceMysqlReadChangesUsingBinaryLogCDC != nil {
		replicationMethod = shared.SourceMysqlUpdateMethod{
			SourceMysqlReadChangesUsingBinaryLogCDC: sourceMysqlReadChangesUsingBinaryLogCDC,
		}
	}
	var sourceMysqlScanChangesWithUserDefinedCursor *shared.SourceMysqlScanChangesWithUserDefinedCursor
	if r.Configuration.ReplicationMethod.ScanChangesWithUserDefinedCursor != nil {
		sourceMysqlScanChangesWithUserDefinedCursor = &shared.SourceMysqlScanChangesWithUserDefinedCursor{}
	}
	if sourceMysqlScanChangesWithUserDefinedCursor != nil {
		replicationMethod = shared.SourceMysqlUpdateMethod{
			SourceMysqlScanChangesWithUserDefinedCursor: sourceMysqlScanChangesWithUserDefinedCursor,
		}
	}
	var sslMode *shared.SourceMysqlSSLModes
	if r.Configuration.SslMode != nil {
		var sourceMysqlPreferred *shared.SourceMysqlPreferred
		if r.Configuration.SslMode.Preferred != nil {
			sourceMysqlPreferred = &shared.SourceMysqlPreferred{}
		}
		if sourceMysqlPreferred != nil {
			sslMode = &shared.SourceMysqlSSLModes{
				SourceMysqlPreferred: sourceMysqlPreferred,
			}
		}
		var sourceMysqlRequired *shared.SourceMysqlRequired
		if r.Configuration.SslMode.Required != nil {
			sourceMysqlRequired = &shared.SourceMysqlRequired{}
		}
		if sourceMysqlRequired != nil {
			sslMode = &shared.SourceMysqlSSLModes{
				SourceMysqlRequired: sourceMysqlRequired,
			}
		}
		var sourceMysqlVerifyCA *shared.SourceMysqlVerifyCA
		if r.Configuration.SslMode.VerifyCA != nil {
			caCertificate := r.Configuration.SslMode.VerifyCA.CaCertificate.ValueString()
			clientCertificate := new(string)
			if !r.Configuration.SslMode.VerifyCA.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.VerifyCA.ClientCertificate.IsNull() {
				*clientCertificate = r.Configuration.SslMode.VerifyCA.ClientCertificate.ValueString()
			} else {
				clientCertificate = nil
			}
			clientKey := new(string)
			if !r.Configuration.SslMode.VerifyCA.ClientKey.IsUnknown() && !r.Configuration.SslMode.VerifyCA.ClientKey.IsNull() {
				*clientKey = r.Configuration.SslMode.VerifyCA.ClientKey.ValueString()
			} else {
				clientKey = nil
			}
			clientKeyPassword := new(string)
			if !r.Configuration.SslMode.VerifyCA.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyCA.ClientKeyPassword.IsNull() {
				*clientKeyPassword = r.Configuration.SslMode.VerifyCA.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword = nil
			}
			sourceMysqlVerifyCA = &shared.SourceMysqlVerifyCA{
				CaCertificate:     caCertificate,
				ClientCertificate: clientCertificate,
				ClientKey:         clientKey,
				ClientKeyPassword: clientKeyPassword,
			}
		}
		if sourceMysqlVerifyCA != nil {
			sslMode = &shared.SourceMysqlSSLModes{
				SourceMysqlVerifyCA: sourceMysqlVerifyCA,
			}
		}
		var sourceMysqlVerifyIdentity *shared.SourceMysqlVerifyIdentity
		if r.Configuration.SslMode.VerifyIdentity != nil {
			caCertificate1 := r.Configuration.SslMode.VerifyIdentity.CaCertificate.ValueString()
			clientCertificate1 := new(string)
			if !r.Configuration.SslMode.VerifyIdentity.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.VerifyIdentity.ClientCertificate.IsNull() {
				*clientCertificate1 = r.Configuration.SslMode.VerifyIdentity.ClientCertificate.ValueString()
			} else {
				clientCertificate1 = nil
			}
			clientKey1 := new(string)
			if !r.Configuration.SslMode.VerifyIdentity.ClientKey.IsUnknown() && !r.Configuration.SslMode.VerifyIdentity.ClientKey.IsNull() {
				*clientKey1 = r.Configuration.SslMode.VerifyIdentity.ClientKey.ValueString()
			} else {
				clientKey1 = nil
			}
			clientKeyPassword1 := new(string)
			if !r.Configuration.SslMode.VerifyIdentity.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyIdentity.ClientKeyPassword.IsNull() {
				*clientKeyPassword1 = r.Configuration.SslMode.VerifyIdentity.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword1 = nil
			}
			sourceMysqlVerifyIdentity = &shared.SourceMysqlVerifyIdentity{
				CaCertificate:     caCertificate1,
				ClientCertificate: clientCertificate1,
				ClientKey:         clientKey1,
				ClientKeyPassword: clientKeyPassword1,
			}
		}
		if sourceMysqlVerifyIdentity != nil {
			sslMode = &shared.SourceMysqlSSLModes{
				SourceMysqlVerifyIdentity: sourceMysqlVerifyIdentity,
			}
		}
	}
	var tunnelMethod *shared.SourceMysqlSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var sourceMysqlNoTunnel *shared.SourceMysqlNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			sourceMysqlNoTunnel = &shared.SourceMysqlNoTunnel{}
		}
		if sourceMysqlNoTunnel != nil {
			tunnelMethod = &shared.SourceMysqlSSHTunnelMethod{
				SourceMysqlNoTunnel: sourceMysqlNoTunnel,
			}
		}
		var sourceMysqlSSHKeyAuthentication *shared.SourceMysqlSSHKeyAuthentication
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
			sourceMysqlSSHKeyAuthentication = &shared.SourceMysqlSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sourceMysqlSSHKeyAuthentication != nil {
			tunnelMethod = &shared.SourceMysqlSSHTunnelMethod{
				SourceMysqlSSHKeyAuthentication: sourceMysqlSSHKeyAuthentication,
			}
		}
		var sourceMysqlPasswordAuthentication *shared.SourceMysqlPasswordAuthentication
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
			sourceMysqlPasswordAuthentication = &shared.SourceMysqlPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if sourceMysqlPasswordAuthentication != nil {
			tunnelMethod = &shared.SourceMysqlSSHTunnelMethod{
				SourceMysqlPasswordAuthentication: sourceMysqlPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceMysql{
		Database:          database,
		Host:              host,
		JdbcURLParams:     jdbcURLParams,
		Password:          password,
		Port:              port,
		ReplicationMethod: replicationMethod,
		SslMode:           sslMode,
		TunnelMethod:      tunnelMethod,
		Username:          username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMysqlCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMysqlResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceMysqlResourceModel) ToSharedSourceMysqlPutRequest() *shared.SourceMysqlPutRequest {
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
	var replicationMethod shared.SourceMysqlUpdateUpdateMethod
	var readChangesUsingBinaryLogCDC *shared.ReadChangesUsingBinaryLogCDC
	if r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC != nil {
		initialWaitingSeconds := new(int64)
		if !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InitialWaitingSeconds.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InitialWaitingSeconds.IsNull() {
			*initialWaitingSeconds = r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InitialWaitingSeconds.ValueInt64()
		} else {
			initialWaitingSeconds = nil
		}
		invalidCdcCursorPositionBehavior := new(shared.SourceMysqlUpdateInvalidCDCPositionBehaviorAdvanced)
		if !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InvalidCdcCursorPositionBehavior.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InvalidCdcCursorPositionBehavior.IsNull() {
			*invalidCdcCursorPositionBehavior = shared.SourceMysqlUpdateInvalidCDCPositionBehaviorAdvanced(r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.InvalidCdcCursorPositionBehavior.ValueString())
		} else {
			invalidCdcCursorPositionBehavior = nil
		}
		serverTimeZone := new(string)
		if !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.ServerTimeZone.IsUnknown() && !r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.ServerTimeZone.IsNull() {
			*serverTimeZone = r.Configuration.ReplicationMethod.ReadChangesUsingBinaryLogCDC.ServerTimeZone.ValueString()
		} else {
			serverTimeZone = nil
		}
		readChangesUsingBinaryLogCDC = &shared.ReadChangesUsingBinaryLogCDC{
			InitialWaitingSeconds:            initialWaitingSeconds,
			InvalidCdcCursorPositionBehavior: invalidCdcCursorPositionBehavior,
			ServerTimeZone:                   serverTimeZone,
		}
	}
	if readChangesUsingBinaryLogCDC != nil {
		replicationMethod = shared.SourceMysqlUpdateUpdateMethod{
			ReadChangesUsingBinaryLogCDC: readChangesUsingBinaryLogCDC,
		}
	}
	var sourceMysqlUpdateScanChangesWithUserDefinedCursor *shared.SourceMysqlUpdateScanChangesWithUserDefinedCursor
	if r.Configuration.ReplicationMethod.ScanChangesWithUserDefinedCursor != nil {
		sourceMysqlUpdateScanChangesWithUserDefinedCursor = &shared.SourceMysqlUpdateScanChangesWithUserDefinedCursor{}
	}
	if sourceMysqlUpdateScanChangesWithUserDefinedCursor != nil {
		replicationMethod = shared.SourceMysqlUpdateUpdateMethod{
			SourceMysqlUpdateScanChangesWithUserDefinedCursor: sourceMysqlUpdateScanChangesWithUserDefinedCursor,
		}
	}
	var sslMode *shared.SourceMysqlUpdateSSLModes
	if r.Configuration.SslMode != nil {
		var preferred *shared.Preferred
		if r.Configuration.SslMode.Preferred != nil {
			preferred = &shared.Preferred{}
		}
		if preferred != nil {
			sslMode = &shared.SourceMysqlUpdateSSLModes{
				Preferred: preferred,
			}
		}
		var required *shared.Required
		if r.Configuration.SslMode.Required != nil {
			required = &shared.Required{}
		}
		if required != nil {
			sslMode = &shared.SourceMysqlUpdateSSLModes{
				Required: required,
			}
		}
		var sourceMysqlUpdateVerifyCA *shared.SourceMysqlUpdateVerifyCA
		if r.Configuration.SslMode.VerifyCA != nil {
			caCertificate := r.Configuration.SslMode.VerifyCA.CaCertificate.ValueString()
			clientCertificate := new(string)
			if !r.Configuration.SslMode.VerifyCA.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.VerifyCA.ClientCertificate.IsNull() {
				*clientCertificate = r.Configuration.SslMode.VerifyCA.ClientCertificate.ValueString()
			} else {
				clientCertificate = nil
			}
			clientKey := new(string)
			if !r.Configuration.SslMode.VerifyCA.ClientKey.IsUnknown() && !r.Configuration.SslMode.VerifyCA.ClientKey.IsNull() {
				*clientKey = r.Configuration.SslMode.VerifyCA.ClientKey.ValueString()
			} else {
				clientKey = nil
			}
			clientKeyPassword := new(string)
			if !r.Configuration.SslMode.VerifyCA.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyCA.ClientKeyPassword.IsNull() {
				*clientKeyPassword = r.Configuration.SslMode.VerifyCA.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword = nil
			}
			sourceMysqlUpdateVerifyCA = &shared.SourceMysqlUpdateVerifyCA{
				CaCertificate:     caCertificate,
				ClientCertificate: clientCertificate,
				ClientKey:         clientKey,
				ClientKeyPassword: clientKeyPassword,
			}
		}
		if sourceMysqlUpdateVerifyCA != nil {
			sslMode = &shared.SourceMysqlUpdateSSLModes{
				SourceMysqlUpdateVerifyCA: sourceMysqlUpdateVerifyCA,
			}
		}
		var verifyIdentity *shared.VerifyIdentity
		if r.Configuration.SslMode.VerifyIdentity != nil {
			caCertificate1 := r.Configuration.SslMode.VerifyIdentity.CaCertificate.ValueString()
			clientCertificate1 := new(string)
			if !r.Configuration.SslMode.VerifyIdentity.ClientCertificate.IsUnknown() && !r.Configuration.SslMode.VerifyIdentity.ClientCertificate.IsNull() {
				*clientCertificate1 = r.Configuration.SslMode.VerifyIdentity.ClientCertificate.ValueString()
			} else {
				clientCertificate1 = nil
			}
			clientKey1 := new(string)
			if !r.Configuration.SslMode.VerifyIdentity.ClientKey.IsUnknown() && !r.Configuration.SslMode.VerifyIdentity.ClientKey.IsNull() {
				*clientKey1 = r.Configuration.SslMode.VerifyIdentity.ClientKey.ValueString()
			} else {
				clientKey1 = nil
			}
			clientKeyPassword1 := new(string)
			if !r.Configuration.SslMode.VerifyIdentity.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyIdentity.ClientKeyPassword.IsNull() {
				*clientKeyPassword1 = r.Configuration.SslMode.VerifyIdentity.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword1 = nil
			}
			verifyIdentity = &shared.VerifyIdentity{
				CaCertificate:     caCertificate1,
				ClientCertificate: clientCertificate1,
				ClientKey:         clientKey1,
				ClientKeyPassword: clientKeyPassword1,
			}
		}
		if verifyIdentity != nil {
			sslMode = &shared.SourceMysqlUpdateSSLModes{
				VerifyIdentity: verifyIdentity,
			}
		}
	}
	var tunnelMethod *shared.SourceMysqlUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var sourceMysqlUpdateNoTunnel *shared.SourceMysqlUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			sourceMysqlUpdateNoTunnel = &shared.SourceMysqlUpdateNoTunnel{}
		}
		if sourceMysqlUpdateNoTunnel != nil {
			tunnelMethod = &shared.SourceMysqlUpdateSSHTunnelMethod{
				SourceMysqlUpdateNoTunnel: sourceMysqlUpdateNoTunnel,
			}
		}
		var sourceMysqlUpdateSSHKeyAuthentication *shared.SourceMysqlUpdateSSHKeyAuthentication
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
			sourceMysqlUpdateSSHKeyAuthentication = &shared.SourceMysqlUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if sourceMysqlUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.SourceMysqlUpdateSSHTunnelMethod{
				SourceMysqlUpdateSSHKeyAuthentication: sourceMysqlUpdateSSHKeyAuthentication,
			}
		}
		var sourceMysqlUpdatePasswordAuthentication *shared.SourceMysqlUpdatePasswordAuthentication
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
			sourceMysqlUpdatePasswordAuthentication = &shared.SourceMysqlUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if sourceMysqlUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.SourceMysqlUpdateSSHTunnelMethod{
				SourceMysqlUpdatePasswordAuthentication: sourceMysqlUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceMysqlUpdate{
		Database:          database,
		Host:              host,
		JdbcURLParams:     jdbcURLParams,
		Password:          password,
		Port:              port,
		ReplicationMethod: replicationMethod,
		SslMode:           sslMode,
		TunnelMethod:      tunnelMethod,
		Username:          username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMysqlPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
