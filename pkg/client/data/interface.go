package data

import (
	"encoding/json"
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type BibleBook struct {
	Abbreviation string `json:"abbreviation"`
	BibleId      string `json:"bibleId"`
	Id           string `json:"id"`
	Name         string `json:"name"`
}

type BibleBooks struct {
	count          int
	Data           []BibleBook `json:"data"`
	lookupByAbbrev map[string]BibleBook
	lookupByName   map[string]BibleBook
	abbrevs        []string
	names          []string
	ids            []string
}

func (bibleBooks *BibleBooks) Count() int {
	return bibleBooks.count
}

// FindByAbbrev does an exact search of all the Bible Book abbreviations
// and returns a blank BibleBook{} struct if said book does not exist,
// and the BibleBook sought after if it does exist.
func (bibleBooks *BibleBooks) FindByAbbrev(abbrev string) (BibleBook, error) {
	bibleBook := bibleBooks.lookupByAbbrev[abbrev]
	if bibleBook == (BibleBook{}) {
		return bibleBook, fmt.Errorf("book abbreviation '%s' does not exist", abbrev)
	}
	return bibleBooks.lookupByAbbrev[abbrev], nil
}

// FindByName does an exact search of all the Bible Book names
// and returns a blank BibleBook{} struct if said book does not exist,
// and the BibleBook sought after if it does exist.
func (bibleBooks *BibleBooks) FindByName(name string) (BibleBook, error) {
	bibleBook := bibleBooks.lookupByName[name]
	if bibleBook == (BibleBook{}) {
		return bibleBook, fmt.Errorf("book name '%s' does not exist", name)
	}
	return bibleBooks.lookupByName[name], nil
}

// FindByOrder takes a number and seeks to find the book in the bible
// at that particular position
func (bibleBooks *BibleBooks) FindByOrder(order int) (BibleBook, error) {
	if order < 1 || bibleBooks.count < order {
		return BibleBook{}, fmt.Errorf("book order out of range: please use a number from 1 to %d", bibleBooks.count)
	} else {
		return bibleBooks.Data[order-1], nil
	}
}

// Abbrevs returns all bible book abbreviations for a particular bible
func (bibleBooks *BibleBooks) Abbrevs() []string {
	return bibleBooks.abbrevs
}

// Names returns all bible book names for a particular bible
func (bibleBooks *BibleBooks) Names() []string {
	return bibleBooks.names
}

// Ids returns all bible book names for a particular bible
func (bibleBooks *BibleBooks) Ids() []string {
	return bibleBooks.ids
}

// FuzzyFind searches through the abbreviations and names of books
// and returns the closest match according to Levenshtein distance
func (bibleBooks *BibleBooks) FuzzyFind(input string) (BibleBook, error) {
	rankByAbbrev := fuzzy.RankFindFold(input, bibleBooks.Abbrevs())
	lenAbbrevs := len(rankByAbbrev)
	rankByName := fuzzy.RankFindFold(input, bibleBooks.Names())
	lenNames := len(rankByName)
	if lenAbbrevs != 0 {
		closestAbbrev := rankByAbbrev[lenAbbrevs-1]
		return bibleBooks.lookupByAbbrev[closestAbbrev.Target], nil
	} else if lenNames != 0 {
		closestName := rankByName[lenNames-1]
		return bibleBooks.lookupByName[closestName.Target], nil
	} else {
		return BibleBook{}, fmt.Errorf("could not find book %s", input)
	}
}

// TODO: integrate this with the client
func ToBibleBooks(body string) (BibleBooks, error) {
	var bibleBooks BibleBooks
	jsonUnmarshalErr := json.Unmarshal([]byte(body), &bibleBooks)
	if jsonUnmarshalErr != nil {
		return BibleBooks{}, fmt.Errorf("failed to unmarshal json %s", body)
	} else {
		var bookAbbrevMap = make(map[string]BibleBook)
		var bookNameMap = make(map[string]BibleBook)
		var names []string
		var abbrevs []string
		var ids []string
		for _, book := range bibleBooks.Data {
			bookAbbrevMap[book.Abbreviation] = book
			bookNameMap[book.Name] = book
			abbrevs = append(abbrevs, book.Abbreviation)
			names = append(names, book.Name)
			ids = append(ids, book.Id)
		}
		bibleBooks.count = len(bibleBooks.Data)
		bibleBooks.lookupByAbbrev = bookAbbrevMap
		bibleBooks.lookupByName = bookNameMap
		bibleBooks.abbrevs = abbrevs
		bibleBooks.names = names
		bibleBooks.ids = ids
		return bibleBooks, nil
	}
}
