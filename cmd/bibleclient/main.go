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
		audioKjv, audioKjvErr := client.GetAudioBibles(apiKey, params.AudioBiblesParams{})
		if audioKjvErr != nil {
			fmt.Println(audioKjvErr)
		} else {
			fmt.Println(audioKjv)
			cleanAudioKjv, cleanErr := utils.Prettify(audioKjv)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanAudioKjv)
			}
		}
	}
}
