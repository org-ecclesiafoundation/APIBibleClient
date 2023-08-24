package internal

import (
	"ecclesiafoundation.org/scripture-api-bible-client/pkg/client/params"
	"io"
	"net/http"
	"net/url"
)

func RetrieveBibles(apiKey string, params params.BiblesParams) (string, error) {
	header := http.Header{
		"Accept":  {"application/json"},
		"api-key": {apiKey},
	}
	apiUrl := &url.URL{
		Scheme:   "https",
		Host:     ApiURL,
		Path:     produceApiPath("/bibles"),
		RawQuery: params.ProduceQueryParameters().Encode(),
	}
	request := http.Request{
		Method: "GET",
		URL:    apiUrl,
		Header: header,
	}
	httpClient := &http.Client{
		Timeout: RetrieveRequestTimeout(),
	}
	response, getRequestErr := httpClient.Do(&request)
	if getRequestErr != nil {
		return "", getRequestErr
	} else {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(response.Body)
		body, bodyReadErr := io.ReadAll(response.Body)
		if bodyReadErr != nil {
			return "", bodyReadErr
		} else {
			return CleanJson(string(body))
		}
	}
}
