//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package params

import (
	"net/url"
	"strings"
)

// The BibleSectionParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/sections/{sectionId}` end point
// All parameter fields are optional, and default to Golang's default values.
// All blank or default parameters do not show up in the query parameter string
// produced by the ProduceQueryParameters() method.
type BibleSectionParams struct {
	ContentType           string
	IncludeNotes          bool
	IncludeTitles         bool
	IncludeChapterNumbers bool
	IncludeVerseNumbers   bool
	IncludeVerseSpans     bool
	Parallels             []string
}

func (params *BibleSectionParams) ProduceQueryParameters() url.Values {
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
