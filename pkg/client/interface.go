package client

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/internal"
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
)

// The GetBibles function calls the
// `/bibles` end point of the API
func GetBibles(apiKey string, params *params.BiblesParams) (string, error) {
	return internal.GetBibles(apiKey, params)
}

// The GetBibleById function calls the
// `/bibles/{bibleId}` end point of the API
func GetBibleById(apiKey string, bibleId string) (string, error) {
	return internal.GetBibleById(apiKey, bibleId)
}

// The GetAudioBibles function calls the
// `/audio-bibles` end point of the API
func GetAudioBibles(apiKey string, params *params.AudioBiblesParams) (string, error) {
	return internal.GetAudioBibles(apiKey, params)
}

// The GetAudioBibleById function calls the
// `/audio-bibles/{audioBibleId}` end point of the API
func GetAudioBibleById(apiKey string, bibleId string) (string, error) {
	return internal.GetAudioBibleById(apiKey, bibleId)
}

// The GetBibleBooks function calls the
// `/bibles/{bibleId}/books` end point of the API
func GetBibleBooks(apiKey string, bibleId string, params *params.BibleBooksParams) (string, error) {
	return internal.GetBibleBooks(apiKey, bibleId, params)
}

// The GetBibleBookById function calls the
// `/bibles/{bibleId}/books/{bookId}` end point of the API
func GetBibleBookById(apiKey string, bibleId string, bookId string, params *params.BibleBookParams) (string, error) {
	return internal.GetBibleBookById(apiKey, bibleId, bookId, params)
}

// The GetAudioBibleBooks function calls the
// `/audio-bibles/{audioBibleId}/books` end point of the API
func GetAudioBibleBooks(apiKey string, audioBibleId string, params *params.AudioBibleBooksParams) (string, error) {
	return internal.GetAudioBibleBooks(apiKey, audioBibleId, params)
}

// The GetAudioBibleBookById function calls the
// `/audio-bibles/{audioBibleId}/books/{bookId}` end point of the API
func GetAudioBibleBookById(apiKey string, audioBibleId string, bookId string, params *params.AudioBibleBookParams) (string, error) {
	return internal.GetAudioBibleBookById(apiKey, audioBibleId, bookId, params)
}
