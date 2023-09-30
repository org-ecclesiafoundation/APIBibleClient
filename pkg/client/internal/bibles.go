//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"fmt"
	"net/url"
	"www.ecclesiafoundation.org/apibibleclient/pkg/params"
)

func GetBibles(apiKey string, params *params.BiblesParams) (string, error) {
	apiUrl := produceBiblesApiUrl(params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleById(apiKey string, bibleId string) (string, error) {
	apiUrl := produceBibleApiUrl(bibleId)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBiblesApiUrl(params *params.BiblesParams) *url.URL {
	path := "/bibles"
	return produceGenericUrlWithQueryParams(path, params)
}

func produceBibleApiUrl(bibleId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s", bibleId)
	return produceGenericUrl(path)
}
