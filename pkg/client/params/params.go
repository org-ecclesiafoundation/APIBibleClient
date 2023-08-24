package params

import (
	"net/url"
	"strings"
)

// The BiblesParams struct contains all optional parameters
// for the API call to the `bibles` end point.
type BiblesParams struct {
	Language           string
	Abbreviation       string
	Name               string
	Ids                []string
	IncludeFullDetails bool
}

func (params *BiblesParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.Language != "" {
		values.Add("language", params.Language)
	}
	if params.Abbreviation != "" {
		values.Add("abbreviation", params.Abbreviation)
	}
	if params.Name != "" {
		values.Add("name", params.Name)
	}
	if params.Ids != nil {
		values.Add("ids", strings.Join(params.Ids, ","))
	}
	if params.IncludeFullDetails == true {
		values.Add("include-full-details", "true")
	}
	return values
}
