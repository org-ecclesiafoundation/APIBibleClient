//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package utils

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/utils/internal"
	"os"
	"time"
)

const (
	RequestTimeout = 5 * time.Second
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

// GetRequestTimeout retrieves the value of the environment variable
// `SCRIPTURE_API_BIBLE_REQUEST_TIMEOUT` as time.Duration. If it is
// not set, then it returns RequestTimeout, which is set to 5 seconds
func GetRequestTimeout() time.Duration {
	fromEnv := os.Getenv("SCRIPTURE_API_BIBLE_REQUEST_TIMEOUT")
	timeout, timeoutErr := time.ParseDuration(fromEnv)
	if timeoutErr != nil {
		return RequestTimeout
	} else {
		return timeout
	}
}
