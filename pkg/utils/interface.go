package utils

import "ecclesiafoundation.org/APIBibleClient/pkg/utils/internal"

// Prettify pretty-prints the JSON response
func Prettify(json string) (string, error) {
	return internal.PrettifyJson(json)
}
