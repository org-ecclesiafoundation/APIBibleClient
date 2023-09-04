package client

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"ecclesiafoundation.org/APIBibleClient/pkg/utils"
	"fmt"
	"testing"
)

const (
	API_KEY_ENV_VAR = "SCRIPTURE_API_BIBLE_KEY"
)

func TestStub(t *testing.T) {}

func ExampleGetBibles() {
	apiKey, apiKeyErr := utils.GetApiKey()
	var output string
	if apiKeyErr != nil {
		fmt.Printf("Failed to get API key.\n"+
			"Please set the environment variable %s to the appropriate value",
			API_KEY_ENV_VAR)
	} else {
		// Here is an example of an API call with all the possible
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
