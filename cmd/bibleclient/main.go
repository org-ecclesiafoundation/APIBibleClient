package main

import (
	"ecclesiafoundation.org/APIBibleClient/pkg/client"
	"ecclesiafoundation.org/APIBibleClient/pkg/client/params"
	"ecclesiafoundation.org/APIBibleClient/pkg/utils"
	"fmt"
	"os"
)

func main() {
	fmt.Println("args:", os.Args)
	kjvBibleId := "de4e12af7f28f599-02"
	fmt.Printf("KJV Bible ID: %s\n", kjvBibleId)
	webAudioBibleId := "105a06b6146d11e7-01"
	bookOfGenesisId := "GEN"
	bookOfJohnId := "JHN"

	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println(apiKeyErr)
	} else {
		fmt.Println("Test `GetBibles`")
		bibles, biblesErr := client.GetBibles(apiKey, &params.BiblesParams{Abbreviation: "kjv"})
		printAndPrettyPrintJsonString(bibles, biblesErr)

		fmt.Println("Test `GetBibleById`")
		kjv, kjvErr := client.GetBibleById(apiKey, kjvBibleId)
		printAndPrettyPrintJsonString(kjv, kjvErr)

		fmt.Println("Test `GetAudioBibles`")
		audioBibles, audioBiblesErr := client.GetAudioBibles(apiKey, &params.AudioBiblesParams{})
		printAndPrettyPrintJsonString(audioBibles, audioBiblesErr)

		fmt.Println("Test `GetAudioBibleById`")
		audioWEB, audioWEBErr := client.GetAudioBibleById(apiKey, webAudioBibleId)
		printAndPrettyPrintJsonString(audioWEB, audioWEBErr)

		fmt.Println("Test `GetBibleBooks`")
		kjvBooks, kjvBooksErr := client.GetBibleBooks(apiKey, kjvBibleId, &params.BibleBooksParams{})
		printAndPrettyPrintJsonString(kjvBooks, kjvBooksErr)

		fmt.Println("Test `GetBibleBookById`")
		kjvBook, kjvBookErr := client.GetBibleBookById(apiKey, kjvBibleId, bookOfGenesisId, &params.BibleBookParams{})
		printAndPrettyPrintJsonString(kjvBook, kjvBookErr)

		fmt.Println("Test `GetAudioBibleBooks`")
		audioBibleBooks, audioBibleBooksErr := client.GetAudioBibleBooks(apiKey, webAudioBibleId, &params.AudioBibleBooksParams{})
		printAndPrettyPrintJsonString(audioBibleBooks, audioBibleBooksErr)

		fmt.Println("Test `GetAudioBibleBookById`")
		audioBibleBook, audioBibleBookErr := client.GetAudioBibleBookById(apiKey, webAudioBibleId, bookOfJohnId, &params.AudioBibleBookParams{})
		printAndPrettyPrintJsonString(audioBibleBook, audioBibleBookErr)
	}
}

func printAndPrettyPrintJsonString(json string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json)
		cleanJson, cleanErr := utils.Prettify(json)
		if cleanErr != nil {
			fmt.Println(cleanErr)
		} else {
			fmt.Println(cleanJson)
		}
	}
}
