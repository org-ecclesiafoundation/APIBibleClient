package params

import "net/url"

// The AudioBibleBookParams struct contains all optional parameters
// for the API call to the `audio-bibles/{audioBibleId}/books/{bookId}` end point
type AudioBibleBookParams struct {
	IncludeChapters bool
}

func (params *AudioBibleBookParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.IncludeChapters {
		values.Add("include-chapters", "true")
	}
	return values
}
