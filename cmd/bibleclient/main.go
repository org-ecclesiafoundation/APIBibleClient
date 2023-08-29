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
	bookOfJohnId := "JHN"
	johnChapter1Id := "JHN.1"
	secondTimPassage := "2TI.3.16-2TI.3.17"
	john3Verse16 := "JHN.3.16"
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
		kjvBook, kjvBookErr := client.GetBibleBookById(apiKey, kjvBibleId, bookOfJohnId, &params.BibleBookParams{})
		printAndPrettyPrintJsonString(kjvBook, kjvBookErr)

		fmt.Println("Test `GetAudioBibleBooks`")
		audioBibleBooks, audioBibleBooksErr := client.GetAudioBibleBooks(apiKey, webAudioBibleId, &params.AudioBibleBooksParams{})
		printAndPrettyPrintJsonString(audioBibleBooks, audioBibleBooksErr)

		fmt.Println("Test `GetAudioBibleBookById`")
		audioBibleBook, audioBibleBookErr := client.GetAudioBibleBookById(apiKey, webAudioBibleId, bookOfJohnId, &params.AudioBibleBookParams{})
		printAndPrettyPrintJsonString(audioBibleBook, audioBibleBookErr)

		fmt.Println("Test `GetBibleChapters`")
		bibleChapters, bibleChaptersErr := client.GetBibleChapters(apiKey, kjvBibleId, bookOfJohnId)
		printAndPrettyPrintJsonString(bibleChapters, bibleChaptersErr)

		fmt.Println("Test `GetBibleChapterById`")
		bibleChapter, bibleChapterErr := client.GetBibleChapterById(apiKey, kjvBibleId, johnChapter1Id, &params.BibleChapterParams{})
		printAndPrettyPrintJsonString(bibleChapter, bibleChapterErr)

		fmt.Println("Test `GetAudioBibleChapters`")
		audioBibleChapters, audioBibleChaptersErr := client.GetAudioBibleChapters(apiKey, webAudioBibleId, bookOfJohnId)
		printAndPrettyPrintJsonString(audioBibleChapters, audioBibleChaptersErr)

		fmt.Println("Test `GetAudioBibleChapterById`")
		audioBibleChapter, audioBibleChapterErr := client.GetAudioBibleChapterById(apiKey, webAudioBibleId, johnChapter1Id)
		printAndPrettyPrintJsonString(audioBibleChapter, audioBibleChapterErr)

		// NOTE: There are currently not any sections available through the API, so querying for them will return a "Not Found" error.
		fmt.Println("Test `GetBibleBookSections`")
		bibleBookSections, bibleBookSectionsErr := client.GetBibleBookSections(apiKey, kjvBibleId, bookOfJohnId)
		printAndPrettyPrintJsonString(bibleBookSections, bibleBookSectionsErr)

		fmt.Println("Test `GetBibleChapterSections`")
		bibleChapterSections, bibleChapterSectionsErr := client.GetBibleChapterSections(apiKey, kjvBibleId, johnChapter1Id)
		printAndPrettyPrintJsonString(bibleChapterSections, bibleChapterSectionsErr)

		// fmt.Println("Test `GetBibleSectionById`")
		// bibleSection, bibleSectionErr := client.GetBibleSectionById(apiKey, kjvBibleId, {SECTION_ID}, &params.BibleChapterParams{})
		// printAndPrettyPrintJsonString(bibleSection, bibleSectionErr)

		fmt.Println("Test `GetBiblePassage`")
		biblePassage, biblePassageErr := client.GetBiblePassage(apiKey, kjvBibleId, secondTimPassage, &params.BiblePassageParams{})
		printAndPrettyPrintJsonString(biblePassage, biblePassageErr)

		fmt.Println("Test `GetBibleChapterVerses`")
		bibleVerses, bibleVersesErr := client.GetBibleChapterVerses(apiKey, kjvBibleId, johnChapter1Id)
		printAndPrettyPrintJsonString(bibleVerses, bibleVersesErr)

		fmt.Println("Test `GetBibleVerseById`")
		bibleVerse, bibleVerseErr := client.GetBibleVerseById(apiKey, kjvBibleId, john3Verse16, &params.BibleVerseParams{})
		printAndPrettyPrintJsonString(bibleVerse, bibleVerseErr)

		fmt.Println("Test `GetBibleSearchResults`")
		bibleSearch, bibleSearchErr := client.GetBibleSearchResults(apiKey, kjvBibleId, &params.BibleSearchParams{Query: "Love"})
		printAndPrettyPrintJsonString(bibleSearch, bibleSearchErr)
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
