//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package params

import (
	"net/url"
	"strings"
)

// The BiblesParams struct contains all optional parameters
// for the API call to the `bibles` end point.
// All parameter fields are optional, and default to Golang's default values.
// All blank or default parameters do not show up in the query parameter string
// produced by the ProduceQueryParameters() method.
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
