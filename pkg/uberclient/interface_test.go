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

package uberclient

import (
	"fmt"
	"testing"
	"www.ecclesiafoundation.org/apibibleclient/pkg/client/params"
	"www.ecclesiafoundation.org/apibibleclient/pkg/utils"
)

func TestStub(t *testing.T) {}

func ExampleGetBibleReference() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// All parameter fields are optional, and default to Golang's default values
		bibleVerseParams := params.BiblePassageParams{
			ContentType:           "text",
			IncludeNotes:          false,
			IncludeTitles:         false,
			IncludeChapterNumbers: false,
			IncludeVerseNumbers:   false,
			IncludeVerseSpans:     false,
			Parallels:             nil,
			UseOrgId:              false,
		}
		kjvBibleId := "de4e12af7f28f599-02"
		reference, refErr := GetBibleReference(apiKey, kjvBibleId, "John 3:16-17", &bibleVerseParams)
		if refErr != nil {
			fmt.Println("Do error handling for failing to get bible verse here")
		} else {
			bibleVerseData, getJsonFieldErr := utils.GetJsonField(reference, "data")
			if getJsonFieldErr != nil {
				fmt.Println("Do error handling for failing to get \"data\" field from JSON here")
			} else {
				prettyVerseData, prettifyErr := utils.Prettify(bibleVerseData)
				if prettifyErr != nil {
					fmt.Println("Do error handling for failing to make pretty JSON here")
				} else {
					fmt.Println(prettyVerseData)
				}
			}
		}
	}
	// Output:
	// {
	//   "bibleId": "de4e12af7f28f599-02",
	//   "bookId": "JHN",
	//   "chapterIds": [
	//     "JHN.3"
	//   ],
	//   "content": "     [16] ¶ For God so loved the world, that he gave his only begotten Son, that whosoever believeth in him should not perish, but have everlasting life.  [17] For God sent not his Son into the world to condemn the world; but that the world through him might be saved.\n",
	//   "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
	//   "id": "JHN.3.16-JHN.3.17",
	//   "orgId": "JHN.3.16-JHN.3.17",
	//   "reference": "John 3:16-17",
	//   "verseCount": 2
	// }
}
