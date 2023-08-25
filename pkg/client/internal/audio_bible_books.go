package internal

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"fmt"
	"net/url"
)

func GetAudioBibleBooks(apiKey string, audioBibleId string, params *params.AudioBibleBooksParams) (string, error) {
	apiUrl := produceAudioBibleBooksApiUrl(audioBibleId, params)
	return genericGetRequest(apiKey, apiUrl)
}

func GetAudioBibleBookById(apiKey string, audioBibleId string, bookId string, params *params.AudioBibleBookParams) (string, error) {
	apiUrl := produceAudioBibleBookApiUrl(audioBibleId, bookId)
	return genericGetRequest(apiKey, apiUrl)
}

func produceAudioBibleBooksApiUrl(audioBibleId string, params *params.AudioBibleBooksParams) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s/books", audioBibleId)
	return produceGenericUrlWithQueryParams(path, params)
}

func produceAudioBibleBookApiUrl(audioBibleId string, bookId string) *url.URL {
	path := fmt.Sprintf("/audio-bibles/%s/books/%s", audioBibleId, bookId)
	return produceGenericUrl(path)
}
