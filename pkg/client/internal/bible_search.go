//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"ecclesiafoundation.org/apibibleclient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBibleSearchResults(apiKey string, bibleId string, params *params.BibleSearchParams) (string, error) {
	apiUrl := produceBibleSearchApiUrl(bibleId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBibleSearchApiUrl(bibleId string, params *params.BibleSearchParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/search", bibleId)
	return produceGenericUrlWithQueryParams(path, params)
}
