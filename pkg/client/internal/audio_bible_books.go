//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"fmt"
	"net/url"
	params2 "www.ecclesiafoundation.org/apibibleclient/pkg/params"
)

func GetAudioBibleBooks(apiKey string, audioBibleId string, params *params2.AudioBibleBooksParams) (string, error) {
	apiUrl := produceAudioBibleBooksApiUrl(audioBibleId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetAudioBibleBookById(apiKey string, audioBibleId string, bookId string, params *params2.AudioBibleBookParams) (string, error) {
	apiUrl := produceAudioBibleBookApiUrl(audioBibleId, bookId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceAudioBibleBooksApiUrl(audioBibleId string, params *params2.AudioBibleBooksParams) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s/books", audioBibleId)
	return produceGenericUrlWithQueryParams(path, params)
}

func produceAudioBibleBookApiUrl(audioBibleId string, bookId string, params *params2.AudioBibleBookParams) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s/books/%s", audioBibleId, bookId)
	return produceGenericUrlWithQueryParams(path, params)
}
