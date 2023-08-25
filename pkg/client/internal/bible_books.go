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

func produceBibleBooksApiUrl(bibleId string, params *params.BibleBooksParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books", bibleId)
	return produceGenericUrlWithQueryParams(path, params)
}
