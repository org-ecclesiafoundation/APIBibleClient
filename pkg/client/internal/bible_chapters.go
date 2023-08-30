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

func GetBibleChapters(apiKey string, bibleId string, bookId string) (string, error) {
	apiUrl := produceBibleChaptersApiUrl(bibleId, bookId)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleChapterById(apiKey string, bibleId string, chapterId string, params *params.BibleChapterParams) (string, error) {
	apiUrl := produceBibleChapterApiUrl(bibleId, chapterId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBibleChaptersApiUrl(bibleId string, bookId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books/%s/chapters", bibleId, bookId)
	return produceGenericUrl(path)
}

func produceBibleChapterApiUrl(bibleId string, chapterId string, params *params.BibleChapterParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/chapters/%s", bibleId, chapterId)
	return produceGenericUrlWithQueryParams(path, params)
}
