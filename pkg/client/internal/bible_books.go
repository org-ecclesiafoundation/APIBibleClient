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

func GetBibleBooks(apiKey string, bibleId string, params *params.BibleBooksParams) (string, error) {
	apiUrl := produceBibleBooksApiUrl(bibleId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleBookById(apiKey string, bibleId string, bookId string, params *params.BibleBookParams) (string, error) {
	apiUrl := produceBibleBookApiUrl(bibleId, bookId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBibleBooksApiUrl(bibleId string, params *params.BibleBooksParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books", bibleId)
	return produceGenericUrlWithQueryParams(path, params)
}

func produceBibleBookApiUrl(bibleId string, bookId string, params *params.BibleBookParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books/%s", bibleId, bookId)
	return produceGenericUrlWithQueryParams(path, params)
}
