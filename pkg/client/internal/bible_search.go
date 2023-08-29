package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBibleSearchResults(apiKey string, bibleId string, params *params.BibleSearchParams) (string, error) {
	apiUrl := produceBibleSearchApiUrl(bibleId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBibleSearchApiUrl(bibleId string, params *params.BibleSearchParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/search", bibleId)
	return produceGenericUrlWithQueryParams(path, params)
}
