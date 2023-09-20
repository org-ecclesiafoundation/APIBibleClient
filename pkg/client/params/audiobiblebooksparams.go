//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package params

import "net/url"

// The AudioBibleBooksParams struct contains all optional parameters
// for the API call to the `bibles/{audioBibleId}/books` end point
// All parameter fields are optional, and default to Golang's default values.
// All blank or default parameters do not show up in the query parameter string
// produced by the ProduceQueryParameters() method.
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
		values.Add("include-chapters-and-sections", "true")
	}
	return values
}
