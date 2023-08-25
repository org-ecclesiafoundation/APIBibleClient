package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"ecclesiafoundation.org/APIBibleClient/pkg/utils"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	ApiURL     = "api.scripture.api.bible"
	ApiVersion = "v1"
)

func produceHttpHeader(apiKey string) http.Header {
	return http.Header{
		"Accept":  {"application/json"},
		"api-key": {apiKey},
	}
}

func produceApiPath(endPoint string) string {
	return path.Join("/", ApiVersion, endPoint)
}

func produceGenericUrl(path string) *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   ApiURL,
		Path:   produceApiPath(path),
	}
}

func produceGenericUrlWithQueryParams(path string, params params.QueryParams) *url.URL {
	return &url.URL{
		Scheme:   "https",
		Host:     ApiURL,
		Path:     produceApiPath(path),
		RawQuery: params.ProduceQueryParameters().Encode(),
	}
}

func produceHttpRequest(apiUrl *url.URL, header http.Header) http.Request {
	return http.Request{
		Method: "GET",
		URL:    apiUrl,
		Header: header,
	}
}

func sendHttpRequest(request http.Request) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: utils.GetRequestTimeout(),
	}
	return httpClient.Do(&request)
}

func handleHttpResponse(response *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
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

func genericGetRequest(apiKey string, apiUrl *url.URL) (string, error) {
	header := produceHttpHeader(apiKey)
	request := produceHttpRequest(apiUrl, header)
	response, getRequestErr := sendHttpRequest(request)
	return handleHttpResponse(response, getRequestErr)
}

func CleanJson(body string) (string, error) {
	var mapForJson map[string]interface{}
	unmarshalErr := json.Unmarshal([]byte(body), &mapForJson)
	if unmarshalErr != nil {
		return "", unmarshalErr
	} else {
		cleanBody, cleanErr := json.Marshal(mapForJson)
		if cleanErr != nil {
			return "", cleanErr
		}
		return string(cleanBody), nil
	}
}
