// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/howly-global/terraform-provider-airbyte/sdk/internal/utils"
)

type WikipediaPageviews string

const (
	WikipediaPageviewsWikipediaPageviews WikipediaPageviews = "wikipedia-pageviews"
)

func (e WikipediaPageviews) ToPointer() *WikipediaPageviews {
	return &e
}
func (e *WikipediaPageviews) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "wikipedia-pageviews":
		*e = WikipediaPageviews(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WikipediaPageviews: %v", v)
	}
}

type SourceWikipediaPageviews struct {
	// If you want to filter by access method, use one of desktop, mobile-app or mobile-web. If you are interested in pageviews regardless of access method, use all-access.
	Access string `json:"access"`
	// If you want to filter by agent type, use one of user, automated or spider. If you are interested in pageviews regardless of agent type, use all-agents.
	Agent string `json:"agent"`
	// The title of any article in the specified project. Any spaces should be replaced with underscores. It also should be URI-encoded, so that non-URI-safe characters like %, / or ? are accepted.
	Article string `json:"article"`
	// The ISO 3166-1 alpha-2 code of a country for which to retrieve top articles.
	Country string `json:"country"`
	// The date of the last day to include, in YYYYMMDD or YYYYMMDDHH format.
	End string `json:"end"`
	// If you want to filter by project, use the domain of any Wikimedia project.
	Project    string             `json:"project"`
	sourceType WikipediaPageviews `const:"wikipedia-pageviews" json:"sourceType"`
	// The date of the first day to include, in YYYYMMDD or YYYYMMDDHH format.
	Start string `json:"start"`
}

func (s SourceWikipediaPageviews) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceWikipediaPageviews) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceWikipediaPageviews) GetAccess() string {
	if o == nil {
		return ""
	}
	return o.Access
}

func (o *SourceWikipediaPageviews) GetAgent() string {
	if o == nil {
		return ""
	}
	return o.Agent
}

func (o *SourceWikipediaPageviews) GetArticle() string {
	if o == nil {
		return ""
	}
	return o.Article
}

func (o *SourceWikipediaPageviews) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *SourceWikipediaPageviews) GetEnd() string {
	if o == nil {
		return ""
	}
	return o.End
}

func (o *SourceWikipediaPageviews) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

func (o *SourceWikipediaPageviews) GetSourceType() WikipediaPageviews {
	return WikipediaPageviewsWikipediaPageviews
}

func (o *SourceWikipediaPageviews) GetStart() string {
	if o == nil {
		return ""
	}
	return o.Start
}
