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
	genesisChapter1Id := "GEN.1"
	johnChapter1Id := "JHN.1"
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

		fmt.Println("Test `GetBibleChapters`")
		bibleChapters, bibleChaptersErr := client.GetBibleChapters(apiKey, kjvBibleId, bookOfGenesisId)
		printAndPrettyPrintJsonString(bibleChapters, bibleChaptersErr)

		fmt.Println("Test `GetBibleChapterById`")
		bibleChapter, bibleChapterErr := client.GetBibleChapterById(apiKey, kjvBibleId, genesisChapter1Id, &params.BibleChapterParams{})
		printAndPrettyPrintJsonString(bibleChapter, bibleChapterErr)

		fmt.Println("Test `GetAudioBibleChapters`")
		audioBibleChapters, audioBibleChaptersErr := client.GetAudioBibleChapters(apiKey, webAudioBibleId, bookOfJohnId)
		printAndPrettyPrintJsonString(audioBibleChapters, audioBibleChaptersErr)

		fmt.Println("Test `GetAudioBibleChapterById`")
		audioBibleChapter, audioBibleChapterErr := client.GetAudioBibleChapterById(apiKey, webAudioBibleId, johnChapter1Id)
		printAndPrettyPrintJsonString(audioBibleChapter, audioBibleChapterErr)

		// NOTE: There are currently not any sections available through the API, so querying for them will return a "Not Found" error.
		fmt.Println("Test `GetBibleBookSections`")
		bibleBookSections, bibleBookSectionsErr := client.GetBibleBookSections(apiKey, kjvBibleId, "JHN")
		printAndPrettyPrintJsonString(bibleBookSections, bibleBookSectionsErr)

		fmt.Println("Test `GetBibleChapterSections`")
		bibleChapterSections, bibleChapterSectionsErr := client.GetBibleChapterSections(apiKey, kjvBibleId, "COL.1")
		printAndPrettyPrintJsonString(bibleChapterSections, bibleChapterSectionsErr)

		// fmt.Println("Test `GetBibleSectionById`")
		// bibleSection, bibleSectionErr := client.GetBibleSectionById(apiKey, kjvBibleId, {SECTION_ID}, &params.BibleChapterParams{})
		// printAndPrettyPrintJsonString(bibleSection, bibleSectionErr)
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
