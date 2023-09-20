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

func GetBibleChapterVerses(apiKey string, bibleId string, chapterId string) (string, error) {
	apiUrl := produceBibleChapterVersesApiUrl(bibleId, chapterId)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleVerseById(apiKey string, bibleId string, verseId string, params *params.BibleVerseParams) (string, error) {
	apiUrl := produceBibleVerseApiUrl(bibleId, verseId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBibleChapterVersesApiUrl(bibleId string, chapterId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s/chapters/%s/verses", bibleId, chapterId)
	return produceGenericUrl(path)
}

func produceBibleVerseApiUrl(bibleId string, verseId string, params *params.BibleVerseParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/verses/%s", bibleId, verseId)
	return produceGenericUrlWithQueryParams(path, params)
}
