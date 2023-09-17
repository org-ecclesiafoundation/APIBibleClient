package params

import "net/url"

// The AudioBibleBooksParams struct contains all optional parameters
// for the API call to the `bibles/{audioBibleId}/books` end point
type AudioBibleBooksParams struct {
	IncludeChapters            bool
	IncludeChaptersAndSections bool
}

func (params *AudioBibleBooksParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.IncludeChapters {
		values.Add("include-chapters", "true")
	}
	if params.IncludeChaptersAndSections {
		values.Add("include-chapters", "true")
	}
	return values
}
