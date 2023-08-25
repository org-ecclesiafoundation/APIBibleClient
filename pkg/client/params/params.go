package params

import (
	"net/url"
	"strings"
)

// The CanProduceQueryParams interface exists
// to help with producing URLs with query parameters
type CanProduceQueryParams interface {
	ProduceQueryParameters() url.Values
}

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
	if params.IncludeFullDetails {
		values.Add("include-full-details", "true")
	}
	return values
}

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

// The BibleBooksParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/books` end point
type BibleBooksParams struct {
	includeChapters            bool
	includeChaptersAndSections bool
}

func (params *BibleBooksParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.includeChapters {
		values.Add("include-chapters", "true")
	}
	if params.includeChaptersAndSections {
		values.Add("include-chapters", "true")
	}
	return values
}

// The BibleBookParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/books/{bookId}` end point
type BibleBookParams struct {
	includeChapters bool
}

func (params *BibleBookParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.includeChapters {
		values.Add("include-chapters", "true")
	}
	return values
}

// The BibleBooksParams struct contains all optional parameters
// for the API call to the `bibles/{audioBibleId}/books` end point
type AudioBibleBooksParams struct {
	includeChapters            bool
	includeChaptersAndSections bool
}

func (params *AudioBibleBooksParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.includeChapters {
		values.Add("include-chapters", "true")
	}
	if params.includeChaptersAndSections {
		values.Add("include-chapters", "true")
	}
	return values
}

// The AudioBibleBookParams struct contains all optional parameters
// for the API call to the `audio-bibles/{audioBibleId}/books/{bookId}` end point
type AudioBibleBookParams struct {
	includeChapters bool
}

func (params *AudioBibleBookParams) ProduceQueryParameters() url.Values {
	values := url.Values{}
	if params.includeChapters {
		values.Add("include-chapters", "true")
	}
	return values
}
