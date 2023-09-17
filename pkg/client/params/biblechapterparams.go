package params

import (
	"net/url"
	"strings"
)

// The BibleChapterParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/chapters/{chapterId}` end point
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
