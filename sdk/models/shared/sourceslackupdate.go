// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"time"
)

type SourceSlackUpdateSchemasOptionTitle string

const (
	SourceSlackUpdateSchemasOptionTitleAPITokenCredentials SourceSlackUpdateSchemasOptionTitle = "API Token Credentials"
)

func (e SourceSlackUpdateSchemasOptionTitle) ToPointer() *SourceSlackUpdateSchemasOptionTitle {
	return &e
}
func (e *SourceSlackUpdateSchemasOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "API Token Credentials":
		*e = SourceSlackUpdateSchemasOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSlackUpdateSchemasOptionTitle: %v", v)
	}
}

type SourceSlackUpdateAPIToken struct {
	// A Slack bot token. See the <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> for instructions on how to generate it.
	APIToken    string                              `json:"api_token"`
	optionTitle SourceSlackUpdateSchemasOptionTitle `const:"API Token Credentials" json:"option_title"`
}

func (s SourceSlackUpdateAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSlackUpdateAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSlackUpdateAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceSlackUpdateAPIToken) GetOptionTitle() SourceSlackUpdateSchemasOptionTitle {
	return SourceSlackUpdateSchemasOptionTitleAPITokenCredentials
}

type SourceSlackUpdateOptionTitle string

const (
	SourceSlackUpdateOptionTitleDefaultOAuth20Authorization SourceSlackUpdateOptionTitle = "Default OAuth2.0 authorization"
)

func (e SourceSlackUpdateOptionTitle) ToPointer() *SourceSlackUpdateOptionTitle {
	return &e
}
func (e *SourceSlackUpdateOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Default OAuth2.0 authorization":
		*e = SourceSlackUpdateOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSlackUpdateOptionTitle: %v", v)
	}
}

type SignInViaSlackOAuth struct {
	// Slack access_token. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help generating the token.
	AccessToken string `json:"access_token"`
	// Slack client_id. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this id.
	ClientID string `json:"client_id"`
	// Slack client_secret. See our <a href="https://docs.airbyte.com/integrations/sources/slack">docs</a> if you need help finding this secret.
	ClientSecret string                       `json:"client_secret"`
	optionTitle  SourceSlackUpdateOptionTitle `const:"Default OAuth2.0 authorization" json:"option_title"`
}

func (s SignInViaSlackOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SignInViaSlackOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SignInViaSlackOAuth) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SignInViaSlackOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SignInViaSlackOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SignInViaSlackOAuth) GetOptionTitle() SourceSlackUpdateOptionTitle {
	return SourceSlackUpdateOptionTitleDefaultOAuth20Authorization
}

type SourceSlackUpdateAuthenticationMechanismType string

const (
	SourceSlackUpdateAuthenticationMechanismTypeSignInViaSlackOAuth       SourceSlackUpdateAuthenticationMechanismType = "Sign in via Slack (OAuth)"
	SourceSlackUpdateAuthenticationMechanismTypeSourceSlackUpdateAPIToken SourceSlackUpdateAuthenticationMechanismType = "source-slack-update_API Token"
)

// SourceSlackUpdateAuthenticationMechanism - Choose how to authenticate into Slack
type SourceSlackUpdateAuthenticationMechanism struct {
	SignInViaSlackOAuth       *SignInViaSlackOAuth
	SourceSlackUpdateAPIToken *SourceSlackUpdateAPIToken

	Type SourceSlackUpdateAuthenticationMechanismType
}

func CreateSourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth(signInViaSlackOAuth SignInViaSlackOAuth) SourceSlackUpdateAuthenticationMechanism {
	typ := SourceSlackUpdateAuthenticationMechanismTypeSignInViaSlackOAuth

	return SourceSlackUpdateAuthenticationMechanism{
		SignInViaSlackOAuth: &signInViaSlackOAuth,
		Type:                typ,
	}
}

func CreateSourceSlackUpdateAuthenticationMechanismSourceSlackUpdateAPIToken(sourceSlackUpdateAPIToken SourceSlackUpdateAPIToken) SourceSlackUpdateAuthenticationMechanism {
	typ := SourceSlackUpdateAuthenticationMechanismTypeSourceSlackUpdateAPIToken

	return SourceSlackUpdateAuthenticationMechanism{
		SourceSlackUpdateAPIToken: &sourceSlackUpdateAPIToken,
		Type:                      typ,
	}
}

func (u *SourceSlackUpdateAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	var sourceSlackUpdateAPIToken SourceSlackUpdateAPIToken = SourceSlackUpdateAPIToken{}
	if err := utils.UnmarshalJSON(data, &sourceSlackUpdateAPIToken, "", true, true); err == nil {
		u.SourceSlackUpdateAPIToken = &sourceSlackUpdateAPIToken
		u.Type = SourceSlackUpdateAuthenticationMechanismTypeSourceSlackUpdateAPIToken
		return nil
	}

	var signInViaSlackOAuth SignInViaSlackOAuth = SignInViaSlackOAuth{}
	if err := utils.UnmarshalJSON(data, &signInViaSlackOAuth, "", true, true); err == nil {
		u.SignInViaSlackOAuth = &signInViaSlackOAuth
		u.Type = SourceSlackUpdateAuthenticationMechanismTypeSignInViaSlackOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSlackUpdateAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SignInViaSlackOAuth != nil {
		return utils.MarshalJSON(u.SignInViaSlackOAuth, "", true)
	}

	if u.SourceSlackUpdateAPIToken != nil {
		return utils.MarshalJSON(u.SourceSlackUpdateAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceSlackUpdate struct {
	// A channel name list (without leading '#' char) which limit the channels from which you'd like to sync. Empty list means no filter.
	ChannelFilter []string `json:"channel_filter,omitempty"`
	// Choose how to authenticate into Slack
	Credentials *SourceSlackUpdateAuthenticationMechanism `json:"credentials,omitempty"`
	// Whether to read information from private channels that the bot is already in.  If false, only public channels will be read.  If true, the bot must be manually added to private channels.
	IncludePrivateChannels *bool `default:"false" json:"include_private_channels"`
	// Whether to join all channels or to sync data only from channels the bot is already in.  If false, you'll need to manually add the bot to all the channels from which you'd like to sync messages.
	JoinChannels *bool `default:"true" json:"join_channels"`
	// How far into the past to look for messages in threads, default is 0 days
	LookbackWindow *int64 `default:"0" json:"lookback_window"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceSlackUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSlackUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSlackUpdate) GetChannelFilter() []string {
	if o == nil {
		return nil
	}
	return o.ChannelFilter
}

func (o *SourceSlackUpdate) GetCredentials() *SourceSlackUpdateAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceSlackUpdate) GetIncludePrivateChannels() *bool {
	if o == nil {
		return nil
	}
	return o.IncludePrivateChannels
}

func (o *SourceSlackUpdate) GetJoinChannels() *bool {
	if o == nil {
		return nil
	}
	return o.JoinChannels
}

func (o *SourceSlackUpdate) GetLookbackWindow() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindow
}

func (o *SourceSlackUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
