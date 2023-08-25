package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"ecclesiafoundation.org/APIBibleClient/pkg/utils"
	"fmt"
	"net/http"
	"net/url"
)

func GetAudioBibles(apiKey string, params *params.AudioBiblesParams) (string, error) {
	header := produceHttpHeader(apiKey)
	apiUrl := produceAudioBiblesApiUrl(params)
	request := produceHttpRequest(apiUrl, header)
	httpClient := &http.Client{
		Timeout: utils.GetRequestTimeout(),
	}
	response, getRequestErr := httpClient.Do(&request)
	return handleHttpResponse(response, getRequestErr)
}

func GetAudioBibleById(apiKey string, audioBibleId string) (string, error) {
	header := produceHttpHeader(apiKey)
	apiUrl := produceAudioBibleApiUrl(audioBibleId)
	request := produceHttpRequest(apiUrl, header)
	response, getRequestErr := sendHttpRequest(request)
	return handleHttpResponse(response, getRequestErr)
}

func produceAudioBiblesApiUrl(params *params.AudioBiblesParams) *url.URL {
	path := "/audio-bibles"
	return &url.URL{
		Scheme:   "https",
		Host:     ApiURL,
		Path:     produceApiPath(path),
		RawQuery: params.ProduceQueryParameters().Encode(),
	}
}

func produceAudioBibleApiUrl(audioBibleId string) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s", audioBibleId)
	return &url.URL{
		Scheme: "https",
		Host:   ApiURL,
		Path:   produceApiPath(path),
	}
}
