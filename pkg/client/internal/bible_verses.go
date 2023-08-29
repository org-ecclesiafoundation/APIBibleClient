package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
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
