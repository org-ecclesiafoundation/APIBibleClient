package params

import "net/url"

// The BibleBooksParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/books` end point
type BibleBooksParams struct {
	IncludeChapters            bool
	IncludeChaptersAndSections bool
}

func (params *BibleBooksParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.IncludeChapters {
		values.Add("include-chapters", "true")
	}
	if params.IncludeChaptersAndSections {
		values.Add("include-chapters", "true")
	}
	return values
}
