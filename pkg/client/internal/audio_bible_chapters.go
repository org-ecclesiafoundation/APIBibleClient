//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"fmt"
	"net/url"
)

func GetAudioBibleChapters(apiKey string, audioBibleId string, bookId string) (string, error) {
	apiUrl := produceAudioBibleChaptersApiUrl(audioBibleId, bookId)
	return genericGetRequest(apiKey, apiUrl)
}

func GetAudioBibleChapterById(apiKey string, audioBibleId string, chapterId string) (string, error) {
	apiUrl := produceAudioBibleChapterApiUrl(audioBibleId, chapterId)
	return genericGetRequest(apiKey, apiUrl)
}

func produceAudioBibleChaptersApiUrl(audioBibleId string, bookId string) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s/books/%s/chapters", audioBibleId, bookId)
	return produceGenericUrl(path)
}

func produceAudioBibleChapterApiUrl(audioBibleId string, chapterId string) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s/chapters/%s", audioBibleId, chapterId)
	return produceGenericUrl(path)
}
