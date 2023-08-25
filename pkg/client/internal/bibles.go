package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBibles(apiKey string, params *params.BiblesParams) (string, error) {
	apiUrl := produceBiblesApiUrl(params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetBibleById(apiKey string, bibleId string) (string, error) {
	apiUrl := produceBibleApiUrl(bibleId)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBiblesApiUrl(params *params.BiblesParams) *url.URL {
	path := "/bibles"
	return produceGenericUrlWithQueryParams(path, params)
}

func produceBibleApiUrl(bibleId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s", bibleId)
	return produceGenericUrl(path)
}
