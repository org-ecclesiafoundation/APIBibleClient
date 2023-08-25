package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"ecclesiafoundation.org/APIBibleClient/pkg/utils"
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

func produceAudioBiblesApiUrl(params *params.AudioBiblesParams) *url.URL {
	path := "/audio-bibles"
	return &url.URL{
		Scheme:   "https",
		Host:     ApiURL,
		Path:     produceApiPath(path),
		RawQuery: params.ProduceQueryParameters().Encode(),
	}
}
