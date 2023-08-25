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

	apiKey, apiKeyErr := utils.GetApiKey()
	if apiKeyErr != nil {
		fmt.Println(apiKeyErr)
	} else {
		bibles, biblesErr := client.GetBibles(apiKey, params.BiblesParams{Abbreviation: "kjv"})
		if biblesErr != nil {
			fmt.Println(biblesErr)
		} else {
			fmt.Println(bibles)
			cleanBibles, cleanErr := utils.Prettify(bibles)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanBibles)
			}
		}
		kjv, kjvErr := client.GetBibleById(apiKey, kjvBibleId)
		if kjvErr != nil {
			fmt.Println(kjvErr)
		} else {
			fmt.Println(kjv)
			cleanKjv, cleanErr := utils.Prettify(kjv)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanKjv)
			}
		}
		audioBibles, audioBiblesErr := client.GetAudioBibles(apiKey, params.AudioBiblesParams{})
		if audioBiblesErr != nil {
			fmt.Println(audioBiblesErr)
		} else {
			fmt.Println(audioBibles)
			cleanAudioKjv, cleanErr := utils.Prettify(audioBibles)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanAudioKjv)
			}
		}
		audioWEB, audioWEBErr := client.GetAudioBibleById(apiKey, webAudioBibleId)
		if audioWEBErr != nil {
			fmt.Println(audioWEBErr)
		} else {
			fmt.Println(audioWEB)
			cleanAudioWEB, cleanErr := utils.Prettify(audioWEB)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanAudioWEB)
			}
		}
		kjvBooks, kjvBooksErr := client.GetBibleBooks(apiKey, kjvBibleId, params.BibleBooksParams{})
		if kjvBooksErr != nil {
			fmt.Println(kjvBooksErr)
		} else {
			fmt.Println(kjvBooks)
			cleanKjvBooks, cleanErr := utils.Prettify(kjvBooks)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanKjvBooks)
			}
		}
	}
}
