//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"www.ecclesiafoundation.org/apibibleclient/pkg/params"
	"www.ecclesiafoundation.org/apibibleclient/pkg/utils"
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

func produceGenericUrlWithQueryParams(path string, params params.CanProduceQueryParams) *url.URL {
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
