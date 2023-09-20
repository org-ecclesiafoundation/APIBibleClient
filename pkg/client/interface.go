//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"ecclesiafoundation.org/apibibleclient/pkg/client/internal"
	"ecclesiafoundation.org/apibibleclient/pkg/client/params"
)

// The GetBibles function calls the
// `/bibles` end point of the API
func GetBibles(apiKey string, params *params.BiblesParams) (string, error) {
	return internal.GetBibles(apiKey, params)
}

// The GetBibleById function calls the
// `/bibles/{bibleId}` end point of the API
func GetBibleById(apiKey string, bibleId string) (string, error) {
	return internal.GetBibleById(apiKey, bibleId)
}

// The GetAudioBibles function calls the
// `/audio-bibles` end point of the API
func GetAudioBibles(apiKey string, params *params.AudioBiblesParams) (string, error) {
	return internal.GetAudioBibles(apiKey, params)
}

// The GetAudioBibleById function calls the
// `/audio-bibles/{audioBibleId}` end point of the API
func GetAudioBibleById(apiKey string, bibleId string) (string, error) {
	return internal.GetAudioBibleById(apiKey, bibleId)
}

// The GetBibleBooks function calls the
// `/bibles/{bibleId}/books` end point of the API
func GetBibleBooks(apiKey string, bibleId string, params *params.BibleBooksParams) (string, error) {
	return internal.GetBibleBooks(apiKey, bibleId, params)
}

// The GetBibleBookById function calls the
// `/bibles/{bibleId}/books/{bookId}` end point of the API
func GetBibleBookById(apiKey string, bibleId string, bookId string, params *params.BibleBookParams) (string, error) {
	return internal.GetBibleBookById(apiKey, bibleId, bookId, params)
}

// The GetAudioBibleBooks function calls the
// `/audio-bibles/{audioBibleId}/books` end point of the API
func GetAudioBibleBooks(apiKey string, audioBibleId string, params *params.AudioBibleBooksParams) (string, error) {
	return internal.GetAudioBibleBooks(apiKey, audioBibleId, params)
}

// The GetAudioBibleBookById function calls the
// `/audio-bibles/{audioBibleId}/books/{bookId}` end point of the API
func GetAudioBibleBookById(apiKey string, audioBibleId string, bookId string, params *params.AudioBibleBookParams) (string, error) {
	return internal.GetAudioBibleBookById(apiKey, audioBibleId, bookId, params)
}

// The GetBibleChapters function calls the
// `/bibles/{bibleId}/books/{bookId}/chapters` end point of the API
func GetBibleChapters(apiKey string, bibleId string, bookId string) (string, error) {
	return internal.GetBibleChapters(apiKey, bibleId, bookId)
}

// The GetBibleChapterById function calls the
// `/bibles/{bibleId}/chapters/{chapterId}` end point of the API
func GetBibleChapterById(apiKey string, bibleId string, chapterId string, params *params.BibleChapterParams) (string, error) {
	return internal.GetBibleChapterById(apiKey, bibleId, chapterId, params)
}

// The GetAudioBibleChapters function calls the
// `/audio-bibles/{audioBibleId}/books/{bookId}/chapters` end point of the API
func GetAudioBibleChapters(apiKey string, audioBibleId string, bookId string) (string, error) {
	return internal.GetAudioBibleChapters(apiKey, audioBibleId, bookId)
}

// The GetAudioBibleChapterById function calls the
// `/audio-bibles/{audioBibleId}/chapters/{chapterId}` end point of the API
func GetAudioBibleChapterById(apiKey string, audioBibleId string, chapterId string) (string, error) {
	return internal.GetAudioBibleChapterById(apiKey, audioBibleId, chapterId)
}

// The GetBibleBookSections function calls the
// `/bibles/{bibleId}/books/{bookId}/sections` end point of the API
func GetBibleBookSections(apiKey string, bibleId string, bookId string) (string, error) {
	return internal.GetBibleBookSections(apiKey, bibleId, bookId)
}

// The GetBibleChapterSections function calls the
// `/bibles/{bibleId}/chapters/{chapterId}/sections` end point of the API
func GetBibleChapterSections(apiKey string, bibleId string, chapterId string) (string, error) {
	return internal.GetBibleChapterSections(apiKey, bibleId, chapterId)
}

// The GetBibleSectionById function calls the
// `/bibles/{bibleId}/sections/{sectionId}` end point of the API
func GetBibleSectionById(apiKey string, bibleId string, sectionId string, params *params.BibleSectionParams) (string, error) {
	return internal.GetBibleSectionById(apiKey, bibleId, sectionId, params)
}

// The GetBiblePassage function calls the
// `/bibles/{bibleId}/passages/{passageId}` end point of the API
func GetBiblePassage(apiKey string, bibleId string, passageId string, params *params.BiblePassageParams) (string, error) {
	return internal.GetBiblePassage(apiKey, bibleId, passageId, params)
}

// The GetBibleChapterVerses function calls the
// `/bibles/{bibleId}/chapters/{chapterId}/verses` end point of the API
func GetBibleChapterVerses(apiKey string, bibleId string, chapterId string) (string, error) {
	return internal.GetBibleChapterVerses(apiKey, bibleId, chapterId)
}

// The GetBibleVerseById function calls the
// `/bibles/{bibleId}/verses/{verseId}` end point of the API
func GetBibleVerseById(apiKey string, bibleId string, verseId string, params *params.BibleVerseParams) (string, error) {
	return internal.GetBibleVerseById(apiKey, bibleId, verseId, params)
}

// The GetBibleSearchResults function calls the
// `/bibles/{bibleId}/search` end point of the API
func GetBibleSearchResults(apiKey string, bibleId string, params *params.BibleSearchParams) (string, error) {
	return internal.GetBibleSearchResults(apiKey, bibleId, params)
}
