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

func GetAudioBibles(apiKey string, params *params.AudioBiblesParams) (string, error) {
	apiUrl := produceAudioBiblesApiUrl(params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetAudioBibleById(apiKey string, audioBibleId string) (string, error) {
	apiUrl := produceAudioBibleApiUrl(audioBibleId)
	return genericGetRequest(apiKey, apiUrl)
}

func produceAudioBiblesApiUrl(params *params.AudioBiblesParams) *url.URL {
	path := "/audio-bibles"
	return produceGenericUrlWithQueryParams(path, params)
}

func produceAudioBibleApiUrl(audioBibleId string) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s", audioBibleId)
	return produceGenericUrl(path)
}
