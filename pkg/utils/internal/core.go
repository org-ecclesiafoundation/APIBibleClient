//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetJsonField(body string, field string) (string, error) {
	var unmarshalledJsonMap map[string]interface{}
	if unmarshalErr := json.Unmarshal([]byte(body), &unmarshalledJsonMap); unmarshalErr != nil {
		return "", fmt.Errorf("Failed to unmarshall json:\n%s", body)
	} else {
		fieldContents, ok := unmarshalledJsonMap[field]
		if !ok {
			return "", fmt.Errorf("Error getting %s field from the following JSON:\n%s", field, body)
		} else {
			fieldString, jsonMarshalErr := json.Marshal(fieldContents)
			if jsonMarshalErr != nil {
				return "", fmt.Errorf("Error re-marshaling field %s from JSON\n%s", fieldString, body)
			} else {
				return string(fieldString), nil
			}
		}
	}
}

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
