//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package params

import (
	"net/url"
	"strconv"
	"strings"
)

// The CanProduceQueryParams interface exists
// to help with producing URLs with query parameters
type CanProduceQueryParams interface {
	ProduceQueryParameters() url.Values
}

// The BiblePassageParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/passages/{passageId}` end point
type BiblePassageParams struct {
	ContentType           string
	IncludeNotes          bool
	IncludeTitles         bool
	IncludeChapterNumbers bool
	IncludeVerseNumbers   bool
	IncludeVerseSpans     bool
	Parallels             []string
	UseOrgId              bool
}

func (params *BiblePassageParams) ProduceQueryParameters() url.Values {
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
	if params.UseOrgId {
		values.Add("use-org-id", "true")
	}
	return values
}

// The BibleVerseParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/verses/{verseId}` end point
type BibleVerseParams struct {
	ContentType           string
	IncludeNotes          bool
	IncludeTitles         bool
	IncludeChapterNumbers bool
	IncludeVerseNumbers   bool
	IncludeVerseSpans     bool
	Parallels             []string
	UseOrgId              bool
}

func (params *BibleVerseParams) ProduceQueryParameters() url.Values {
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
	if params.UseOrgId {
		values.Add("use-org-id", "true")
	}
	return values
}

// The BibleSearchParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/search` end point
type BibleSearchParams struct {
	Query     string
	Limit     int
	Offset    int
	Sort      string
	Range     string
	Fuzziness string
}

func (params *BibleSearchParams) ProduceQueryParameters() url.Values {
	values := url.Values{}

	if params.Query != "" {
		values.Add("query", params.Query)
	}
	if params.Limit != 0 {
		values.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset != 0 {
		values.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Sort != "" {
		values.Add("sort", params.Sort)
	}
	if params.Range != "" {
		values.Add("range", params.Range)
	}
	if params.Fuzziness != "" {
		values.Add("fuzziness", params.Fuzziness)
	}

	return values
}
