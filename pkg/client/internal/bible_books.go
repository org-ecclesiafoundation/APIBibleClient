package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetBibleBooks(apiKey string, bibleId string, params params.BibleBookParams) (string, error) {
	header := produceHttpHeader(apiKey)
	apiUrl := produceBibleBooksApiUrl(bibleId, params)
	request := produceHttpRequest(apiUrl, header)
	response, getRequestErr := sendHttpRequest(request)
	return handleHttpResponse(response, getRequestErr)
}

func produceBibleBooksApiUrl(bibleId string, params params.BibleBookParams) *url.URL {
	path := fmt.Sprintf("/bibles/%s/books", bibleId)
	return &url.URL{
		Scheme:   "https",
		Host:     ApiURL,
		Path:     produceApiPath(path),
		RawQuery: params.ProduceQueryParameters().Encode(),
	}
}
