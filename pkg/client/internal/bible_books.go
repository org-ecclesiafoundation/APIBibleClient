package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBibleBooks(apiKey string, bibleId string, params *params.BibleBooksParams) (string, error) {
	apiUrl := produceBibleBooksApiUrl(bibleId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleBookById(apiKey string, bibleId string, bookId string, params *params.BibleBooksParams) (string, error) {
	apiUrl := produceBibleBookApiUrl(bibleId, bookId)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBibleBooksApiUrl(bibleId string, params *params.BibleBooksParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books", bibleId)
	return produceGenericUrlWithQueryParams(path, params)
}

func produceBibleBookApiUrl(bibleId string, bookId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books/%s", bibleId, bookId)
	return produceGenericUrl(path)
}
