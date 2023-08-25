package client

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/internal"
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
)

// The GetBibles function calls the `/bibles` end point of the API
func GetBibles(apiKey string, params params.BiblesParams) (string, error) {
	return internal.GetBibles(apiKey, &params)
}

// The GetBibleById function calls the `/bibles/{bibleId}` end point of the API
func GetBibleById(apiKey string, bibleId string) (string, error) {
	return internal.GetBibleById(apiKey, bibleId)
}
