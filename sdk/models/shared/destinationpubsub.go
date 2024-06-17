// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type Pubsub string

const (
	PubsubPubsub Pubsub = "pubsub"
)

func (e Pubsub) ToPointer() *Pubsub {
	return &e
}
func (e *Pubsub) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pubsub":
		*e = Pubsub(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pubsub: %v", v)
	}
}

type DestinationPubsub struct {
	// Number of ms before the buffer is flushed
	BatchingDelayThreshold *int64 `default:"1" json:"batching_delay_threshold"`
	// Number of messages before the buffer is flushed
	BatchingElementCountThreshold *int64 `default:"1" json:"batching_element_count_threshold"`
	// If TRUE messages will be buffered instead of sending them one by one
	BatchingEnabled *bool `default:"false" json:"batching_enabled"`
	// Number of bytes before the buffer is flushed
	BatchingRequestBytesThreshold *int64 `default:"1" json:"batching_request_bytes_threshold"`
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/pubsub">docs</a> if you need help generating this key.
	CredentialsJSON string `json:"credentials_json"`
	destinationType Pubsub `const:"pubsub" json:"destinationType"`
	// If TRUE PubSub publisher will have <a href="https://cloud.google.com/pubsub/docs/ordering">message ordering</a> enabled. Every message will have an ordering key of stream
	OrderingEnabled *bool `default:"false" json:"ordering_enabled"`
	// The GCP project ID for the project containing the target PubSub.
	ProjectID string `json:"project_id"`
	// The PubSub topic ID in the given GCP project ID.
	TopicID string `json:"topic_id"`
}

func (d DestinationPubsub) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPubsub) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPubsub) GetBatchingDelayThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchingDelayThreshold
}

func (o *DestinationPubsub) GetBatchingElementCountThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchingElementCountThreshold
}

func (o *DestinationPubsub) GetBatchingEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.BatchingEnabled
}

func (o *DestinationPubsub) GetBatchingRequestBytesThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchingRequestBytesThreshold
}

func (o *DestinationPubsub) GetCredentialsJSON() string {
	if o == nil {
		return ""
	}
	return o.CredentialsJSON
}

func (o *DestinationPubsub) GetDestinationType() Pubsub {
	return PubsubPubsub
}

func (o *DestinationPubsub) GetOrderingEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.OrderingEnabled
}

func (o *DestinationPubsub) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *DestinationPubsub) GetTopicID() string {
	if o == nil {
		return ""
	}
	return o.TopicID
}
