package params

import "net/url"

// The AudioBiblesParams struct contains all optional parameters
// for the API call to the `audio-bibles` end point.
type AudioBiblesParams struct {
	Language           string
	Abbreviation       string
	Name               string
	RelatedTextBibleId string
	IncludeFullDetails bool
}

func (params *AudioBiblesParams) ProduceQueryParameters() url.Values {
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
	if params.RelatedTextBibleId != "" {
		values.Add("bibleId", params.RelatedTextBibleId)
	}
	if params.IncludeFullDetails {
		values.Add("include-full-details", "true")
	}
	return values
}
