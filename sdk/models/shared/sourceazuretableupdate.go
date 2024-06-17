// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type SourceAzureTableUpdate struct {
	// Azure Table Storage Access Key. See the <a href="https://docs.airbyte.com/integrations/sources/azure-table">docs</a> for more information on how to obtain this key.
	StorageAccessKey string `json:"storage_access_key"`
	// The name of your storage account.
	StorageAccountName string `json:"storage_account_name"`
	// Azure Table Storage service account URL suffix. See the <a href="https://docs.airbyte.com/integrations/sources/azure-table">docs</a> for more information on how to obtain endpoint suffix
	StorageEndpointSuffix *string `default:"core.windows.net" json:"storage_endpoint_suffix"`
}

func (s SourceAzureTableUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAzureTableUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAzureTableUpdate) GetStorageAccessKey() string {
	if o == nil {
		return ""
	}
	return o.StorageAccessKey
}

func (o *SourceAzureTableUpdate) GetStorageAccountName() string {
	if o == nil {
		return ""
	}
	return o.StorageAccountName
}

func (o *SourceAzureTableUpdate) GetStorageEndpointSuffix() *string {
	if o == nil {
		return nil
	}
	return o.StorageEndpointSuffix
}
