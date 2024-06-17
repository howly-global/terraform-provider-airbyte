// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

// DestinationBigqueryDatasetLocation - The location of the dataset. Warning: Changes made after creation will not be applied. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
type DestinationBigqueryDatasetLocation string

const (
	DestinationBigqueryDatasetLocationUs                     DestinationBigqueryDatasetLocation = "US"
	DestinationBigqueryDatasetLocationEu                     DestinationBigqueryDatasetLocation = "EU"
	DestinationBigqueryDatasetLocationAsiaEast1              DestinationBigqueryDatasetLocation = "asia-east1"
	DestinationBigqueryDatasetLocationAsiaEast2              DestinationBigqueryDatasetLocation = "asia-east2"
	DestinationBigqueryDatasetLocationAsiaNortheast1         DestinationBigqueryDatasetLocation = "asia-northeast1"
	DestinationBigqueryDatasetLocationAsiaNortheast2         DestinationBigqueryDatasetLocation = "asia-northeast2"
	DestinationBigqueryDatasetLocationAsiaNortheast3         DestinationBigqueryDatasetLocation = "asia-northeast3"
	DestinationBigqueryDatasetLocationAsiaSouth1             DestinationBigqueryDatasetLocation = "asia-south1"
	DestinationBigqueryDatasetLocationAsiaSouth2             DestinationBigqueryDatasetLocation = "asia-south2"
	DestinationBigqueryDatasetLocationAsiaSoutheast1         DestinationBigqueryDatasetLocation = "asia-southeast1"
	DestinationBigqueryDatasetLocationAsiaSoutheast2         DestinationBigqueryDatasetLocation = "asia-southeast2"
	DestinationBigqueryDatasetLocationAustraliaSoutheast1    DestinationBigqueryDatasetLocation = "australia-southeast1"
	DestinationBigqueryDatasetLocationAustraliaSoutheast2    DestinationBigqueryDatasetLocation = "australia-southeast2"
	DestinationBigqueryDatasetLocationEuropeCentral1         DestinationBigqueryDatasetLocation = "europe-central1"
	DestinationBigqueryDatasetLocationEuropeCentral2         DestinationBigqueryDatasetLocation = "europe-central2"
	DestinationBigqueryDatasetLocationEuropeNorth1           DestinationBigqueryDatasetLocation = "europe-north1"
	DestinationBigqueryDatasetLocationEuropeSouthwest1       DestinationBigqueryDatasetLocation = "europe-southwest1"
	DestinationBigqueryDatasetLocationEuropeWest1            DestinationBigqueryDatasetLocation = "europe-west1"
	DestinationBigqueryDatasetLocationEuropeWest2            DestinationBigqueryDatasetLocation = "europe-west2"
	DestinationBigqueryDatasetLocationEuropeWest3            DestinationBigqueryDatasetLocation = "europe-west3"
	DestinationBigqueryDatasetLocationEuropeWest4            DestinationBigqueryDatasetLocation = "europe-west4"
	DestinationBigqueryDatasetLocationEuropeWest6            DestinationBigqueryDatasetLocation = "europe-west6"
	DestinationBigqueryDatasetLocationEuropeWest7            DestinationBigqueryDatasetLocation = "europe-west7"
	DestinationBigqueryDatasetLocationEuropeWest8            DestinationBigqueryDatasetLocation = "europe-west8"
	DestinationBigqueryDatasetLocationEuropeWest9            DestinationBigqueryDatasetLocation = "europe-west9"
	DestinationBigqueryDatasetLocationEuropeWest12           DestinationBigqueryDatasetLocation = "europe-west12"
	DestinationBigqueryDatasetLocationMeCentral1             DestinationBigqueryDatasetLocation = "me-central1"
	DestinationBigqueryDatasetLocationMeCentral2             DestinationBigqueryDatasetLocation = "me-central2"
	DestinationBigqueryDatasetLocationMeWest1                DestinationBigqueryDatasetLocation = "me-west1"
	DestinationBigqueryDatasetLocationNorthamericaNortheast1 DestinationBigqueryDatasetLocation = "northamerica-northeast1"
	DestinationBigqueryDatasetLocationNorthamericaNortheast2 DestinationBigqueryDatasetLocation = "northamerica-northeast2"
	DestinationBigqueryDatasetLocationSouthamericaEast1      DestinationBigqueryDatasetLocation = "southamerica-east1"
	DestinationBigqueryDatasetLocationSouthamericaWest1      DestinationBigqueryDatasetLocation = "southamerica-west1"
	DestinationBigqueryDatasetLocationUsCentral1             DestinationBigqueryDatasetLocation = "us-central1"
	DestinationBigqueryDatasetLocationUsEast1                DestinationBigqueryDatasetLocation = "us-east1"
	DestinationBigqueryDatasetLocationUsEast2                DestinationBigqueryDatasetLocation = "us-east2"
	DestinationBigqueryDatasetLocationUsEast3                DestinationBigqueryDatasetLocation = "us-east3"
	DestinationBigqueryDatasetLocationUsEast4                DestinationBigqueryDatasetLocation = "us-east4"
	DestinationBigqueryDatasetLocationUsEast5                DestinationBigqueryDatasetLocation = "us-east5"
	DestinationBigqueryDatasetLocationUsSouth1               DestinationBigqueryDatasetLocation = "us-south1"
	DestinationBigqueryDatasetLocationUsWest1                DestinationBigqueryDatasetLocation = "us-west1"
	DestinationBigqueryDatasetLocationUsWest2                DestinationBigqueryDatasetLocation = "us-west2"
	DestinationBigqueryDatasetLocationUsWest3                DestinationBigqueryDatasetLocation = "us-west3"
	DestinationBigqueryDatasetLocationUsWest4                DestinationBigqueryDatasetLocation = "us-west4"
)

func (e DestinationBigqueryDatasetLocation) ToPointer() *DestinationBigqueryDatasetLocation {
	return &e
}
func (e *DestinationBigqueryDatasetLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		fallthrough
	case "asia-east1":
		fallthrough
	case "asia-east2":
		fallthrough
	case "asia-northeast1":
		fallthrough
	case "asia-northeast2":
		fallthrough
	case "asia-northeast3":
		fallthrough
	case "asia-south1":
		fallthrough
	case "asia-south2":
		fallthrough
	case "asia-southeast1":
		fallthrough
	case "asia-southeast2":
		fallthrough
	case "australia-southeast1":
		fallthrough
	case "australia-southeast2":
		fallthrough
	case "europe-central1":
		fallthrough
	case "europe-central2":
		fallthrough
	case "europe-north1":
		fallthrough
	case "europe-southwest1":
		fallthrough
	case "europe-west1":
		fallthrough
	case "europe-west2":
		fallthrough
	case "europe-west3":
		fallthrough
	case "europe-west4":
		fallthrough
	case "europe-west6":
		fallthrough
	case "europe-west7":
		fallthrough
	case "europe-west8":
		fallthrough
	case "europe-west9":
		fallthrough
	case "europe-west12":
		fallthrough
	case "me-central1":
		fallthrough
	case "me-central2":
		fallthrough
	case "me-west1":
		fallthrough
	case "northamerica-northeast1":
		fallthrough
	case "northamerica-northeast2":
		fallthrough
	case "southamerica-east1":
		fallthrough
	case "southamerica-west1":
		fallthrough
	case "us-central1":
		fallthrough
	case "us-east1":
		fallthrough
	case "us-east2":
		fallthrough
	case "us-east3":
		fallthrough
	case "us-east4":
		fallthrough
	case "us-east5":
		fallthrough
	case "us-south1":
		fallthrough
	case "us-west1":
		fallthrough
	case "us-west2":
		fallthrough
	case "us-west3":
		fallthrough
	case "us-west4":
		*e = DestinationBigqueryDatasetLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDatasetLocation: %v", v)
	}
}

type Bigquery string

const (
	BigqueryBigquery Bigquery = "bigquery"
)

func (e Bigquery) ToPointer() *Bigquery {
	return &e
}
func (e *Bigquery) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bigquery":
		*e = Bigquery(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Bigquery: %v", v)
	}
}

type DestinationBigquerySchemasMethod string

const (
	DestinationBigquerySchemasMethodStandard DestinationBigquerySchemasMethod = "Standard"
)

func (e DestinationBigquerySchemasMethod) ToPointer() *DestinationBigquerySchemasMethod {
	return &e
}
func (e *DestinationBigquerySchemasMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		*e = DestinationBigquerySchemasMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigquerySchemasMethod: %v", v)
	}
}

// DestinationBigqueryStandardInserts - <i>(not recommended)</i> Direct loading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In all other cases, you should use GCS staging.
type DestinationBigqueryStandardInserts struct {
	method DestinationBigquerySchemasMethod `const:"Standard" json:"method"`
}

func (d DestinationBigqueryStandardInserts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryStandardInserts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryStandardInserts) GetMethod() DestinationBigquerySchemasMethod {
	return DestinationBigquerySchemasMethodStandard
}

type DestinationBigqueryCredentialType string

const (
	DestinationBigqueryCredentialTypeHmacKey DestinationBigqueryCredentialType = "HMAC_KEY"
)

func (e DestinationBigqueryCredentialType) ToPointer() *DestinationBigqueryCredentialType {
	return &e
}
func (e *DestinationBigqueryCredentialType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HMAC_KEY":
		*e = DestinationBigqueryCredentialType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryCredentialType: %v", v)
	}
}

type DestinationBigqueryHMACKey struct {
	credentialType DestinationBigqueryCredentialType `const:"HMAC_KEY" json:"credential_type"`
	// HMAC key access ID. When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long.
	HmacKeyAccessID string `json:"hmac_key_access_id"`
	// The corresponding secret for the access ID. It is a 40-character base-64 encoded string.
	HmacKeySecret string `json:"hmac_key_secret"`
}

func (d DestinationBigqueryHMACKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryHMACKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryHMACKey) GetCredentialType() DestinationBigqueryCredentialType {
	return DestinationBigqueryCredentialTypeHmacKey
}

func (o *DestinationBigqueryHMACKey) GetHmacKeyAccessID() string {
	if o == nil {
		return ""
	}
	return o.HmacKeyAccessID
}

func (o *DestinationBigqueryHMACKey) GetHmacKeySecret() string {
	if o == nil {
		return ""
	}
	return o.HmacKeySecret
}

type DestinationBigqueryCredentialUnionType string

const (
	DestinationBigqueryCredentialUnionTypeDestinationBigqueryHMACKey DestinationBigqueryCredentialUnionType = "destination-bigquery_HMAC key"
)

// DestinationBigqueryCredential - An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.
type DestinationBigqueryCredential struct {
	DestinationBigqueryHMACKey *DestinationBigqueryHMACKey

	Type DestinationBigqueryCredentialUnionType
}

func CreateDestinationBigqueryCredentialDestinationBigqueryHMACKey(destinationBigqueryHMACKey DestinationBigqueryHMACKey) DestinationBigqueryCredential {
	typ := DestinationBigqueryCredentialUnionTypeDestinationBigqueryHMACKey

	return DestinationBigqueryCredential{
		DestinationBigqueryHMACKey: &destinationBigqueryHMACKey,
		Type:                       typ,
	}
}

func (u *DestinationBigqueryCredential) UnmarshalJSON(data []byte) error {

	var destinationBigqueryHMACKey DestinationBigqueryHMACKey = DestinationBigqueryHMACKey{}
	if err := utils.UnmarshalJSON(data, &destinationBigqueryHMACKey, "", true, true); err == nil {
		u.DestinationBigqueryHMACKey = &destinationBigqueryHMACKey
		u.Type = DestinationBigqueryCredentialUnionTypeDestinationBigqueryHMACKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationBigqueryCredential) MarshalJSON() ([]byte, error) {
	if u.DestinationBigqueryHMACKey != nil {
		return utils.MarshalJSON(u.DestinationBigqueryHMACKey, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationBigqueryGCSTmpFilesAfterwardProcessing - This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.
type DestinationBigqueryGCSTmpFilesAfterwardProcessing string

const (
	DestinationBigqueryGCSTmpFilesAfterwardProcessingDeleteAllTmpFilesFromGcs DestinationBigqueryGCSTmpFilesAfterwardProcessing = "Delete all tmp files from GCS"
	DestinationBigqueryGCSTmpFilesAfterwardProcessingKeepAllTmpFilesInGcs     DestinationBigqueryGCSTmpFilesAfterwardProcessing = "Keep all tmp files in GCS"
)

func (e DestinationBigqueryGCSTmpFilesAfterwardProcessing) ToPointer() *DestinationBigqueryGCSTmpFilesAfterwardProcessing {
	return &e
}
func (e *DestinationBigqueryGCSTmpFilesAfterwardProcessing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Delete all tmp files from GCS":
		fallthrough
	case "Keep all tmp files in GCS":
		*e = DestinationBigqueryGCSTmpFilesAfterwardProcessing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryGCSTmpFilesAfterwardProcessing: %v", v)
	}
}

type DestinationBigqueryMethod string

const (
	DestinationBigqueryMethodGcsStaging DestinationBigqueryMethod = "GCS Staging"
)

func (e DestinationBigqueryMethod) ToPointer() *DestinationBigqueryMethod {
	return &e
}
func (e *DestinationBigqueryMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS Staging":
		*e = DestinationBigqueryMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryMethod: %v", v)
	}
}

// DestinationBigqueryGCSStaging - <i>(recommended)</i> Writes large batches of records to a file, uploads the file to GCS, then uses COPY INTO to load your data into BigQuery. Provides best-in-class speed, reliability and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
type DestinationBigqueryGCSStaging struct {
	// An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.
	Credential DestinationBigqueryCredential `json:"credential"`
	// The name of the GCS bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.
	GcsBucketName string `json:"gcs_bucket_name"`
	// Directory under the GCS bucket where data will be written.
	GcsBucketPath string `json:"gcs_bucket_path"`
	// This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.
	KeepFilesInGcsBucket *DestinationBigqueryGCSTmpFilesAfterwardProcessing `default:"Delete all tmp files from GCS" json:"keep_files_in_gcs-bucket"`
	method               DestinationBigqueryMethod                          `const:"GCS Staging" json:"method"`
}

func (d DestinationBigqueryGCSStaging) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryGCSStaging) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryGCSStaging) GetCredential() DestinationBigqueryCredential {
	if o == nil {
		return DestinationBigqueryCredential{}
	}
	return o.Credential
}

func (o *DestinationBigqueryGCSStaging) GetGcsBucketName() string {
	if o == nil {
		return ""
	}
	return o.GcsBucketName
}

func (o *DestinationBigqueryGCSStaging) GetGcsBucketPath() string {
	if o == nil {
		return ""
	}
	return o.GcsBucketPath
}

func (o *DestinationBigqueryGCSStaging) GetKeepFilesInGcsBucket() *DestinationBigqueryGCSTmpFilesAfterwardProcessing {
	if o == nil {
		return nil
	}
	return o.KeepFilesInGcsBucket
}

func (o *DestinationBigqueryGCSStaging) GetMethod() DestinationBigqueryMethod {
	return DestinationBigqueryMethodGcsStaging
}

type DestinationBigqueryLoadingMethodType string

const (
	DestinationBigqueryLoadingMethodTypeDestinationBigqueryGCSStaging      DestinationBigqueryLoadingMethodType = "destination-bigquery_GCS Staging"
	DestinationBigqueryLoadingMethodTypeDestinationBigqueryStandardInserts DestinationBigqueryLoadingMethodType = "destination-bigquery_Standard Inserts"
)

// DestinationBigqueryLoadingMethod - The way data will be uploaded to BigQuery.
type DestinationBigqueryLoadingMethod struct {
	DestinationBigqueryGCSStaging      *DestinationBigqueryGCSStaging
	DestinationBigqueryStandardInserts *DestinationBigqueryStandardInserts

	Type DestinationBigqueryLoadingMethodType
}

func CreateDestinationBigqueryLoadingMethodDestinationBigqueryGCSStaging(destinationBigqueryGCSStaging DestinationBigqueryGCSStaging) DestinationBigqueryLoadingMethod {
	typ := DestinationBigqueryLoadingMethodTypeDestinationBigqueryGCSStaging

	return DestinationBigqueryLoadingMethod{
		DestinationBigqueryGCSStaging: &destinationBigqueryGCSStaging,
		Type:                          typ,
	}
}

func CreateDestinationBigqueryLoadingMethodDestinationBigqueryStandardInserts(destinationBigqueryStandardInserts DestinationBigqueryStandardInserts) DestinationBigqueryLoadingMethod {
	typ := DestinationBigqueryLoadingMethodTypeDestinationBigqueryStandardInserts

	return DestinationBigqueryLoadingMethod{
		DestinationBigqueryStandardInserts: &destinationBigqueryStandardInserts,
		Type:                               typ,
	}
}

func (u *DestinationBigqueryLoadingMethod) UnmarshalJSON(data []byte) error {

	var destinationBigqueryStandardInserts DestinationBigqueryStandardInserts = DestinationBigqueryStandardInserts{}
	if err := utils.UnmarshalJSON(data, &destinationBigqueryStandardInserts, "", true, true); err == nil {
		u.DestinationBigqueryStandardInserts = &destinationBigqueryStandardInserts
		u.Type = DestinationBigqueryLoadingMethodTypeDestinationBigqueryStandardInserts
		return nil
	}

	var destinationBigqueryGCSStaging DestinationBigqueryGCSStaging = DestinationBigqueryGCSStaging{}
	if err := utils.UnmarshalJSON(data, &destinationBigqueryGCSStaging, "", true, true); err == nil {
		u.DestinationBigqueryGCSStaging = &destinationBigqueryGCSStaging
		u.Type = DestinationBigqueryLoadingMethodTypeDestinationBigqueryGCSStaging
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationBigqueryLoadingMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationBigqueryGCSStaging != nil {
		return utils.MarshalJSON(u.DestinationBigqueryGCSStaging, "", true)
	}

	if u.DestinationBigqueryStandardInserts != nil {
		return utils.MarshalJSON(u.DestinationBigqueryStandardInserts, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationBigqueryTransformationQueryRunType - Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type <a href="https://cloud.google.com/bigquery/docs/running-queries#queries">here</a>. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries <a href="https://cloud.google.com/bigquery/docs/running-queries#batch">here</a>. The default "interactive" value is used if not set explicitly.
type DestinationBigqueryTransformationQueryRunType string

const (
	DestinationBigqueryTransformationQueryRunTypeInteractive DestinationBigqueryTransformationQueryRunType = "interactive"
	DestinationBigqueryTransformationQueryRunTypeBatch       DestinationBigqueryTransformationQueryRunType = "batch"
)

func (e DestinationBigqueryTransformationQueryRunType) ToPointer() *DestinationBigqueryTransformationQueryRunType {
	return &e
}
func (e *DestinationBigqueryTransformationQueryRunType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "interactive":
		fallthrough
	case "batch":
		*e = DestinationBigqueryTransformationQueryRunType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryTransformationQueryRunType: %v", v)
	}
}

type DestinationBigquery struct {
	// Google BigQuery client's chunk (buffer) size (MIN=1, MAX = 15) for each table. The size that will be written by a single RPC. Written data will be buffered and only flushed upon reaching this size or closing the channel. The default 15MB value is used if not set explicitly. Read more <a href="https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.client.Client.html">here</a>.
	BigQueryClientBufferSizeMb *int64 `default:"15" json:"big_query_client_buffer_size_mb"`
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/bigquery#service-account-key">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.
	CredentialsJSON *string `json:"credentials_json,omitempty"`
	// The default BigQuery Dataset ID that tables are replicated to if the source does not specify a namespace. Read more <a href="https://cloud.google.com/bigquery/docs/datasets#create-dataset">here</a>.
	DatasetID string `json:"dataset_id"`
	// The location of the dataset. Warning: Changes made after creation will not be applied. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
	DatasetLocation DestinationBigqueryDatasetLocation `json:"dataset_location"`
	destinationType Bigquery                           `const:"bigquery" json:"destinationType"`
	// Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions
	DisableTypeDedupe *bool `default:"false" json:"disable_type_dedupe"`
	// The way data will be uploaded to BigQuery.
	LoadingMethod *DestinationBigqueryLoadingMethod `json:"loading_method,omitempty"`
	// The GCP project ID for the project containing the target BigQuery dataset. Read more <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">here</a>.
	ProjectID string `json:"project_id"`
	// The dataset to write raw tables into (default: airbyte_internal)
	RawDataDataset *string `json:"raw_data_dataset,omitempty"`
	// Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type <a href="https://cloud.google.com/bigquery/docs/running-queries#queries">here</a>. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries <a href="https://cloud.google.com/bigquery/docs/running-queries#batch">here</a>. The default "interactive" value is used if not set explicitly.
	TransformationPriority *DestinationBigqueryTransformationQueryRunType `default:"interactive" json:"transformation_priority"`
}

func (d DestinationBigquery) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigquery) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigquery) GetBigQueryClientBufferSizeMb() *int64 {
	if o == nil {
		return nil
	}
	return o.BigQueryClientBufferSizeMb
}

func (o *DestinationBigquery) GetCredentialsJSON() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsJSON
}

func (o *DestinationBigquery) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

func (o *DestinationBigquery) GetDatasetLocation() DestinationBigqueryDatasetLocation {
	if o == nil {
		return DestinationBigqueryDatasetLocation("")
	}
	return o.DatasetLocation
}

func (o *DestinationBigquery) GetDestinationType() Bigquery {
	return BigqueryBigquery
}

func (o *DestinationBigquery) GetDisableTypeDedupe() *bool {
	if o == nil {
		return nil
	}
	return o.DisableTypeDedupe
}

func (o *DestinationBigquery) GetLoadingMethod() *DestinationBigqueryLoadingMethod {
	if o == nil {
		return nil
	}
	return o.LoadingMethod
}

func (o *DestinationBigquery) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *DestinationBigquery) GetRawDataDataset() *string {
	if o == nil {
		return nil
	}
	return o.RawDataDataset
}

func (o *DestinationBigquery) GetTransformationPriority() *DestinationBigqueryTransformationQueryRunType {
	if o == nil {
		return nil
	}
	return o.TransformationPriority
}
