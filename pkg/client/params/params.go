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

// The BibleBooksParams struct contains all optional parameters
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

// The AudioBibleBookParams struct contains all optional parameters
// for the API call to the `audio-bibles/{audioBibleId}/books/{bookId}` end point
type BibleChapterParams struct {
	ContentType           string
	IncludeNotes          bool
	IncludeTitles         bool
	IncludeChapterNumbers bool
	IncludeVerseNumbers   bool
	IncludeVerseSpans     bool
	Parallels             []string
}

func (params *BibleChapterParams) ProduceQueryParameters() url.Values {
	values := url.Values{}

	if params.ContentType != "" {
		values.Add("content-type", params.ContentType)
	}
	if params.IncludeNotes {
		values.Add("include-notes", "true")
	}
	if params.IncludeTitles {
		values.Add("include-titles", "true")
	}
	if params.IncludeChapterNumbers {
		values.Add("include-chapter-numbers", "true")
	}
	if params.IncludeVerseNumbers {
		values.Add("include-verse-numbers", "true")
	}
	if params.IncludeVerseSpans {
		values.Add("include-verse-spans", "true")
	}
	if params.Parallels != nil {
		values.Add("parallels", strings.Join(params.Parallels, ","))
	}
	return values
}
