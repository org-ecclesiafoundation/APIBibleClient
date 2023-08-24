package main

import (
	"ecclesiafoundation.org/scripture-api-bible-client/pkg/client"
	"ecclesiafoundation.org/scripture-api-bible-client/pkg/client/params"
	"fmt"
	"os"
)

func main() {
	fmt.Println("args:", os.Args)
	kjvBibleId := "de4e12af7f28f599-02"
	fmt.Printf("KJV Bible ID: %s\n", kjvBibleId)

	apiKey, apiKeyErr := client.RetrieveApiKey()
	if apiKeyErr != nil {
		fmt.Println(apiKeyErr)
	} else {
		bibles, biblesErr := client.RetrieveBibles(apiKey, params.BiblesParams{Abbreviation: "kjv"})
		if biblesErr != nil {
			fmt.Println(biblesErr)
		} else {
			fmt.Println(bibles)
			cleanBibles, cleanErr := client.Prettify(bibles)
			if cleanErr != nil {
				fmt.Println(cleanErr)
			} else {
				fmt.Println(cleanBibles)
			}
		}
	}
}
