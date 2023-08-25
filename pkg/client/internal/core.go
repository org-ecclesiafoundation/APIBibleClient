package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

const (
	ApiURL         = "api.scripture.api.bible"
	ApiVersion     = "v1"
	RequestTimeout = 5 * time.Second
)

func GetRequestTimeout() time.Duration {
	fromEnv := os.Getenv("SCRIPTURE_API_BIBLE_REQUEST_TIMEOUT")
	timeout, timeoutErr := time.ParseDuration(fromEnv)
	if timeoutErr != nil {
		return RequestTimeout
	} else {
		return timeout
	}
}

func GetApiKey() (string, error) {
	apiKey := os.Getenv("SCRIPTURE_API_BIBLE_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("Please set environment variable SCRIPTURE_API_BIBLE_KEY \n" +
			"to a valid API key for the scripture.api.bible site")
	}
	return apiKey, nil
}

func produceHttpHeader(apiKey string) http.Header {
	return http.Header{
		"Accept":  {"application/json"},
		"api-key": {apiKey},
	}
}

func produceApiPath(endPoint string) string {
	return path.Join("/", ApiVersion, endPoint)
}

func produceHttpRequest(apiUrl *url.URL, header http.Header) http.Request {
	return http.Request{
		Method: "GET",
		URL:    apiUrl,
		Header: header,
	}
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
