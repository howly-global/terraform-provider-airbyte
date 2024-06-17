// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
	"github.com/howly-global/terraform-provider-airbyte/sdk/types"
)

type SourcePinterestUpdateAuthMethod string

const (
	SourcePinterestUpdateAuthMethodOauth20 SourcePinterestUpdateAuthMethod = "oauth2.0"
)

func (e SourcePinterestUpdateAuthMethod) ToPointer() *SourcePinterestUpdateAuthMethod {
	return &e
}
func (e *SourcePinterestUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourcePinterestUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestUpdateAuthMethod: %v", v)
	}
}

type OAuth20 struct {
	authMethod SourcePinterestUpdateAuthMethod `const:"oauth2.0" json:"auth_method"`
	// The Client ID of your OAuth application
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken string `json:"refresh_token"`
}

func (o OAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OAuth20) GetAuthMethod() SourcePinterestUpdateAuthMethod {
	return SourcePinterestUpdateAuthMethodOauth20
}

func (o *OAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *OAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *OAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

// SourcePinterestUpdateValidEnums - An enumeration.
type SourcePinterestUpdateValidEnums string

const (
	SourcePinterestUpdateValidEnumsIndividual SourcePinterestUpdateValidEnums = "INDIVIDUAL"
	SourcePinterestUpdateValidEnumsHousehold  SourcePinterestUpdateValidEnums = "HOUSEHOLD"
)

func (e SourcePinterestUpdateValidEnums) ToPointer() *SourcePinterestUpdateValidEnums {
	return &e
}
func (e *SourcePinterestUpdateValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INDIVIDUAL":
		fallthrough
	case "HOUSEHOLD":
		*e = SourcePinterestUpdateValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestUpdateValidEnums: %v", v)
	}
}

// ClickWindowDays - Number of days to use as the conversion attribution window for a pin click action.
type ClickWindowDays int64

const (
	ClickWindowDaysZero     ClickWindowDays = 0
	ClickWindowDaysOne      ClickWindowDays = 1
	ClickWindowDaysSeven    ClickWindowDays = 7
	ClickWindowDaysFourteen ClickWindowDays = 14
	ClickWindowDaysThirty   ClickWindowDays = 30
	ClickWindowDaysSixty    ClickWindowDays = 60
)

func (e ClickWindowDays) ToPointer() *ClickWindowDays {
	return &e
}
func (e *ClickWindowDays) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 14:
		fallthrough
	case 30:
		fallthrough
	case 60:
		*e = ClickWindowDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClickWindowDays: %v", v)
	}
}

// SourcePinterestUpdateSchemasValidEnums - An enumeration.
type SourcePinterestUpdateSchemasValidEnums string

const (
	SourcePinterestUpdateSchemasValidEnumsAdvertiserID                                 SourcePinterestUpdateSchemasValidEnums = "ADVERTISER_ID"
	SourcePinterestUpdateSchemasValidEnumsAdAccountID                                  SourcePinterestUpdateSchemasValidEnums = "AD_ACCOUNT_ID"
	SourcePinterestUpdateSchemasValidEnumsAdGroupEntityStatus                          SourcePinterestUpdateSchemasValidEnums = "AD_GROUP_ENTITY_STATUS"
	SourcePinterestUpdateSchemasValidEnumsAdGroupID                                    SourcePinterestUpdateSchemasValidEnums = "AD_GROUP_ID"
	SourcePinterestUpdateSchemasValidEnumsAdID                                         SourcePinterestUpdateSchemasValidEnums = "AD_ID"
	SourcePinterestUpdateSchemasValidEnumsCampaignDailySpendCap                        SourcePinterestUpdateSchemasValidEnums = "CAMPAIGN_DAILY_SPEND_CAP"
	SourcePinterestUpdateSchemasValidEnumsCampaignEntityStatus                         SourcePinterestUpdateSchemasValidEnums = "CAMPAIGN_ENTITY_STATUS"
	SourcePinterestUpdateSchemasValidEnumsCampaignID                                   SourcePinterestUpdateSchemasValidEnums = "CAMPAIGN_ID"
	SourcePinterestUpdateSchemasValidEnumsCampaignLifetimeSpendCap                     SourcePinterestUpdateSchemasValidEnums = "CAMPAIGN_LIFETIME_SPEND_CAP"
	SourcePinterestUpdateSchemasValidEnumsCampaignName                                 SourcePinterestUpdateSchemasValidEnums = "CAMPAIGN_NAME"
	SourcePinterestUpdateSchemasValidEnumsCheckoutRoas                                 SourcePinterestUpdateSchemasValidEnums = "CHECKOUT_ROAS"
	SourcePinterestUpdateSchemasValidEnumsClickthrough1                                SourcePinterestUpdateSchemasValidEnums = "CLICKTHROUGH_1"
	SourcePinterestUpdateSchemasValidEnumsClickthrough1Gross                           SourcePinterestUpdateSchemasValidEnums = "CLICKTHROUGH_1_GROSS"
	SourcePinterestUpdateSchemasValidEnumsClickthrough2                                SourcePinterestUpdateSchemasValidEnums = "CLICKTHROUGH_2"
	SourcePinterestUpdateSchemasValidEnumsCpcInMicroDollar                             SourcePinterestUpdateSchemasValidEnums = "CPC_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsCpmInDollar                                  SourcePinterestUpdateSchemasValidEnums = "CPM_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsCpmInMicroDollar                             SourcePinterestUpdateSchemasValidEnums = "CPM_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsCtr                                          SourcePinterestUpdateSchemasValidEnums = "CTR"
	SourcePinterestUpdateSchemasValidEnumsCtr2                                         SourcePinterestUpdateSchemasValidEnums = "CTR_2"
	SourcePinterestUpdateSchemasValidEnumsEcpcvInDollar                                SourcePinterestUpdateSchemasValidEnums = "ECPCV_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEcpcvP95InDollar                             SourcePinterestUpdateSchemasValidEnums = "ECPCV_P95_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEcpcInDollar                                 SourcePinterestUpdateSchemasValidEnums = "ECPC_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEcpcInMicroDollar                            SourcePinterestUpdateSchemasValidEnums = "ECPC_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEcpeInDollar                                 SourcePinterestUpdateSchemasValidEnums = "ECPE_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEcpmInMicroDollar                            SourcePinterestUpdateSchemasValidEnums = "ECPM_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEcpvInDollar                                 SourcePinterestUpdateSchemasValidEnums = "ECPV_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsEctr                                         SourcePinterestUpdateSchemasValidEnums = "ECTR"
	SourcePinterestUpdateSchemasValidEnumsEengagementRate                              SourcePinterestUpdateSchemasValidEnums = "EENGAGEMENT_RATE"
	SourcePinterestUpdateSchemasValidEnumsEngagement1                                  SourcePinterestUpdateSchemasValidEnums = "ENGAGEMENT_1"
	SourcePinterestUpdateSchemasValidEnumsEngagement2                                  SourcePinterestUpdateSchemasValidEnums = "ENGAGEMENT_2"
	SourcePinterestUpdateSchemasValidEnumsEngagementRate                               SourcePinterestUpdateSchemasValidEnums = "ENGAGEMENT_RATE"
	SourcePinterestUpdateSchemasValidEnumsIdeaPinProductTagVisit1                      SourcePinterestUpdateSchemasValidEnums = "IDEA_PIN_PRODUCT_TAG_VISIT_1"
	SourcePinterestUpdateSchemasValidEnumsIdeaPinProductTagVisit2                      SourcePinterestUpdateSchemasValidEnums = "IDEA_PIN_PRODUCT_TAG_VISIT_2"
	SourcePinterestUpdateSchemasValidEnumsImpression1                                  SourcePinterestUpdateSchemasValidEnums = "IMPRESSION_1"
	SourcePinterestUpdateSchemasValidEnumsImpression1Gross                             SourcePinterestUpdateSchemasValidEnums = "IMPRESSION_1_GROSS"
	SourcePinterestUpdateSchemasValidEnumsImpression2                                  SourcePinterestUpdateSchemasValidEnums = "IMPRESSION_2"
	SourcePinterestUpdateSchemasValidEnumsInappCheckoutCostPerAction                   SourcePinterestUpdateSchemasValidEnums = "INAPP_CHECKOUT_COST_PER_ACTION"
	SourcePinterestUpdateSchemasValidEnumsOutboundClick1                               SourcePinterestUpdateSchemasValidEnums = "OUTBOUND_CLICK_1"
	SourcePinterestUpdateSchemasValidEnumsOutboundClick2                               SourcePinterestUpdateSchemasValidEnums = "OUTBOUND_CLICK_2"
	SourcePinterestUpdateSchemasValidEnumsPageVisitCostPerAction                       SourcePinterestUpdateSchemasValidEnums = "PAGE_VISIT_COST_PER_ACTION"
	SourcePinterestUpdateSchemasValidEnumsPageVisitRoas                                SourcePinterestUpdateSchemasValidEnums = "PAGE_VISIT_ROAS"
	SourcePinterestUpdateSchemasValidEnumsPaidImpression                               SourcePinterestUpdateSchemasValidEnums = "PAID_IMPRESSION"
	SourcePinterestUpdateSchemasValidEnumsPinID                                        SourcePinterestUpdateSchemasValidEnums = "PIN_ID"
	SourcePinterestUpdateSchemasValidEnumsPinPromotionID                               SourcePinterestUpdateSchemasValidEnums = "PIN_PROMOTION_ID"
	SourcePinterestUpdateSchemasValidEnumsRepin1                                       SourcePinterestUpdateSchemasValidEnums = "REPIN_1"
	SourcePinterestUpdateSchemasValidEnumsRepin2                                       SourcePinterestUpdateSchemasValidEnums = "REPIN_2"
	SourcePinterestUpdateSchemasValidEnumsRepinRate                                    SourcePinterestUpdateSchemasValidEnums = "REPIN_RATE"
	SourcePinterestUpdateSchemasValidEnumsSpendInDollar                                SourcePinterestUpdateSchemasValidEnums = "SPEND_IN_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsSpendInMicroDollar                           SourcePinterestUpdateSchemasValidEnums = "SPEND_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalCheckout                                SourcePinterestUpdateSchemasValidEnums = "TOTAL_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalCheckoutValueInMicroDollar              SourcePinterestUpdateSchemasValidEnums = "TOTAL_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalClickthrough                            SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICKTHROUGH"
	SourcePinterestUpdateSchemasValidEnumsTotalClickAddToCart                          SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICK_ADD_TO_CART"
	SourcePinterestUpdateSchemasValidEnumsTotalClickCheckout                           SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICK_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalClickCheckoutValueInMicroDollar         SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalClickLead                               SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICK_LEAD"
	SourcePinterestUpdateSchemasValidEnumsTotalClickSignup                             SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICK_SIGNUP"
	SourcePinterestUpdateSchemasValidEnumsTotalClickSignupValueInMicroDollar           SourcePinterestUpdateSchemasValidEnums = "TOTAL_CLICK_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalConversions                             SourcePinterestUpdateSchemasValidEnums = "TOTAL_CONVERSIONS"
	SourcePinterestUpdateSchemasValidEnumsTotalCustom                                  SourcePinterestUpdateSchemasValidEnums = "TOTAL_CUSTOM"
	SourcePinterestUpdateSchemasValidEnumsTotalEngagement                              SourcePinterestUpdateSchemasValidEnums = "TOTAL_ENGAGEMENT"
	SourcePinterestUpdateSchemasValidEnumsTotalEngagementCheckout                      SourcePinterestUpdateSchemasValidEnums = "TOTAL_ENGAGEMENT_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalEngagementCheckoutValueInMicroDollar    SourcePinterestUpdateSchemasValidEnums = "TOTAL_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalEngagementLead                          SourcePinterestUpdateSchemasValidEnums = "TOTAL_ENGAGEMENT_LEAD"
	SourcePinterestUpdateSchemasValidEnumsTotalEngagementSignup                        SourcePinterestUpdateSchemasValidEnums = "TOTAL_ENGAGEMENT_SIGNUP"
	SourcePinterestUpdateSchemasValidEnumsTotalEngagementSignupValueInMicroDollar      SourcePinterestUpdateSchemasValidEnums = "TOTAL_ENGAGEMENT_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalIdeaPinProductTagVisit                  SourcePinterestUpdateSchemasValidEnums = "TOTAL_IDEA_PIN_PRODUCT_TAG_VISIT"
	SourcePinterestUpdateSchemasValidEnumsTotalImpressionFrequency                     SourcePinterestUpdateSchemasValidEnums = "TOTAL_IMPRESSION_FREQUENCY"
	SourcePinterestUpdateSchemasValidEnumsTotalImpressionUser                          SourcePinterestUpdateSchemasValidEnums = "TOTAL_IMPRESSION_USER"
	SourcePinterestUpdateSchemasValidEnumsTotalLead                                    SourcePinterestUpdateSchemasValidEnums = "TOTAL_LEAD"
	SourcePinterestUpdateSchemasValidEnumsTotalOfflineCheckout                         SourcePinterestUpdateSchemasValidEnums = "TOTAL_OFFLINE_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalPageVisit                               SourcePinterestUpdateSchemasValidEnums = "TOTAL_PAGE_VISIT"
	SourcePinterestUpdateSchemasValidEnumsTotalRepinRate                               SourcePinterestUpdateSchemasValidEnums = "TOTAL_REPIN_RATE"
	SourcePinterestUpdateSchemasValidEnumsTotalSignup                                  SourcePinterestUpdateSchemasValidEnums = "TOTAL_SIGNUP"
	SourcePinterestUpdateSchemasValidEnumsTotalSignupValueInMicroDollar                SourcePinterestUpdateSchemasValidEnums = "TOTAL_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalVideo3SecViews                          SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_3SEC_VIEWS"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoAvgWatchtimeInSecond               SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_AVG_WATCHTIME_IN_SECOND"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoMrcViews                           SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_MRC_VIEWS"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoP0Combined                         SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_P0_COMBINED"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoP100Complete                       SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_P100_COMPLETE"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoP25Combined                        SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_P25_COMBINED"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoP50Combined                        SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_P50_COMBINED"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoP75Combined                        SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_P75_COMBINED"
	SourcePinterestUpdateSchemasValidEnumsTotalVideoP95Combined                        SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIDEO_P95_COMBINED"
	SourcePinterestUpdateSchemasValidEnumsTotalViewAddToCart                           SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIEW_ADD_TO_CART"
	SourcePinterestUpdateSchemasValidEnumsTotalViewCheckout                            SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIEW_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalViewCheckoutValueInMicroDollar          SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalViewLead                                SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIEW_LEAD"
	SourcePinterestUpdateSchemasValidEnumsTotalViewSignup                              SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIEW_SIGNUP"
	SourcePinterestUpdateSchemasValidEnumsTotalViewSignupValueInMicroDollar            SourcePinterestUpdateSchemasValidEnums = "TOTAL_VIEW_SIGNUP_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalWebCheckout                             SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalWebCheckoutValueInMicroDollar           SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalWebClickCheckout                        SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_CLICK_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalWebClickCheckoutValueInMicroDollar      SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalWebEngagementCheckout                   SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_ENGAGEMENT_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalWebEngagementCheckoutValueInMicroDollar SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsTotalWebSessions                             SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_SESSIONS"
	SourcePinterestUpdateSchemasValidEnumsTotalWebViewCheckout                         SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_VIEW_CHECKOUT"
	SourcePinterestUpdateSchemasValidEnumsTotalWebViewCheckoutValueInMicroDollar       SourcePinterestUpdateSchemasValidEnums = "TOTAL_WEB_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR"
	SourcePinterestUpdateSchemasValidEnumsVideo3SecViews2                              SourcePinterestUpdateSchemasValidEnums = "VIDEO_3SEC_VIEWS_2"
	SourcePinterestUpdateSchemasValidEnumsVideoLength                                  SourcePinterestUpdateSchemasValidEnums = "VIDEO_LENGTH"
	SourcePinterestUpdateSchemasValidEnumsVideoMrcViews2                               SourcePinterestUpdateSchemasValidEnums = "VIDEO_MRC_VIEWS_2"
	SourcePinterestUpdateSchemasValidEnumsVideoP0Combined2                             SourcePinterestUpdateSchemasValidEnums = "VIDEO_P0_COMBINED_2"
	SourcePinterestUpdateSchemasValidEnumsVideoP100Complete2                           SourcePinterestUpdateSchemasValidEnums = "VIDEO_P100_COMPLETE_2"
	SourcePinterestUpdateSchemasValidEnumsVideoP25Combined2                            SourcePinterestUpdateSchemasValidEnums = "VIDEO_P25_COMBINED_2"
	SourcePinterestUpdateSchemasValidEnumsVideoP50Combined2                            SourcePinterestUpdateSchemasValidEnums = "VIDEO_P50_COMBINED_2"
	SourcePinterestUpdateSchemasValidEnumsVideoP75Combined2                            SourcePinterestUpdateSchemasValidEnums = "VIDEO_P75_COMBINED_2"
	SourcePinterestUpdateSchemasValidEnumsVideoP95Combined2                            SourcePinterestUpdateSchemasValidEnums = "VIDEO_P95_COMBINED_2"
	SourcePinterestUpdateSchemasValidEnumsWebCheckoutCostPerAction                     SourcePinterestUpdateSchemasValidEnums = "WEB_CHECKOUT_COST_PER_ACTION"
	SourcePinterestUpdateSchemasValidEnumsWebCheckoutRoas                              SourcePinterestUpdateSchemasValidEnums = "WEB_CHECKOUT_ROAS"
	SourcePinterestUpdateSchemasValidEnumsWebSessions1                                 SourcePinterestUpdateSchemasValidEnums = "WEB_SESSIONS_1"
	SourcePinterestUpdateSchemasValidEnumsWebSessions2                                 SourcePinterestUpdateSchemasValidEnums = "WEB_SESSIONS_2"
)

func (e SourcePinterestUpdateSchemasValidEnums) ToPointer() *SourcePinterestUpdateSchemasValidEnums {
	return &e
}
func (e *SourcePinterestUpdateSchemasValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADVERTISER_ID":
		fallthrough
	case "AD_ACCOUNT_ID":
		fallthrough
	case "AD_GROUP_ENTITY_STATUS":
		fallthrough
	case "AD_GROUP_ID":
		fallthrough
	case "AD_ID":
		fallthrough
	case "CAMPAIGN_DAILY_SPEND_CAP":
		fallthrough
	case "CAMPAIGN_ENTITY_STATUS":
		fallthrough
	case "CAMPAIGN_ID":
		fallthrough
	case "CAMPAIGN_LIFETIME_SPEND_CAP":
		fallthrough
	case "CAMPAIGN_NAME":
		fallthrough
	case "CHECKOUT_ROAS":
		fallthrough
	case "CLICKTHROUGH_1":
		fallthrough
	case "CLICKTHROUGH_1_GROSS":
		fallthrough
	case "CLICKTHROUGH_2":
		fallthrough
	case "CPC_IN_MICRO_DOLLAR":
		fallthrough
	case "CPM_IN_DOLLAR":
		fallthrough
	case "CPM_IN_MICRO_DOLLAR":
		fallthrough
	case "CTR":
		fallthrough
	case "CTR_2":
		fallthrough
	case "ECPCV_IN_DOLLAR":
		fallthrough
	case "ECPCV_P95_IN_DOLLAR":
		fallthrough
	case "ECPC_IN_DOLLAR":
		fallthrough
	case "ECPC_IN_MICRO_DOLLAR":
		fallthrough
	case "ECPE_IN_DOLLAR":
		fallthrough
	case "ECPM_IN_MICRO_DOLLAR":
		fallthrough
	case "ECPV_IN_DOLLAR":
		fallthrough
	case "ECTR":
		fallthrough
	case "EENGAGEMENT_RATE":
		fallthrough
	case "ENGAGEMENT_1":
		fallthrough
	case "ENGAGEMENT_2":
		fallthrough
	case "ENGAGEMENT_RATE":
		fallthrough
	case "IDEA_PIN_PRODUCT_TAG_VISIT_1":
		fallthrough
	case "IDEA_PIN_PRODUCT_TAG_VISIT_2":
		fallthrough
	case "IMPRESSION_1":
		fallthrough
	case "IMPRESSION_1_GROSS":
		fallthrough
	case "IMPRESSION_2":
		fallthrough
	case "INAPP_CHECKOUT_COST_PER_ACTION":
		fallthrough
	case "OUTBOUND_CLICK_1":
		fallthrough
	case "OUTBOUND_CLICK_2":
		fallthrough
	case "PAGE_VISIT_COST_PER_ACTION":
		fallthrough
	case "PAGE_VISIT_ROAS":
		fallthrough
	case "PAID_IMPRESSION":
		fallthrough
	case "PIN_ID":
		fallthrough
	case "PIN_PROMOTION_ID":
		fallthrough
	case "REPIN_1":
		fallthrough
	case "REPIN_2":
		fallthrough
	case "REPIN_RATE":
		fallthrough
	case "SPEND_IN_DOLLAR":
		fallthrough
	case "SPEND_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CHECKOUT":
		fallthrough
	case "TOTAL_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CLICKTHROUGH":
		fallthrough
	case "TOTAL_CLICK_ADD_TO_CART":
		fallthrough
	case "TOTAL_CLICK_CHECKOUT":
		fallthrough
	case "TOTAL_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CLICK_LEAD":
		fallthrough
	case "TOTAL_CLICK_SIGNUP":
		fallthrough
	case "TOTAL_CLICK_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_CONVERSIONS":
		fallthrough
	case "TOTAL_CUSTOM":
		fallthrough
	case "TOTAL_ENGAGEMENT":
		fallthrough
	case "TOTAL_ENGAGEMENT_CHECKOUT":
		fallthrough
	case "TOTAL_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_ENGAGEMENT_LEAD":
		fallthrough
	case "TOTAL_ENGAGEMENT_SIGNUP":
		fallthrough
	case "TOTAL_ENGAGEMENT_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_IDEA_PIN_PRODUCT_TAG_VISIT":
		fallthrough
	case "TOTAL_IMPRESSION_FREQUENCY":
		fallthrough
	case "TOTAL_IMPRESSION_USER":
		fallthrough
	case "TOTAL_LEAD":
		fallthrough
	case "TOTAL_OFFLINE_CHECKOUT":
		fallthrough
	case "TOTAL_PAGE_VISIT":
		fallthrough
	case "TOTAL_REPIN_RATE":
		fallthrough
	case "TOTAL_SIGNUP":
		fallthrough
	case "TOTAL_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_VIDEO_3SEC_VIEWS":
		fallthrough
	case "TOTAL_VIDEO_AVG_WATCHTIME_IN_SECOND":
		fallthrough
	case "TOTAL_VIDEO_MRC_VIEWS":
		fallthrough
	case "TOTAL_VIDEO_P0_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P100_COMPLETE":
		fallthrough
	case "TOTAL_VIDEO_P25_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P50_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P75_COMBINED":
		fallthrough
	case "TOTAL_VIDEO_P95_COMBINED":
		fallthrough
	case "TOTAL_VIEW_ADD_TO_CART":
		fallthrough
	case "TOTAL_VIEW_CHECKOUT":
		fallthrough
	case "TOTAL_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_VIEW_LEAD":
		fallthrough
	case "TOTAL_VIEW_SIGNUP":
		fallthrough
	case "TOTAL_VIEW_SIGNUP_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_CLICK_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_CLICK_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_ENGAGEMENT_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_ENGAGEMENT_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "TOTAL_WEB_SESSIONS":
		fallthrough
	case "TOTAL_WEB_VIEW_CHECKOUT":
		fallthrough
	case "TOTAL_WEB_VIEW_CHECKOUT_VALUE_IN_MICRO_DOLLAR":
		fallthrough
	case "VIDEO_3SEC_VIEWS_2":
		fallthrough
	case "VIDEO_LENGTH":
		fallthrough
	case "VIDEO_MRC_VIEWS_2":
		fallthrough
	case "VIDEO_P0_COMBINED_2":
		fallthrough
	case "VIDEO_P100_COMPLETE_2":
		fallthrough
	case "VIDEO_P25_COMBINED_2":
		fallthrough
	case "VIDEO_P50_COMBINED_2":
		fallthrough
	case "VIDEO_P75_COMBINED_2":
		fallthrough
	case "VIDEO_P95_COMBINED_2":
		fallthrough
	case "WEB_CHECKOUT_COST_PER_ACTION":
		fallthrough
	case "WEB_CHECKOUT_ROAS":
		fallthrough
	case "WEB_SESSIONS_1":
		fallthrough
	case "WEB_SESSIONS_2":
		*e = SourcePinterestUpdateSchemasValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestUpdateSchemasValidEnums: %v", v)
	}
}

// ConversionReportTime - The date by which the conversion metrics returned from this endpoint will be reported. There are two dates associated with a conversion event: the date that the user interacted with the ad, and the date that the user completed a conversion event..
type ConversionReportTime string

const (
	ConversionReportTimeTimeOfAdAction   ConversionReportTime = "TIME_OF_AD_ACTION"
	ConversionReportTimeTimeOfConversion ConversionReportTime = "TIME_OF_CONVERSION"
)

func (e ConversionReportTime) ToPointer() *ConversionReportTime {
	return &e
}
func (e *ConversionReportTime) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIME_OF_AD_ACTION":
		fallthrough
	case "TIME_OF_CONVERSION":
		*e = ConversionReportTime(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConversionReportTime: %v", v)
	}
}

// EngagementWindowDays - Number of days to use as the conversion attribution window for an engagement action.
type EngagementWindowDays int64

const (
	EngagementWindowDaysZero     EngagementWindowDays = 0
	EngagementWindowDaysOne      EngagementWindowDays = 1
	EngagementWindowDaysSeven    EngagementWindowDays = 7
	EngagementWindowDaysFourteen EngagementWindowDays = 14
	EngagementWindowDaysThirty   EngagementWindowDays = 30
	EngagementWindowDaysSixty    EngagementWindowDays = 60
)

func (e EngagementWindowDays) ToPointer() *EngagementWindowDays {
	return &e
}
func (e *EngagementWindowDays) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 14:
		fallthrough
	case 30:
		fallthrough
	case 60:
		*e = EngagementWindowDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EngagementWindowDays: %v", v)
	}
}

// Granularity - Chosen granularity for API
type Granularity string

const (
	GranularityTotal Granularity = "TOTAL"
	GranularityDay   Granularity = "DAY"
	GranularityHour  Granularity = "HOUR"
	GranularityWeek  Granularity = "WEEK"
	GranularityMonth Granularity = "MONTH"
)

func (e Granularity) ToPointer() *Granularity {
	return &e
}
func (e *Granularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TOTAL":
		fallthrough
	case "DAY":
		fallthrough
	case "HOUR":
		fallthrough
	case "WEEK":
		fallthrough
	case "MONTH":
		*e = Granularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Granularity: %v", v)
	}
}

// SourcePinterestUpdateLevel - Chosen level for API
type SourcePinterestUpdateLevel string

const (
	SourcePinterestUpdateLevelAdvertiser            SourcePinterestUpdateLevel = "ADVERTISER"
	SourcePinterestUpdateLevelAdvertiserTargeting   SourcePinterestUpdateLevel = "ADVERTISER_TARGETING"
	SourcePinterestUpdateLevelCampaign              SourcePinterestUpdateLevel = "CAMPAIGN"
	SourcePinterestUpdateLevelCampaignTargeting     SourcePinterestUpdateLevel = "CAMPAIGN_TARGETING"
	SourcePinterestUpdateLevelAdGroup               SourcePinterestUpdateLevel = "AD_GROUP"
	SourcePinterestUpdateLevelAdGroupTargeting      SourcePinterestUpdateLevel = "AD_GROUP_TARGETING"
	SourcePinterestUpdateLevelPinPromotion          SourcePinterestUpdateLevel = "PIN_PROMOTION"
	SourcePinterestUpdateLevelPinPromotionTargeting SourcePinterestUpdateLevel = "PIN_PROMOTION_TARGETING"
	SourcePinterestUpdateLevelKeyword               SourcePinterestUpdateLevel = "KEYWORD"
	SourcePinterestUpdateLevelProductGroup          SourcePinterestUpdateLevel = "PRODUCT_GROUP"
	SourcePinterestUpdateLevelProductGroupTargeting SourcePinterestUpdateLevel = "PRODUCT_GROUP_TARGETING"
	SourcePinterestUpdateLevelProductItem           SourcePinterestUpdateLevel = "PRODUCT_ITEM"
)

func (e SourcePinterestUpdateLevel) ToPointer() *SourcePinterestUpdateLevel {
	return &e
}
func (e *SourcePinterestUpdateLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADVERTISER":
		fallthrough
	case "ADVERTISER_TARGETING":
		fallthrough
	case "CAMPAIGN":
		fallthrough
	case "CAMPAIGN_TARGETING":
		fallthrough
	case "AD_GROUP":
		fallthrough
	case "AD_GROUP_TARGETING":
		fallthrough
	case "PIN_PROMOTION":
		fallthrough
	case "PIN_PROMOTION_TARGETING":
		fallthrough
	case "KEYWORD":
		fallthrough
	case "PRODUCT_GROUP":
		fallthrough
	case "PRODUCT_GROUP_TARGETING":
		fallthrough
	case "PRODUCT_ITEM":
		*e = SourcePinterestUpdateLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestUpdateLevel: %v", v)
	}
}

// ViewWindowDays - Number of days to use as the conversion attribution window for a view action.
type ViewWindowDays int64

const (
	ViewWindowDaysZero     ViewWindowDays = 0
	ViewWindowDaysOne      ViewWindowDays = 1
	ViewWindowDaysSeven    ViewWindowDays = 7
	ViewWindowDaysFourteen ViewWindowDays = 14
	ViewWindowDaysThirty   ViewWindowDays = 30
	ViewWindowDaysSixty    ViewWindowDays = 60
)

func (e ViewWindowDays) ToPointer() *ViewWindowDays {
	return &e
}
func (e *ViewWindowDays) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 7:
		fallthrough
	case 14:
		fallthrough
	case 30:
		fallthrough
	case 60:
		*e = ViewWindowDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ViewWindowDays: %v", v)
	}
}

// ReportConfig - Config for custom report
type ReportConfig struct {
	// List of types of attribution for the conversion report
	AttributionTypes []SourcePinterestUpdateValidEnums `json:"attribution_types,omitempty"`
	// Number of days to use as the conversion attribution window for a pin click action.
	ClickWindowDays *ClickWindowDays `default:"30" json:"click_window_days"`
	// A list of chosen columns
	Columns []SourcePinterestUpdateSchemasValidEnums `json:"columns"`
	// The date by which the conversion metrics returned from this endpoint will be reported. There are two dates associated with a conversion event: the date that the user interacted with the ad, and the date that the user completed a conversion event..
	ConversionReportTime *ConversionReportTime `default:"TIME_OF_AD_ACTION" json:"conversion_report_time"`
	// Number of days to use as the conversion attribution window for an engagement action.
	EngagementWindowDays *EngagementWindowDays `default:"30" json:"engagement_window_days"`
	// Chosen granularity for API
	Granularity *Granularity `default:"TOTAL" json:"granularity"`
	// Chosen level for API
	Level *SourcePinterestUpdateLevel `default:"ADVERTISER" json:"level"`
	// The name value of report
	Name string `json:"name"`
	// A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by report api (913 days from today).
	StartDate *types.Date `json:"start_date,omitempty"`
	// Number of days to use as the conversion attribution window for a view action.
	ViewWindowDays *ViewWindowDays `default:"30" json:"view_window_days"`
}

func (r ReportConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReportConfig) GetAttributionTypes() []SourcePinterestUpdateValidEnums {
	if o == nil {
		return nil
	}
	return o.AttributionTypes
}

func (o *ReportConfig) GetClickWindowDays() *ClickWindowDays {
	if o == nil {
		return nil
	}
	return o.ClickWindowDays
}

func (o *ReportConfig) GetColumns() []SourcePinterestUpdateSchemasValidEnums {
	if o == nil {
		return []SourcePinterestUpdateSchemasValidEnums{}
	}
	return o.Columns
}

func (o *ReportConfig) GetConversionReportTime() *ConversionReportTime {
	if o == nil {
		return nil
	}
	return o.ConversionReportTime
}

func (o *ReportConfig) GetEngagementWindowDays() *EngagementWindowDays {
	if o == nil {
		return nil
	}
	return o.EngagementWindowDays
}

func (o *ReportConfig) GetGranularity() *Granularity {
	if o == nil {
		return nil
	}
	return o.Granularity
}

func (o *ReportConfig) GetLevel() *SourcePinterestUpdateLevel {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *ReportConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ReportConfig) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ReportConfig) GetViewWindowDays() *ViewWindowDays {
	if o == nil {
		return nil
	}
	return o.ViewWindowDays
}

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusPaused   Status = "PAUSED"
	StatusArchived Status = "ARCHIVED"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "PAUSED":
		fallthrough
	case "ARCHIVED":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type SourcePinterestUpdate struct {
	Credentials *OAuth20 `json:"credentials,omitempty"`
	// A list which contains ad statistics entries, each entry must have a name and can contains fields, breakdowns or action_breakdowns. Click on "add" to fill this field.
	CustomReports []ReportConfig `json:"custom_reports,omitempty"`
	// A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by api (89 days from today).
	StartDate *types.Date `json:"start_date,omitempty"`
	// For the ads, ad_groups, and campaigns streams, specifying a status will filter out records that do not match the specified ones. If a status is not specified, the source will default to records with a status of either ACTIVE or PAUSED.
	Status []Status `json:"status,omitempty"`
}

func (s SourcePinterestUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePinterestUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePinterestUpdate) GetCredentials() *OAuth20 {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourcePinterestUpdate) GetCustomReports() []ReportConfig {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourcePinterestUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourcePinterestUpdate) GetStatus() []Status {
	if o == nil {
		return nil
	}
	return o.Status
}
