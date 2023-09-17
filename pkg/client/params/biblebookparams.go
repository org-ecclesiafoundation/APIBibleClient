//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
