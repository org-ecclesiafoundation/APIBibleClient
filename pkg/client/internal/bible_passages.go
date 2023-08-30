//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBiblePassage(apiKey string, bibleId string, passageId string, params *params.BiblePassageParams) (string, error) {
	apiUrl := produceBiblePassageApiUrl(bibleId, passageId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBiblePassageApiUrl(bibleId string, passageId string, params *params.BiblePassageParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/passages/%s", bibleId, passageId)
	return produceGenericUrlWithQueryParams(path, params)
}
