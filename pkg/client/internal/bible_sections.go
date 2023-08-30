package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBibleBookSections(apiKey string, bibleId string, bookId string) (string, error) {
	apiUrl := produceBookSectionsApiUrl(bibleId, bookId)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleChapterSections(apiKey string, bibleId string, chapterId string) (string, error) {
	apiUrl := produceChapterSectionsApiUrl(bibleId, chapterId)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleSectionById(apiKey string, bibleId string, sectionId string, params *params.BibleSectionParams) (string, error) {
	apiUrl := produceBibleSectionApiUrl(bibleId, sectionId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBookSectionsApiUrl(bibleId string, bookId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books/%s/sections", bibleId, bookId)
	return produceGenericUrl(path)
}

func produceChapterSectionsApiUrl(bibleId string, chapterId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s/chapters/%s/sections", bibleId, chapterId)
	return produceGenericUrl(path)
}

func produceBibleSectionApiUrl(bibleId string, sectionId string, params *params.BibleSectionParams) *url.URL {
	path := fmt.Sprintf("bibles/%s/sections/%s", bibleId, sectionId)
	return produceGenericUrlWithQueryParams(path, params)

}
