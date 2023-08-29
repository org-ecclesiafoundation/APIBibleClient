package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBiblePassage(apiKey string, bibleId string, passageId string, params *params.BiblePassageParams) (string, error) {
	apiUrl := produceBiblePassageApiUrl(bibleId, passageId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func produceBiblePassageApiUrl(bibleId string, passageId string, params *params.BiblePassageParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/passages/%s", bibleId, passageId)
	return produceGenericUrlWithQueryParams(path, params)
}
