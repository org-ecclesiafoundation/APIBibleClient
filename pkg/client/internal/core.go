package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

const (
	ApiURL         = "api.scripture.api.bible"
	ApiVersion     = "v1"
	RequestTimeout = 5 * time.Second
)

func RetrieveRequestTimeout() time.Duration {
	fromEnv := os.Getenv("SCRIPTURE_API_BIBLE_REQUEST_TIMEOUT")
	timeout, timeoutErr := time.ParseDuration(fromEnv)
	if timeoutErr != nil {
		return RequestTimeout
	} else {
		return timeout
	}
}

func RetrieveApiKey() (string, error) {
	apiKey := os.Getenv("SCRIPTURE_API_BIBLE_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("Please set environment variable SCRIPTURE_API_BIBLE_KEY \n" +
			"to a valid API key for the scripture.api.bible site")
	}
	return apiKey, nil
}

func produceApiPath(endPoint string) string {
	return path.Join("/", ApiVersion, endPoint)
}

type BiblesString string

func (biblesString *BiblesString) Prettify() (string, error) {
	prettyJson, prettifyErr := PrettifyJson(string(*biblesString))
	if prettifyErr != nil {
		return "", prettifyErr
	} else {
		return prettyJson, nil
	}
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
