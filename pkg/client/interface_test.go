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
	var output string
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.BibleParams{} struct
		// to the call to GetBibles.
		biblesParams := params.BiblesParams{
			Language:           "eng",
			Abbreviation:       "kjv",
			Name:               "King James",
			Ids:                []string{"de4e12af7f28f599-01"},
			IncludeFullDetails: true,
		}
		bibles, biblesErr := GetBibles(apiKey, &biblesParams)
		// Here is some boilerplate for handling errors and pretty-printing
		if biblesErr != nil {
			fmt.Println("Do error handling for failing to get bibles here")
		} else {
			prettyBibles, prettifyErr := utils.Prettify(bibles)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to print bibles JSON here")
			} else {
				output = prettyBibles
			}
		}
	}
	fmt.Println(output)
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
	//       "description": "Ecumenical",
	//       "descriptionLocal": "Ecumenical",
	//       "id": "de4e12af7f28f599-01",
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
	//       "updatedAt": "2023-05-03T21:21:08.000Z"
	//     }
	//   ]
	// }
}

func ExampleGetBibleById() {
	apiKey, apiKeyErr := utils.GetApiKey()
	kjvBibleId := "de4e12af7f28f599-02"
	var output string
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of an API call
		bible, bibleErr := GetBibleById(apiKey, kjvBibleId)
		// Here is some boilerplate for handling errors and pretty-printing
		if bibleErr != nil {
			fmt.Println("Do error handling for failing to get bible here")
		} else {
			prettyBible, prettifyErr := utils.Prettify(bible)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to print bible JSON here")
			} else {
				output = prettyBible
			}
		}
	}
	fmt.Println(output)
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
	var output string
	if apiKeyErr != nil {
		fmt.Println("Failed to get API key.\n" +
			"Please set the environment variable SCRIPTURE_API_BIBLE_KEY to the appropriate value")
	} else {
		// Here is an example of all the possible API
		// parameters used.
		// You may use any subset of these parameters,
		// including passing in a blank params.AudioBiblesParams{} struct
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
		if audioBiblesErr != nil {
			fmt.Println("Do error handling for failing to get audio bibles here")
		} else {
			prettyAudioBible, prettifyErr := utils.Prettify(audioBibles)
			if prettifyErr != nil {
				fmt.Println("Do error handling for failing to print audio bible JSON here")
			} else {
				output = prettyAudioBible
			}
		}
	}
	fmt.Println(output)
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
