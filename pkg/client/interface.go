package client

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/internal"
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
)

// Prettify pretty-prints the JSON response
func Prettify(json string) (string, error) {
	return internal.PrettifyJson(json)
}

// RetrieveApiKey retrieves API key from
// the environment variable, SCRIPTURE_API_BIBLE_KEY
func RetrieveApiKey() (string, error) {
	return internal.RetrieveApiKey()
}

// The RetrieveBibles function calls the `bibles` end point of the API
func RetrieveBibles(apiKey string, params params.BiblesParams) (string, error) {
	return internal.RetrieveBibles(apiKey, params)
}
