//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.
//  -----------------------------------------------------------------------------
//  License for Examples Produced by the DocTests:
//  MIT License
//  Copyright 2023 Ecclesia Foundation
//  Licensed under the MIT license:
//
//  http://www.opensource.org/licenses/mit-license.php
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in
//  all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//  THE SOFTWARE.

package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestStub(t *testing.T) {}

func ExampleGetJsonField() {
	json := `
{
  "foo": "bar",
  "splat": {
    "baz": [
      "element 1",
      "element 2",
      "element 3"
    ]
  }
}`
	splat, getJsonFieldErr := GetJsonField(json, "splat")
	if getJsonFieldErr != nil {
		fmt.Println("Handle getJsonFieldErr here")
	} else {
		fmt.Println(splat)
	}
	// Output:
	// {"baz":["element 1","element 2","element 3"]}
}

func ExamplePrettify() {
	json := `{"baz":["element 1","element 2","element 3"]}`
	prettyJson, prettifyJsonError := Prettify(json)
	if prettifyJsonError != nil {
		fmt.Println("Handle prettification error here")
	} else {
		fmt.Println(prettyJson)
	}
	// Output:
	// {
	//   "baz": [
	//     "element 1",
	//     "element 2",
	//     "element 3"
	//   ]
	// }
}

func ExampleGetRequestTimeout() {
	// Here we set and unset the environment variable for the timeout
	// to show both how to set, and how to use the default timeout.
	os.Setenv("SCRIPTURE_API_BIBLE_REQUEST_TIMEOUT", "2s")
	timeout := GetRequestTimeout()
	fmt.Println(timeout)

	os.Setenv("SCRIPTURE_API_BIBLE_REQUEST_TIMEOUT", "")
	timeout2 := GetRequestTimeout()
	fmt.Println(timeout2)

	// Output:
	// 2s
	// 5s
}
