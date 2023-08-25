package utils

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/utils/internal"
	"time"
)

// Prettify pretty-prints the JSON response
func Prettify(json string) (string, error) {
	return internal.PrettifyJson(json)
}

// GetApiKey retrieves API key from
// the environment variable, SCRIPTURE_API_BIBLE_KEY
func GetApiKey() (string, error) {
	return internal.GetApiKey()
}

func GetRequestTimeout() time.Duration {
	return internal.GetRequestTimeout()
}
