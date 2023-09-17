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

package client

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"ecclesiafoundation.org/APIBibleClient/pkg/utils"
	"fmt"
	"testing"
)

func TestStub(t *testing.T) {}

func ExampleGetBibles() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleParams{} struct ref
		// to the call to GetBibles.
		kjvBibleId := "de4e12af7f28f599-02"
		biblesParams := params.BiblesParams{
			Language:           "eng",
			Abbreviation:       "kjv",
			Name:               "King James",
			Ids:                []string{kjvBibleId},
			IncludeFullDetails: true,
		}
		// Here is an example of an API call
		bibles, biblesErr := GetBibles(apiKey, &biblesParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if biblesErr != nil {
			fmt.Println("Do error handling for failing to get bibles here")
		} else {
			prettyBibles, prettifyErr := utils.Prettify(bibles)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBibles)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "abbreviation": "engKJV",
	//       "abbreviationLocal": "KJV",
	//       "audioBibles": [],
	//       "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
	//       "countries": [
	//         {
	//           "id": "GB",
	//           "name": "United Kingdom of Great Britain and Northern Ireland",
	//           "nameLocal": "United Kingdom of Great Britain and Northern Ireland"
	//         }
	//       ],
	//       "dblId": "de4e12af7f28f599",
	//       "description": "Protestant",
	//       "descriptionLocal": "Protestant",
	//       "id": "de4e12af7f28f599-02",
	//       "info": "\u003cp\u003eThis historical translation of the Holy Bible is brought to you in quality digital format by the \u003ca href=\"https://crosswire.org/\"\u003eCrosswire Bible Society\u003c/a\u003e and \u003ca href=\"https://eBible.org\"\u003eeBible.org\u003c/a\u003e.\u003c/p\u003e",
	//       "language": {
	//         "id": "eng",
	//         "name": "English",
	//         "nameLocal": "English",
	//         "script": "Latin",
	//         "scriptDirection": "LTR"
	//       },
	//       "name": "King James (Authorised) Version",
	//       "nameLocal": "King James Version",
	//       "relatedDbl": null,
	//       "type": "text",
	//       "updatedAt": "2023-05-03T09:23:59.000Z"
	//     }
	//   ]
	// }
}

func ExampleGetBibleById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		bible, bibleErr := GetBibleById(apiKey, kjvBibleId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleErr != nil {
			fmt.Println("Do error handling for failing to get bible here")
		} else {
			prettyBible, prettifyErr := utils.Prettify(bible)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBible)
			}
		}
	}
	// Output:
	// {
	//   "data": {
	//     "abbreviation": "engKJV",
	//     "abbreviationLocal": "KJV",
	//     "audioBibles": [],
	//     "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
	//     "countries": [
	//       {
	//         "id": "GB",
	//         "name": "United Kingdom of Great Britain and Northern Ireland",
	//         "nameLocal": "United Kingdom of Great Britain and Northern Ireland"
	//       }
	//     ],
	//     "dblId": "de4e12af7f28f599",
	//     "description": "Protestant",
	//     "descriptionLocal": "Protestant",
	//     "id": "de4e12af7f28f599-02",
	//     "info": "\u003cp\u003eThis historical translation of the Holy Bible is brought to you in quality digital format by the \u003ca href=\"https://crosswire.org/\"\u003eCrosswire Bible Society\u003c/a\u003e and \u003ca href=\"https://eBible.org\"\u003eeBible.org\u003c/a\u003e.\u003c/p\u003e",
	//     "language": {
	//       "id": "eng",
	//       "name": "English",
	//       "nameLocal": "English",
	//       "script": "Latin",
	//       "scriptDirection": "LTR"
	//     },
	//     "name": "King James (Authorised) Version",
	//     "nameLocal": "King James Version",
	//     "relatedDbl": null,
	//     "type": "text",
	//     "updatedAt": "2023-05-03T09:23:59.000Z"
	//   }
	// }
}

func ExampleGetAudioBibles() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.AudioBiblesParams{} struct ref
		// to the call to Foo.
		worldEnglishBibleID := "9879dbb7cfe39e4d-01"
		audioBiblesParams := params.AudioBiblesParams{
			Language:           "eng",
			Abbreviation:       "WEB13",
			Name:               "World English",
			RelatedTextBibleId: worldEnglishBibleID,
			IncludeFullDetails: true,
		}
		// Here is an example of an API call
		audioBibles, audioBiblesErr := GetAudioBibles(apiKey, &audioBiblesParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if audioBiblesErr != nil {
			fmt.Println("Do error handling for failing to get audio bibles here")
		} else {
			prettyAudioBible, prettifyErr := utils.Prettify(audioBibles)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyAudioBible)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "abbreviation": "WEB13",
	//       "abbreviationLocal": "WEB13",
	//       "copyright": "℗ 2013 Hosanna",
	//       "countries": [
	//         {
	//           "id": "US",
	//           "name": "United States",
	//           "nameLocal": "United States"
	//         }
	//       ],
	//       "dblId": "105a06b6146d11e7",
	//       "description": null,
	//       "descriptionLocal": null,
	//       "id": "105a06b6146d11e7-01",
	//       "info": null,
	//       "language": {
	//         "id": "eng",
	//         "name": "English",
	//         "nameLocal": "English",
	//         "script": "Latin",
	//         "scriptDirection": "LTR"
	//       },
	//       "name": "English - World English Bible 2013 (Drama NT)",
	//       "nameLocal": "English - World English Bible 2013 (Drama NT)",
	//       "relatedDbl": "9879dbb7cfe39e4d",
	//       "type": "audio",
	//       "updatedAt": "2022-01-07T18:49:57.000Z"
	//     }
	//   ]
	// }
}

func ExampleGetAudioBibleById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		webAudioBibleId := "105a06b6146d11e7-01"
		audioBible, audioBibleErr := GetAudioBibleById(apiKey, webAudioBibleId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if audioBibleErr != nil {
			fmt.Println("Do error handling for failing to get audio audioBible here")
		} else {
			prettyAudioBible, prettifyErr := utils.Prettify(audioBible)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyAudioBible)
			}
		}
	}
	// Output:
	// {
	//   "data": {
	//     "abbreviation": "WEB13",
	//     "abbreviationLocal": "WEB13",
	//     "copyright": "℗ 2013 Hosanna",
	//     "countries": [
	//       {
	//         "id": "US",
	//         "name": "United States",
	//         "nameLocal": "United States"
	//       }
	//     ],
	//     "dblId": "105a06b6146d11e7",
	//     "description": null,
	//     "descriptionLocal": null,
	//     "id": "105a06b6146d11e7-01",
	//     "info": null,
	//     "language": {
	//       "id": "eng",
	//       "name": "English",
	//       "nameLocal": "English",
	//       "script": "Latin",
	//       "scriptDirection": "LTR"
	//     },
	//     "name": "English - World English Bible 2013 (Drama NT)",
	//     "nameLocal": "English - World English Bible 2013 (Drama NT)",
	//     "relatedDbl": "9879dbb7cfe39e4d",
	//     "type": "audio",
	//     "updatedAt": "2022-01-07T18:49:57.000Z"
	//   }
	// }
}

func ExampleGetBibleBooks() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleBooksParams{} struct ref
		// to the call to Foo.
		bibleBooksParams := params.BibleBooksParams{
			IncludeChapters:            false,
			IncludeChaptersAndSections: false,
		}
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		bibleBooks, bibleBooksErr := GetBibleBooks(apiKey, kjvBibleId, &bibleBooksParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleBooksErr != nil {
			fmt.Println("Do error handling for failing to get audio audioBible here")
		} else {
			prettyBibleBooks, prettifyErr := utils.Prettify(bibleBooks)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBibleBooks)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "abbreviation": "Gen",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "GEN",
	//       "name": "Genesis",
	//       "nameLong": "The First Book of Moses, called Genesis"
	//     },
	//     {
	//       "abbreviation": "Exo",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "EXO",
	//       "name": "Exodus",
	//       "nameLong": "The Second Book of Moses, called Exodus"
	//     },
	//     {
	//       "abbreviation": "Lev",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "LEV",
	//       "name": "Leviticus",
	//       "nameLong": "The Third Book of Moses, called Leviticus"
	//     },
	//     {
	//       "abbreviation": "Num",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "NUM",
	//       "name": "Numbers",
	//       "nameLong": "The Fourth Book of Moses, called Numbers"
	//     },
	//     {
	//       "abbreviation": "Deu",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "DEU",
	//       "name": "Deuteronomy",
	//       "nameLong": "The Fifth Book of Moses, called Deuteronomy"
	//     },
	//     {
	//       "abbreviation": "Jos",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JOS",
	//       "name": "Joshua",
	//       "nameLong": "The Book of Joshua"
	//     },
	//     {
	//       "abbreviation": "Jdg",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JDG",
	//       "name": "Judges",
	//       "nameLong": "The Book of Judges"
	//     },
	//     {
	//       "abbreviation": "Rut",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "RUT",
	//       "name": "Ruth",
	//       "nameLong": "The Book of Ruth"
	//     },
	//     {
	//       "abbreviation": "1Sa",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1SA",
	//       "name": "1 Samuel",
	//       "nameLong": "The First Book of Samuel Otherwise Called The First Book of the Kings"
	//     },
	//     {
	//       "abbreviation": "2Sa",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2SA",
	//       "name": "2 Samuel",
	//       "nameLong": "The Second Book of Samuel Otherwise Called The Second Book of the Kings"
	//     },
	//     {
	//       "abbreviation": "1Ki",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1KI",
	//       "name": "1 Kings",
	//       "nameLong": "The First Book of the Kings, Commonly Called the Third Book of the Kings"
	//     },
	//     {
	//       "abbreviation": "2Ki",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2KI",
	//       "name": "2 Kings",
	//       "nameLong": "The Second Book of the Kings, Commonly Called the Fourth Book of the Kings"
	//     },
	//     {
	//       "abbreviation": "1Ch",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1CH",
	//       "name": "1 Chronicles",
	//       "nameLong": "The First Book of the Chronicles"
	//     },
	//     {
	//       "abbreviation": "2Ch",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2CH",
	//       "name": "2 Chronicles",
	//       "nameLong": "The Second Book of the Chronicles"
	//     },
	//     {
	//       "abbreviation": "Ezr",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "EZR",
	//       "name": "Ezra",
	//       "nameLong": "Ezra"
	//     },
	//     {
	//       "abbreviation": "Neh",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "NEH",
	//       "name": "Nehemiah",
	//       "nameLong": "The Book of Nehemiah"
	//     },
	//     {
	//       "abbreviation": "Est",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "EST",
	//       "name": "Esther",
	//       "nameLong": "The Book of Esther"
	//     },
	//     {
	//       "abbreviation": "Job",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JOB",
	//       "name": "Job",
	//       "nameLong": "The Book of Job"
	//     },
	//     {
	//       "abbreviation": "Psa",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "PSA",
	//       "name": "Psalms",
	//       "nameLong": "The Book of Psalms"
	//     },
	//     {
	//       "abbreviation": "Pro",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "PRO",
	//       "name": "Proverbs",
	//       "nameLong": "The Proverbs"
	//     },
	//     {
	//       "abbreviation": "Ecc",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "ECC",
	//       "name": "Ecclesiastes",
	//       "nameLong": "Ecclesiastes or, the Preacher"
	//     },
	//     {
	//       "abbreviation": "Sng",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "SNG",
	//       "name": "Song of Solomon",
	//       "nameLong": "The Song of Solomon"
	//     },
	//     {
	//       "abbreviation": "Isa",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "ISA",
	//       "name": "Isaiah",
	//       "nameLong": "The Book of the Prophet Isaiah"
	//     },
	//     {
	//       "abbreviation": "Jer",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JER",
	//       "name": "Jeremiah",
	//       "nameLong": "The Book of the Prophet Jeremiah"
	//     },
	//     {
	//       "abbreviation": "Lam",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "LAM",
	//       "name": "Lamentations",
	//       "nameLong": "The Lamentations of Jeremiah"
	//     },
	//     {
	//       "abbreviation": "Ezk",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "EZK",
	//       "name": "Ezekiel",
	//       "nameLong": "The Book of the Prophet Ezekiel"
	//     },
	//     {
	//       "abbreviation": "Dan",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "DAN",
	//       "name": "Daniel",
	//       "nameLong": "The Book of Daniel"
	//     },
	//     {
	//       "abbreviation": "Hos",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "HOS",
	//       "name": "Hosea",
	//       "nameLong": "Hosea"
	//     },
	//     {
	//       "abbreviation": "Jol",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JOL",
	//       "name": "Joel",
	//       "nameLong": "Joel"
	//     },
	//     {
	//       "abbreviation": "Amo",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "AMO",
	//       "name": "Amos",
	//       "nameLong": "Amos"
	//     },
	//     {
	//       "abbreviation": "Oba",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "OBA",
	//       "name": "Obadiah",
	//       "nameLong": "Obadiah"
	//     },
	//     {
	//       "abbreviation": "Jon",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JON",
	//       "name": "Jonah",
	//       "nameLong": "Jonah"
	//     },
	//     {
	//       "abbreviation": "Mic",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "MIC",
	//       "name": "Micah",
	//       "nameLong": "Micah"
	//     },
	//     {
	//       "abbreviation": "Nam",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "NAM",
	//       "name": "Nahum",
	//       "nameLong": "Nahum"
	//     },
	//     {
	//       "abbreviation": "Hab",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "HAB",
	//       "name": "Habakkuk",
	//       "nameLong": "Habakkuk"
	//     },
	//     {
	//       "abbreviation": "Zep",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "ZEP",
	//       "name": "Zephaniah",
	//       "nameLong": "Zephaniah"
	//     },
	//     {
	//       "abbreviation": "Hag",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "HAG",
	//       "name": "Haggai",
	//       "nameLong": "Haggai"
	//     },
	//     {
	//       "abbreviation": "Zec",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "ZEC",
	//       "name": "Zechariah",
	//       "nameLong": "Zechariah"
	//     },
	//     {
	//       "abbreviation": "Mal",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "MAL",
	//       "name": "Malachi",
	//       "nameLong": "Malachi"
	//     },
	//     {
	//       "abbreviation": "Mat",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "MAT",
	//       "name": "Matthew",
	//       "nameLong": "THE GOSPEL ACCORDING TO ST. MATTHEW"
	//     },
	//     {
	//       "abbreviation": "Mrk",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "MRK",
	//       "name": "Mark",
	//       "nameLong": "THE GOSPEL ACCORDING TO ST. MARK"
	//     },
	//     {
	//       "abbreviation": "Luk",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "LUK",
	//       "name": "Luke",
	//       "nameLong": "THE GOSPEL ACCORDING TO ST. LUKE"
	//     },
	//     {
	//       "abbreviation": "Jhn",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JHN",
	//       "name": "John",
	//       "nameLong": "THE GOSPEL ACCORDING TO ST. JOHN"
	//     },
	//     {
	//       "abbreviation": "Act",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "ACT",
	//       "name": "Acts",
	//       "nameLong": "THE ACTS OF THE APOSTLES"
	//     },
	//     {
	//       "abbreviation": "Rom",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "ROM",
	//       "name": "Romans",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO THE ROMANS"
	//     },
	//     {
	//       "abbreviation": "1Co",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1CO",
	//       "name": "1 Corinthians",
	//       "nameLong": "THE FIRST EPISTLE OF PAUL THE APOSTLE TO THE CORINTHIANS"
	//     },
	//     {
	//       "abbreviation": "2Co",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2CO",
	//       "name": "2 Corinthians",
	//       "nameLong": "THE SECOND EPISTLE OF PAUL THE APOSTLE TO THE CORINTHIANS"
	//     },
	//     {
	//       "abbreviation": "Gal",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "GAL",
	//       "name": "Galatians",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO THE GALATIANS"
	//     },
	//     {
	//       "abbreviation": "Eph",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "EPH",
	//       "name": "Ephesians",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO THE EPHESIANS"
	//     },
	//     {
	//       "abbreviation": "Php",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "PHP",
	//       "name": "Philippians",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO THE PHILIPPIANS"
	//     },
	//     {
	//       "abbreviation": "Col",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "COL",
	//       "name": "Colossians",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO THE COLOSSIANS"
	//     },
	//     {
	//       "abbreviation": "1Th",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1TH",
	//       "name": "1 Thessalonians",
	//       "nameLong": "THE FIRST EPISTLE OF PAUL THE APOSTLE TO THE THESSALONIANS"
	//     },
	//     {
	//       "abbreviation": "2Th",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2TH",
	//       "name": "2 Thessalonians",
	//       "nameLong": "THE SECOND EPISTLE OF PAUL THE APOSTLE TO THE THESSALONIANS"
	//     },
	//     {
	//       "abbreviation": "1Ti",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1TI",
	//       "name": "1 Timothy",
	//       "nameLong": "THE FIRST EPISTLE OF PAUL THE APOSTLE TO TIMOTHY"
	//     },
	//     {
	//       "abbreviation": "2Ti",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2TI",
	//       "name": "2 Timothy",
	//       "nameLong": "THE SECOND EPISTLE OF PAUL THE APOSTLE TO TIMOTHY"
	//     },
	//     {
	//       "abbreviation": "Tit",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "TIT",
	//       "name": "Titus",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO TITUS"
	//     },
	//     {
	//       "abbreviation": "Phm",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "PHM",
	//       "name": "Philemon",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO PHILEMON"
	//     },
	//     {
	//       "abbreviation": "Heb",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "HEB",
	//       "name": "Hebrews",
	//       "nameLong": "THE EPISTLE OF PAUL THE APOSTLE TO THE HEBREWS"
	//     },
	//     {
	//       "abbreviation": "Jas",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JAS",
	//       "name": "James",
	//       "nameLong": "THE GENERAL EPISTLE OF JAMES"
	//     },
	//     {
	//       "abbreviation": "1Pe",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1PE",
	//       "name": "1 Peter",
	//       "nameLong": "THE FIRST EPISTLE GENERAL OF PETER"
	//     },
	//     {
	//       "abbreviation": "2Pe",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2PE",
	//       "name": "2 Peter",
	//       "nameLong": "THE SECOND EPISTLE GENERAL OF PETER"
	//     },
	//     {
	//       "abbreviation": "1Jn",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "1JN",
	//       "name": "1 John",
	//       "nameLong": "THE FIRST EPISTLE GENERAL OF JOHN"
	//     },
	//     {
	//       "abbreviation": "2Jn",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "2JN",
	//       "name": "2 John",
	//       "nameLong": "THE SECOND EPISTLE OF JOHN"
	//     },
	//     {
	//       "abbreviation": "3Jn",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "3JN",
	//       "name": "3 John",
	//       "nameLong": "THE THIRD EPISTLE OF JOHN"
	//     },
	//     {
	//       "abbreviation": "Jud",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "JUD",
	//       "name": "Jude",
	//       "nameLong": "THE GENERAL EPISTLE OF JUDE"
	//     },
	//     {
	//       "abbreviation": "Rev",
	//       "bibleId": "de4e12af7f28f599-02",
	//       "id": "REV",
	//       "name": "Revelation",
	//       "nameLong": "THE REVELATION OF ST. JOHN THE DIVINE"
	//     }
	//   ]
	// }
}

func ExampleGetBibleBookById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleBookParams{} struct ref
		// to the call to GetBibleBookById.
		bibleBookParams := params.BibleBookParams{
			IncludeChapters: true,
		}
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		secondJohnId := "2JN"
		bibleBook, bibleBookErr := GetBibleBookById(apiKey, kjvBibleId, secondJohnId, &bibleBookParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleBookErr != nil {
			fmt.Println("Do error handling for failing to get bible book here")
		} else {
			prettyBibleBook, prettifyErr := utils.Prettify(bibleBook)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBibleBook)
			}
		}
	}
	// Output:
	// {
	//   "data": {
	//     "abbreviation": "2Jn",
	//     "bibleId": "de4e12af7f28f599-02",
	//     "chapters": [
	//       {
	//         "bibleId": "de4e12af7f28f599-02",
	//         "bookId": "2JN",
	//         "id": "2JN.intro",
	//         "number": "intro",
	//         "position": 1226
	//       },
	//       {
	//         "bibleId": "de4e12af7f28f599-02",
	//         "bookId": "2JN",
	//         "id": "2JN.1",
	//         "number": "1",
	//         "position": 1227
	//       }
	//     ],
	//     "id": "2JN",
	//     "name": "2 John",
	//     "nameLong": "THE SECOND EPISTLE OF JOHN"
	//   }
	// }
}

func ExampleGetAudioBibleBooks() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.AudioBibleBooksParams{} struct ref
		// to the call to GetAudioBibleBooks.
		audioBibleBooksParams := params.AudioBibleBooksParams{
			IncludeChapters:            false,
			IncludeChaptersAndSections: false,
		}
		// Here is an example of an API call
		webAudioBibleId := "105a06b6146d11e7-01"
		audioBibleBooks, audioBibleBooksErr := GetAudioBibleBooks(apiKey, webAudioBibleId, &audioBibleBooksParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if audioBibleBooksErr != nil {
			fmt.Println("Do error handling for failing to get audio bible books here")
		} else {
			prettyAudioBibleBooks, prettifyErr := utils.Prettify(audioBibleBooks)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyAudioBibleBooks)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "abbreviation": "1 Corinthians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "1CO",
	//       "name": "1 Corinthians",
	//       "nameLong": "Paul’s First Letter to the Corinthians"
	//     },
	//     {
	//       "abbreviation": "1 John",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "1JN",
	//       "name": "1 John",
	//       "nameLong": "John’s First Letter"
	//     },
	//     {
	//       "abbreviation": "1 Peter",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "1PE",
	//       "name": "1 Peter",
	//       "nameLong": "Peter’s First Letter"
	//     },
	//     {
	//       "abbreviation": "1 Thessalonians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "1TH",
	//       "name": "1 Thessalonians",
	//       "nameLong": "Paul’s First Letter to the Thessalonians"
	//     },
	//     {
	//       "abbreviation": "1 Timothy",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "1TI",
	//       "name": "1 Timothy",
	//       "nameLong": "Paul’s First Letter to Timothy"
	//     },
	//     {
	//       "abbreviation": "2 Corinthians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "2CO",
	//       "name": "2 Corinthians",
	//       "nameLong": "Paul’s Second Letter to the Corinthians"
	//     },
	//     {
	//       "abbreviation": "2 John",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "2JN",
	//       "name": "2 John",
	//       "nameLong": "John’s Second Letter"
	//     },
	//     {
	//       "abbreviation": "2 Peter",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "2PE",
	//       "name": "2 Peter",
	//       "nameLong": "Peter’s Second Letter"
	//     },
	//     {
	//       "abbreviation": "2 Thessalonians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "2TH",
	//       "name": "2 Thessalonians",
	//       "nameLong": "Paul’s Second Letter to the Thessalonians"
	//     },
	//     {
	//       "abbreviation": "2 Timothy",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "2TI",
	//       "name": "2 Timothy",
	//       "nameLong": "Paul’s Second Letter to Timothy"
	//     },
	//     {
	//       "abbreviation": "3 John",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "3JN",
	//       "name": "3 John",
	//       "nameLong": "John’s Third Letter"
	//     },
	//     {
	//       "abbreviation": "Acts",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "ACT",
	//       "name": "Acts",
	//       "nameLong": "The Acts of the Apostles"
	//     },
	//     {
	//       "abbreviation": "Colossians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "COL",
	//       "name": "Colossians",
	//       "nameLong": "Paul’s Letter to the Colossians"
	//     },
	//     {
	//       "abbreviation": "Ephesians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "EPH",
	//       "name": "Ephesians",
	//       "nameLong": "Paul’s Letter to the Ephesians"
	//     },
	//     {
	//       "abbreviation": "Galatians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "GAL",
	//       "name": "Galatians",
	//       "nameLong": "Paul’s Letter to the Galatians"
	//     },
	//     {
	//       "abbreviation": "Hebrews",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "HEB",
	//       "name": "Hebrews",
	//       "nameLong": "The Letter to the Hebrews"
	//     },
	//     {
	//       "abbreviation": "James",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "JAS",
	//       "name": "James",
	//       "nameLong": "The Letter from James"
	//     },
	//     {
	//       "abbreviation": "John",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "JHN",
	//       "name": "John",
	//       "nameLong": "The Good News According to John"
	//     },
	//     {
	//       "abbreviation": "Jude",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "JUD",
	//       "name": "Jude",
	//       "nameLong": "The Letter from Jude"
	//     },
	//     {
	//       "abbreviation": "Luke",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "LUK",
	//       "name": "Luke",
	//       "nameLong": "The Good News According to Luke"
	//     },
	//     {
	//       "abbreviation": "Matthew",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "MAT",
	//       "name": "Matthew",
	//       "nameLong": "The Good News According to Matthew"
	//     },
	//     {
	//       "abbreviation": "Mark",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "MRK",
	//       "name": "Mark",
	//       "nameLong": "The Good News According to Mark"
	//     },
	//     {
	//       "abbreviation": "Philemon",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "PHM",
	//       "name": "Philemon",
	//       "nameLong": "Paul’s Letter to Philemon"
	//     },
	//     {
	//       "abbreviation": "Philippians",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "PHP",
	//       "name": "Philippians",
	//       "nameLong": "Paul’s Letter to the Philippians"
	//     },
	//     {
	//       "abbreviation": "Revelation",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "REV",
	//       "name": "Revelation",
	//       "nameLong": "The Revelation to John"
	//     },
	//     {
	//       "abbreviation": "Romans",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "ROM",
	//       "name": "Romans",
	//       "nameLong": "Paul’s Letter to the Romans"
	//     },
	//     {
	//       "abbreviation": "Titus",
	//       "bibleId": "105a06b6146d11e7-01",
	//       "id": "TIT",
	//       "name": "Titus",
	//       "nameLong": "Paul’s Letter to Titus"
	//     }
	//   ]
	// }
}

func ExampleGetAudioBibleBookById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.AudioBibleBookParams{} struct ref
		// to the call to GetAudioBibleBookById.
		audioBibleBookParams := params.AudioBibleBookParams{
			IncludeChapters: true,
		}
		// Here is an example of an API call
		webAudioBibleId := "105a06b6146d11e7-01"
		bookId := "3JN"
		audioBibleBook, audioBibleBookErr :=
			GetAudioBibleBookById(apiKey, webAudioBibleId, bookId, &audioBibleBookParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if audioBibleBookErr != nil {
			fmt.Println("Do error handling for failing to get audio bible book here")
		} else {
			prettyAudioBibleBook, prettifyErr := utils.Prettify(audioBibleBook)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyAudioBibleBook)
			}
		}
	}
	// Output:
	// {
	//   "data": {
	//     "abbreviation": "3 John",
	//     "bibleId": "105a06b6146d11e7-01",
	//     "chapters": [
	//       {
	//         "bibleId": "105a06b6146d11e7-01",
	//         "bookId": "3JN",
	//         "id": "3JN.1",
	//         "number": "1",
	//         "position": 61
	//       }
	//     ],
	//     "id": "3JN",
	//     "name": "3 John",
	//     "nameLong": "John’s Third Letter"
	//   }
	// }
}

func ExampleGetBibleChapters() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		kjvBibleId := "de4e12af7f28f599-02"
		bookOf3JohnId := "3JN"
		bibleChapters, bibleChaptersErr := GetBibleChapters(apiKey, kjvBibleId, bookOf3JohnId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleChaptersErr != nil {
			fmt.Println("Do error handling for failing to get bible chapters here")
		} else {
			prettyBibleChapters, prettifyErr := utils.Prettify(bibleChapters)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBibleChapters)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "3JN",
	//       "id": "3JN.intro",
	//       "number": "intro",
	//       "reference": "3 John"
	//     },
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "3JN",
	//       "id": "3JN.1",
	//       "number": "1",
	//       "reference": "3 John 1"
	//     }
	//   ]
	// }
}

func ExampleGetBibleChapterById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleChapterParams{} struct ref
		// to the call to GetBibleChapterById.
		bibleChapterParams := params.BibleChapterParams{
			ContentType:           "text", // choices are html, json, and text (html is default)
			IncludeNotes:          false,
			IncludeTitles:         false,
			IncludeChapterNumbers: true,
			IncludeVerseNumbers:   false,
			IncludeVerseSpans:     false,
			Parallels:             nil,
		}
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		psalmChapter117Id := "PSA.117"
		bibleChapter, bibleChapterErr := GetBibleChapterById(apiKey, kjvBibleId, psalmChapter117Id, &bibleChapterParams)
		if bibleChapterErr != nil {
			fmt.Println("Do error handling for failing to get bible chapter here")
		} else {
			// Here is an example of getting a single field from the response.
			// This is required for duplicability of testing, as the metadata
			// changes with every single API call.
			bibleChapterData, getJsonFieldErr := utils.GetJsonField(bibleChapter, "data")
			// Here is some boilerplate for handling errors and pretty-printing
			// Note: The pretty-printing is just to make the output readable.
			// You may not need to do this in your own production environment.
			if getJsonFieldErr != nil {
				fmt.Println("Do error handling for failing to get \"data\" field from JSON here")
			} else {
				prettyBibleChapterData, prettifyErr := utils.Prettify(bibleChapterData)
				if prettifyErr != nil {
					fmt.Println("Do error handling for failing to make pretty JSON here")
				} else {
					fmt.Println(prettyBibleChapterData)
				}
			}
		}
	}
	// Output:
	// {
	//   "bibleId": "de4e12af7f28f599-02",
	//   "bookId": "PSA",
	//   "content": "117\n   [1] O praise the LORD, all ye nations: praise him, all ye people.\n   [2] For his merciful kindness is great toward us: and the truth of the LORD endureth for ever. Praise ye the LORD.\n",
	//   "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
	//   "id": "PSA.117",
	//   "next": {
	//     "bookId": "PSA",
	//     "id": "PSA.118",
	//     "number": "118"
	//   },
	//   "number": "117",
	//   "previous": {
	//     "bookId": "PSA",
	//     "id": "PSA.116",
	//     "number": "116"
	//   },
	//   "reference": "Psalms 117",
	//   "verseCount": 2
	// }
}

func ExampleGetAudioBibleChapters() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		webAudioBibleId := "105a06b6146d11e7-01"
		bookId := "3JN"
		audioBibleChapters, audioBibleChaptersErr := GetAudioBibleChapters(apiKey, webAudioBibleId, bookId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if audioBibleChaptersErr != nil {
			fmt.Println("Do error handling for failing to get audio bible chapters here")
		} else {
			prettyBibleChaptersData, prettifyErr := utils.Prettify(audioBibleChapters)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBibleChaptersData)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "bibleId": "105a06b6146d11e7-01",
	//       "bookId": "3JN",
	//       "id": "3JN.1",
	//       "number": "1",
	//       "reference": "3 John 1"
	//     }
	//   ]
	// }
}

func ExampleGetAudioBibleChapterById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		webAudioBibleId := "105a06b6146d11e7-01"
		chapterId := "JHN.1"
		audioBibleChapter, audioBibleChapterErr := GetAudioBibleChapterById(apiKey, webAudioBibleId, chapterId)
		if audioBibleChapterErr != nil {
			fmt.Println("Do error handling for failing to get audio bible chapters here")
		} else {
			// Here is an example of getting a single field from the response.
			// This is required for duplicability of testing, as the metadata
			// changes with every single API call.
			audioBibleChapterData, getJsonFieldErr := utils.GetJsonField(audioBibleChapter, "data")
			// Here is some boilerplate for handling errors and pretty-printing
			// Note: The pretty-printing is just to make the output readable.
			// You may not need to do this in your own production environment.
			if getJsonFieldErr != nil {
				fmt.Println("Do error handling for failing to get \"data\" field from JSON here")
			} else {
				audioBibleNextChapter, getJsonField2Err := utils.GetJsonField(audioBibleChapterData, "next")
				if getJsonField2Err != nil {
					fmt.Println("Do error handling for failing to get \"next\" field from JSON here")
				} else {
					prettyAudioBibleChapterData, prettifyErr := utils.Prettify(audioBibleNextChapter)
					if prettifyErr != nil {
						fmt.Println("Do error handling for failing to make pretty JSON here")
					} else {
						fmt.Println(prettyAudioBibleChapterData)
					}
				}
			}
		}
	}
	// Output:
	// {
	//   "bookId": "JHN",
	//   "id": "JHN.2",
	//   "number": "2"
	// }
}

func ExampleGetBibleBookSections() {
	// NOTE: As of 2023-09-15, getting bible book sections
	// is experimental on the part of scripture.api.bible
	// Thus, the output of this is an error message.
	// Once we become aware of this API either working or getting deprecated
	// We will modify or remove this code accordingly.
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		bookId := "3JN"
		bibleBookSections, bibleBookSectionsErr := GetBibleBookSections(apiKey, kjvBibleId, bookId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleBookSectionsErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			prettyBookSections, prettifyErr := utils.Prettify(bibleBookSections)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyBookSections)
			}
		}
	}
	// Output:
	// {
	//   "error": "Not Found",
	//   "message": "Sections not found for bookId",
	//   "statusCode": 404
	// }
}

func ExampleGetBibleChapterSections() {
	// NOTE: As of 2023-09-15, getting bible chapter sections
	// is experimental on the part of scripture.api.bible
	// Thus, the output of this is an error message.
	// Once we become aware of this API either working or getting deprecated
	// We will modify or remove this code accordingly.
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		chapterId := "3JN.1"
		bibleChapterSections, bibleChapterSectionsErr := GetBibleChapterSections(apiKey, kjvBibleId, chapterId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleChapterSectionsErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			prettyChapterSections, prettifyErr := utils.Prettify(bibleChapterSections)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyChapterSections)
			}
		}
	}
	// Output:
	// {
	//   "error": "Not Found",
	//   "message": "Sections not found for bookId",
	//   "statusCode": 404
	// }
}

func ExampleGetBibleSectionById() {
	// NOTE: As of 2023-09-15, getting bible sections
	// is experimental on the part of scripture.api.bible
	// Thus, the output of this is an error message.
	// Once we become aware of this API either working or getting deprecated
	// We will modify or remove this code accordingly.
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleSectionParams{} struct ref
		// to the call to GetBibleSectionById.
		kjvBibleId := "de4e12af7f28f599-02"
		webBibleId := "32664dc3288a28df-03"
		asvBibleId := "685d1470fe4d5c3b-01"
		bibleSectionParams := params.BibleSectionParams{
			ContentType:           "text",
			IncludeNotes:          false,
			IncludeTitles:         false,
			IncludeChapterNumbers: false,
			IncludeVerseNumbers:   false,
			IncludeVerseSpans:     false,
			Parallels:             []string{asvBibleId, webBibleId},
		}
		// Here is an example of an API call
		sectionId := "abc123"
		bibleSection, bibleSectionErr := GetBibleSectionById(apiKey, kjvBibleId, sectionId, &bibleSectionParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleSectionErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			prettySection, prettifyErr := utils.Prettify(bibleSection)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettySection)
			}
		}
	}
	// Output:
	// {
	//   "error": "Bad Request",
	//   "message": "Invalid request params input",
	//   "statusCode": 400
	// }
}

func ExampleGetBiblePassage() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BiblePassageParams{} struct ref
		// to the call to GetBiblePassage.
		webBibleId := "32664dc3288a28df-03"
		asvBibleId := "685d1470fe4d5c3b-01"
		biblePassageParams := params.BiblePassageParams{
			ContentType:           "text",
			IncludeNotes:          false,
			IncludeTitles:         false,
			IncludeChapterNumbers: false,
			IncludeVerseNumbers:   false,
			IncludeVerseSpans:     false,
			Parallels:             []string{asvBibleId, webBibleId},
			UseOrgId:              false,
		}
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		passage := "2TI.3.16-2TI.3.17"
		biblePassage, biblePassageErr := GetBiblePassage(apiKey, kjvBibleId, passage, &biblePassageParams)
		// Here is an example of getting a single field from the response.
		// This is required for duplicability of testing, as the metadata
		// changes with every single API call.
		biblePassageData, getJsonFieldErr := utils.GetJsonField(biblePassage, "data")
		if biblePassageErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			// Here is some boilerplate for handling errors and pretty-printing
			// Note: The pretty-printing is just to make the output readable.
			// You may not need to do this in your own production environment.
			if getJsonFieldErr != nil {
				fmt.Println("Do error handling for failing to get \"data\" field from JSON here")
			} else {
				prettyPassageData, prettifyErr := utils.Prettify(biblePassageData)
				if prettifyErr != nil {
					fmt.Println("Do error handling for failing to make pretty JSON here")
				} else {
					fmt.Println(prettyPassageData)
				}
			}
		}
	}
	// Output:
	// {
	//   "bibleId": "de4e12af7f28f599-02",
	//   "bookId": "2TI",
	//   "chapterIds": [
	//     "2TI.3"
	//   ],
	//   "content": "     [16] All scripture is given by inspiration of God, and is profitable for doctrine, for reproof, for correction, for instruction in righteousness:  [17] That the man of God may be perfect, throughly furnished unto all good works.\n",
	//   "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
	//   "id": "2TI.3.16-2TI.3.17",
	//   "orgId": "2TI.3.16-2TI.3.17",
	//   "parallels": [
	//     {
	//       "bibleId": "685d1470fe4d5c3b-01",
	//       "bookId": "2TI",
	//       "chapterIds": [
	//         "2TI.3"
	//       ],
	//       "content": "     [16] Every scripture is inspired of God, and profitable for teaching, for reproof, for correction, for instruction which is in righteousness:  [17] that the man of God may be complete, furnished completely unto every good work.\n",
	//       "copyright": "PUBLIC DOMAIN",
	//       "id": "2TI.3.16-2TI.3.17",
	//       "orgId": "2TI.3.16-2TI.3.17",
	//       "reference": "2 Timothy 3:16-17",
	//       "verseCount": 2
	//     },
	//     {
	//       "bibleId": "32664dc3288a28df-03",
	//       "bookId": "2TI",
	//       "chapterIds": [
	//         "2TI.3"
	//       ],
	//       "content": "     [16] Every Scripture is God-breathed and profitable for teaching, for reproof, for correction, and for instruction in righteousness,  [17] that each person who belongs to God may be complete, thoroughly equipped for every good work.\n",
	//       "copyright": "PUBLIC DOMAIN (not copyrighted)",
	//       "id": "2TI.3.16-2TI.3.17",
	//       "orgId": "2TI.3.16-2TI.3.17",
	//       "reference": "2 Timothy 3:16-17",
	//       "verseCount": 2
	//     }
	//   ],
	//   "reference": "2 Timothy 3:16-17",
	//   "verseCount": 2
	// }
}

func ExampleGetBibleChapterVerses() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		chapterId := "PSA.117"
		bibleChapterVerses, bibleChapterVersesErr := GetBibleChapterVerses(apiKey, kjvBibleId, chapterId)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleChapterVersesErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			prettyChapterVerses, prettifyErr := utils.Prettify(bibleChapterVerses)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to make pretty JSON here")
			} else {
				fmt.Println(prettyChapterVerses)
			}
		}
	}
	// Output:
	// {
	//   "data": [
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "PSA",
	//       "chapterId": "PSA.117",
	//       "id": "PSA.117.1",
	//       "orgId": "PSA.117.1",
	//       "reference": "Psalms 117:1"
	//     },
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "PSA",
	//       "chapterId": "PSA.117",
	//       "id": "PSA.117.2",
	//       "orgId": "PSA.117.2",
	//       "reference": "Psalms 117:2"
	//     }
	//   ]
	// }
}

func ExampleGetBibleVerseById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleVerseParams{} struct ref
		// to the call to GetBibleVerseById.
		kjvBibleId := "de4e12af7f28f599-02"
		webBibleId := "32664dc3288a28df-03"
		asvBibleId := "685d1470fe4d5c3b-01"
		bibleVerseParams := params.BibleVerseParams{
			ContentType:           "text",
			IncludeNotes:          false,
			IncludeTitles:         false,
			IncludeChapterNumbers: false,
			IncludeVerseNumbers:   false,
			IncludeVerseSpans:     false,
			Parallels:             []string{asvBibleId, webBibleId},
			UseOrgId:              false,
		}
		verse := "JHN.3.16"
		bibleVerse, bibleVerseErr := GetBibleVerseById(apiKey, kjvBibleId, verse, &bibleVerseParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleVerseErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			bibleVerseData, getJsonFieldErr := utils.GetJsonField(bibleVerse, "data")
			if bibleVerseErr != nil {
				fmt.Println("Do error handling for failing to get bible book sections here")
			} else {
				// Here is some boilerplate for handling errors and pretty-printing
				// Note: The pretty-printing is just to make the output readable.
				// You may not need to do this in your own production environment.
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
	}
	// Output:
	// {
	//   "bibleId": "de4e12af7f28f599-02",
	//   "bookId": "JHN",
	//   "chapterId": "JHN.3",
	//   "content": "     [16] ¶ For God so loved the world, that he gave his only begotten Son, that whosoever believeth in him should not perish, but have everlasting life. \n",
	//   "copyright": "PUBLIC DOMAIN except in the United Kingdom, where a Crown Copyright applies to printing the KJV. See http://www.cambridge.org/about-us/who-we-are/queens-printers-patent",
	//   "id": "JHN.3.16",
	//   "next": {
	//     "id": "JHN.3.17",
	//     "number": "17"
	//   },
	//   "orgId": "JHN.3.16",
	//   "parallels": [
	//     {
	//       "bibleId": "32664dc3288a28df-03",
	//       "bookId": "JHN",
	//       "chapterId": "JHN.3",
	//       "content": "     [16] For God so loved the world, that he gave his only born Son, that whoever believes in him should not perish, but have eternal life.  \n",
	//       "copyright": "PUBLIC DOMAIN (not copyrighted)",
	//       "id": "JHN.3.16",
	//       "orgId": "JHN.3.16",
	//       "reference": "John 3:16",
	//       "verseCount": 1
	//     }
	//   ],
	//   "previous": {
	//     "id": "JHN.3.15",
	//     "number": "15"
	//   },
	//   "reference": "John 3:16",
	//   "verseCount": 1
	// }
}

func ExampleGetBibleSearchResults() {
	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleSearchParams{} struct ref
		// to the call to GetBibleSearchResults.
		bibleSearchParams := params.BibleSearchParams{
			Query:     "Love",
			Limit:     3,           // Default value is 10
			Offset:    5,           // Pagination
			Sort:      "canonical", // Supported values are relevance (default), canonical and reverse-canonical
			Range:     "MAT.1-REV.22",
			Fuzziness: "0", // Default value is AUTO. This accounts for difference in spelling
		}
		// Here is an example of an API call
		kjvBibleId := "de4e12af7f28f599-02"
		bibleSearchResult, bibleSearchErr := GetBibleSearchResults(apiKey, kjvBibleId, &bibleSearchParams)
		// Here is some boilerplate for handling errors and pretty-printing
		// Note: The pretty-printing is just to make the output readable.
		// You may not need to do this in your own production environment.
		if bibleSearchErr != nil {
			fmt.Println("Do error handling for failing to get bible book sections here")
		} else {
			// Here is an example of getting a single field from the response.
			// This is required for duplicability of testing, as the metadata
			// changes with every single API call.
			bibleSearchResultData, getJsonFieldErr := utils.GetJsonField(bibleSearchResult, "data")
			// Here is some boilerplate for handling errors and pretty-printing
			// Note: The pretty-printing is just to make the output readable.
			// You may not need to do this in your own production environment.
			if getJsonFieldErr != nil {
				fmt.Println("Do error handling for failing to get \"data\" field from JSON here")
			} else {
				prettyBibleSearchResultData, prettifyErr := utils.Prettify(bibleSearchResultData)
				if prettifyErr != nil {
					fmt.Println("Do error handling for failing to make pretty JSON here")
				} else {
					fmt.Println(prettyBibleSearchResultData)
				}
			}
		}
	}
	// Output:
	// {
	//   "limit": 3,
	//   "offset": 5,
	//   "query": "Love",
	//   "total": 183,
	//   "verseCount": 3,
	//   "verses": [
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "MAT",
	//       "chapterId": "MAT.19",
	//       "id": "MAT.19.19",
	//       "orgId": "MAT.19.19",
	//       "reference": "Matthew 19:19",
	//       "text": "Honour thy father and thy  mother: and, Thou shalt love thy neighbour as thyself."
	//     },
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "MAT",
	//       "chapterId": "MAT.22",
	//       "id": "MAT.22.37",
	//       "orgId": "MAT.22.37",
	//       "reference": "Matthew 22:37",
	//       "text": "Jesus said unto him, Thou shalt love the Lord thy God with all thy heart, and with all thy soul, and with all thy mind."
	//     },
	//     {
	//       "bibleId": "de4e12af7f28f599-02",
	//       "bookId": "MAT",
	//       "chapterId": "MAT.22",
	//       "id": "MAT.22.39",
	//       "orgId": "MAT.22.39",
	//       "reference": "Matthew 22:39",
	//       "text": "And the second is  like unto it, Thou shalt love thy neighbour as thyself."
	//     }
	//   ]
	// }
}
