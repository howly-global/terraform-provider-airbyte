// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGitlabResourceModel) ToSharedSourceGitlabCreateRequest() *shared.SourceGitlabCreateRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var credentials shared.SourceGitlabAuthorizationMethod
	var sourceGitlabOAuth20 *shared.SourceGitlabOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceGitlabOAuth20 = &shared.SourceGitlabOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceGitlabOAuth20 != nil {
		credentials = shared.SourceGitlabAuthorizationMethod{
			SourceGitlabOAuth20: sourceGitlabOAuth20,
		}
	}
	var sourceGitlabPrivateToken *shared.SourceGitlabPrivateToken
	if r.Configuration.Credentials.PrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.PrivateToken.AccessToken.ValueString()
		sourceGitlabPrivateToken = &shared.SourceGitlabPrivateToken{
			AccessToken: accessToken1,
		}
	}
	if sourceGitlabPrivateToken != nil {
		credentials = shared.SourceGitlabAuthorizationMethod{
			SourceGitlabPrivateToken: sourceGitlabPrivateToken,
		}
	}
	groups := new(string)
	if !r.Configuration.Groups.IsUnknown() && !r.Configuration.Groups.IsNull() {
		*groups = r.Configuration.Groups.ValueString()
	} else {
		groups = nil
	}
	var groupsList []string = []string{}
	for _, groupsListItem := range r.Configuration.GroupsList {
		groupsList = append(groupsList, groupsListItem.ValueString())
	}
	projects := new(string)
	if !r.Configuration.Projects.IsUnknown() && !r.Configuration.Projects.IsNull() {
		*projects = r.Configuration.Projects.ValueString()
	} else {
		projects = nil
	}
	var projectsList []string = []string{}
	for _, projectsListItem := range r.Configuration.ProjectsList {
		projectsList = append(projectsList, projectsListItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceGitlab{
		APIURL:       apiURL,
		Credentials:  credentials,
		Groups:       groups,
		GroupsList:   groupsList,
		Projects:     projects,
		ProjectsList: projectsList,
		StartDate:    startDate,
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
	out := shared.SourceGitlabCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGitlabResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGitlabResourceModel) ToSharedSourceGitlabPutRequest() *shared.SourceGitlabPutRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var credentials shared.SourceGitlabUpdateAuthorizationMethod
	var sourceGitlabUpdateOAuth20 *shared.SourceGitlabUpdateOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceGitlabUpdateOAuth20 = &shared.SourceGitlabUpdateOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceGitlabUpdateOAuth20 != nil {
		credentials = shared.SourceGitlabUpdateAuthorizationMethod{
			SourceGitlabUpdateOAuth20: sourceGitlabUpdateOAuth20,
		}
	}
	var privateToken *shared.PrivateToken
	if r.Configuration.Credentials.PrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.PrivateToken.AccessToken.ValueString()
		privateToken = &shared.PrivateToken{
			AccessToken: accessToken1,
		}
	}
	if privateToken != nil {
		credentials = shared.SourceGitlabUpdateAuthorizationMethod{
			PrivateToken: privateToken,
		}
	}
	groups := new(string)
	if !r.Configuration.Groups.IsUnknown() && !r.Configuration.Groups.IsNull() {
		*groups = r.Configuration.Groups.ValueString()
	} else {
		groups = nil
	}
	var groupsList []string = []string{}
	for _, groupsListItem := range r.Configuration.GroupsList {
		groupsList = append(groupsList, groupsListItem.ValueString())
	}
	projects := new(string)
	if !r.Configuration.Projects.IsUnknown() && !r.Configuration.Projects.IsNull() {
		*projects = r.Configuration.Projects.ValueString()
	} else {
		projects = nil
	}
	var projectsList []string = []string{}
	for _, projectsListItem := range r.Configuration.ProjectsList {
		projectsList = append(projectsList, projectsListItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceGitlabUpdate{
		APIURL:       apiURL,
		Credentials:  credentials,
		Groups:       groups,
		GroupsList:   groupsList,
		Projects:     projects,
		ProjectsList: projectsList,
		StartDate:    startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGitlabPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}