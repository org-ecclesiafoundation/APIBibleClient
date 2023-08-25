package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetApiKey() (string, error) {
	apiKey := os.Getenv("SCRIPTURE_API_BIBLE_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("Please set environment variable SCRIPTURE_API_BIBLE_KEY \n" +
			"to a valid API key for the scripture.api.bible site")
	}
	return apiKey, nil
}

func PrettifyJson(body string) (string, error) {
	var mapForJson map[string]interface{}
	unmarshalErr := json.Unmarshal([]byte(body), &mapForJson)
	if unmarshalErr != nil {
		return "", unmarshalErr
	} else {
		prettyBody, prettifyErr := json.MarshalIndent(mapForJson, "", "  ")
		if prettifyErr != nil {
			return "", prettifyErr
		}
		return string(prettyBody), nil
	}
}
