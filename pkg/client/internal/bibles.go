package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/http"
	"net/url"
)

func GetBibles(apiKey string, params *params.BiblesParams) (string, error) {
	header := produceHttpHeader(apiKey)
	apiUrl := produceBiblesApiUrl(params)
	request := produceHttpRequest(apiUrl, header)
	httpClient := &http.Client{
		Timeout: GetRequestTimeout(),
	}
	response, getRequestErr := httpClient.Do(&request)
	return handleHttpResponse(response, getRequestErr)
}

func GetBibleById(apiKey string, bibleId string) (string, error) {
	header := produceHttpHeader(apiKey)
	apiUrl := produceBibleApiUrl(bibleId)
	request := produceHttpRequest(apiUrl, header)
	httpClient := &http.Client{
		Timeout: GetRequestTimeout(),
	}
	response, getRequestErr := httpClient.Do(&request)
	return handleHttpResponse(response, getRequestErr)
}

func produceBiblesApiUrl(params *params.BiblesParams) *url.URL {
	path := "/bibles"
	return &url.URL{
		Scheme:   "https",
		Host:     ApiURL,
		Path:     produceApiPath(path),
		RawQuery: params.ProduceQueryParameters().Encode(),
	}
}

func produceBibleApiUrl(bibleId string) *url.URL {
	path := fmt.Sprintf("/bibles/%s", bibleId)
	return &url.URL{
		Scheme: "https",
		Host:   ApiURL,
		Path:   produceApiPath(path),
	}
}
