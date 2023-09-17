package params

import "net/url"

// The BibleBookParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/books/{bookId}` end point
type BibleBookParams struct {
	IncludeChapters bool
}

func (params *BibleBookParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.IncludeChapters {
		values.Add("include-chapters", "true")
	}
	return values
}
