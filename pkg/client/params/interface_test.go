package params

import (
	"fmt"
	"testing"
)

func TestStub(t *testing.T) {}

func ExampleBiblesParams_ProduceQueryParameters() {
	kjvBibleId := "de4e12af7f28f599-02"
	webBibleId := "32664dc3288a28df-03"
	asvBibleId := "685d1470fe4d5c3b-01"
	biblesParams := BiblesParams{
		Language:           "eng",
		Abbreviation:       "kjv",
		Name:               "King James",
		Ids:                []string{kjvBibleId, webBibleId, asvBibleId},
		IncludeFullDetails: true,
	}
	queryParameters := biblesParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// abbreviation=kjv&ids=de4e12af7f28f599-02%2C32664dc3288a28df-03%2C685d1470fe4d5c3b-01&include-full-details=true&language=eng&name=King+James
}

func ExampleAudioBiblesParams_ProduceQueryParameters() {
	worldEnglishBibleID := "9879dbb7cfe39e4d-01"
	audioBiblesParams := AudioBiblesParams{
		Language:           "eng",
		Abbreviation:       "WEB13",
		Name:               "World English",
		RelatedTextBibleId: worldEnglishBibleID,
		IncludeFullDetails: true,
	}
	queryParameters := audioBiblesParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// abbreviation=WEB13&bibleId=9879dbb7cfe39e4d-01&include-full-details=true&language=eng&name=World+English
}

func ExampleBibleBooksParams_ProduceQueryParameters() {
	bibleBooksParams := BibleBooksParams{
		IncludeChapters:            true,
		IncludeChaptersAndSections: true,
	}
	queryParameters := bibleBooksParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// include-chapters=true&include-chapters-and-sections=true
}

func ExampleBibleBookParams_ProduceQueryParameters() {
	bibleBookParams := BibleBookParams{
		IncludeChapters: true,
	}
	queryParameters := bibleBookParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// include-chapters=true
}

func ExampleAudioBibleBooksParams_ProduceQueryParameters() {
	audioBibleBooksParams := AudioBibleBooksParams{
		IncludeChapters:            true,
		IncludeChaptersAndSections: true,
	}
	queryParameters := audioBibleBooksParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// include-chapters=true&include-chapters-and-sections=true
}

func ExampleAudioBibleBookParams_ProduceQueryParameters() {
	audioBibleBookParams := AudioBibleBookParams{
		IncludeChapters: true,
	}
	queryParameters := audioBibleBookParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// include-chapters=true
}

func ExampleBibleChapterParams_ProduceQueryParameters() {
	asvBibleId := "685d1470fe4d5c3b-01"
	webBibleId := "32664dc3288a28df-03"
	bibleChapterParams := BibleChapterParams{
		ContentType:           "text", // choices are html, json, and text (html is default)
		IncludeNotes:          true,
		IncludeTitles:         true,
		IncludeChapterNumbers: true,
		IncludeVerseNumbers:   true,
		IncludeVerseSpans:     true,
		Parallels:             []string{asvBibleId, webBibleId},
	}
	queryParameters := bibleChapterParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// content-type=text&include-chapter-numbers=true&include-notes=true&include-titles=true&include-verse-numbers=true&include-verse-spans=true&parallels=685d1470fe4d5c3b-01%2C32664dc3288a28df-03
}

func ExampleBibleSectionParams_ProduceQueryParameters() {
	asvBibleId := "685d1470fe4d5c3b-01"
	webBibleId := "32664dc3288a28df-03"
	bibleSectionParams := BibleSectionParams{
		ContentType:           "text",
		IncludeNotes:          true,
		IncludeTitles:         true,
		IncludeChapterNumbers: true,
		IncludeVerseNumbers:   true,
		IncludeVerseSpans:     true,
		Parallels:             []string{asvBibleId, webBibleId},
	}
	queryParameters := bibleSectionParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// content-type=text&include-chapter-numbers=true&include-notes=true&include-titles=true&include-verse-numbers=true&include-verse-spans=true&parallels=685d1470fe4d5c3b-01%2C32664dc3288a28df-03
}

func ExampleBiblePassageParams_ProduceQueryParameters() {
	asvBibleId := "685d1470fe4d5c3b-01"
	webBibleId := "32664dc3288a28df-03"
	biblePassageParams := BiblePassageParams{
		ContentType:           "text",
		IncludeNotes:          true,
		IncludeTitles:         true,
		IncludeChapterNumbers: true,
		IncludeVerseNumbers:   true,
		IncludeVerseSpans:     true,
		Parallels:             []string{asvBibleId, webBibleId},
		UseOrgId:              true,
	}
	queryParameters := biblePassageParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// content-type=text&include-chapter-numbers=true&include-notes=true&include-titles=true&include-verse-numbers=true&include-verse-spans=true&parallels=685d1470fe4d5c3b-01%2C32664dc3288a28df-03&use-org-id=true
}

func ExampleBibleVerseParams_ProduceQueryParameters() {
	asvBibleId := "685d1470fe4d5c3b-01"
	webBibleId := "32664dc3288a28df-03"
	bibleVerseParams := BibleVerseParams{
		ContentType:           "text",
		IncludeNotes:          true,
		IncludeTitles:         true,
		IncludeChapterNumbers: true,
		IncludeVerseNumbers:   true,
		IncludeVerseSpans:     true,
		Parallels:             []string{asvBibleId, webBibleId},
		UseOrgId:              true,
	}
	queryParameters := bibleVerseParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// content-type=text&include-chapter-numbers=true&include-notes=true&include-titles=true&include-verse-numbers=true&include-verse-spans=true&parallels=685d1470fe4d5c3b-01%2C32664dc3288a28df-03&use-org-id=true
}

func ExampleBibleSearchParams_ProduceQueryParameters() {
	bibleSearchParams := BibleSearchParams{
		Query:     "Love",
		Limit:     3,           // Default value is 10
		Offset:    5,           // Pagination
		Sort:      "canonical", // Supported values are relevance (default), canonical and reverse-canonical
		Range:     "MAT.1-REV.22",
		Fuzziness: "0", // Default value is AUTO. This accounts for difference in spelling
	}
	queryParameters := bibleSearchParams.ProduceQueryParameters().Encode()
	fmt.Println(queryParameters)
	// Output:
	// fuzziness=0&limit=3&offset=5&query=Love&range=MAT.1-REV.22&sort=canonical
}
