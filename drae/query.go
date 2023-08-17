package drae

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/rulu158/gorae/scraper"
)

func QueryWord(word string) (*WordDefinition, error) {
	doc, err := scraper.Scrape(word)
	if err != nil {
		return nil, err
	}

	definition, err := getWordDefinition(doc)
	if err != nil {
		return nil, err
	}

	return definition, nil
}

func getWordDefinition(doc *goquery.Document) (*WordDefinition, error) {
	wordDefinition := &WordDefinition{}

	doc.Find("#resultados > article").Each(func(i int, articleSelection *goquery.Selection) {
		word := articleSelection.Find("header").Text()

		entries, err := getWordDefinitionEntries(articleSelection)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		wordDefinition.Word = word
		wordDefinition.Entries = entries
	})

	return wordDefinition, nil
}

func getWordDefinitionEntries(articleSelection *goquery.Selection) ([]*WordDefinitionEntry, error) {
	entries := []*WordDefinitionEntry{}

	articleSelection.Find("p.j, p.j2").Each(func(i int, entrySelection *goquery.Selection) {
		var num int
		var types, definition string

		entrySelection.Each(func(i int, chunk *goquery.Selection) {
			if chunk.Find("span.n_acep").Text() != "" {
				n, err := SanitizeStrNum(chunk.Find("span.n_acep").Text())
				if err != nil {
					log.Fatal(err)
				}

				num = n
			}

			chunk.Find("abbr.d, abbr.g").Each(func(i int, title *goquery.Selection) {
				typeOfEntry, _ := title.Attr("title")
				types += typeOfEntry + ", "
			})
			types = types[0 : len(types)-2]

			chunk.Find("span.n_acep").Remove()
			chunk.Find("abbr.d, abbr.g").Remove()
			chunk.Find("span.h").Remove()
			definition = chunk.Text()
		})

		entry := &WordDefinitionEntry{
			Num:        num,
			Types:      types,
			Definition: definition,
		}

		entries = append(entries, entry)
	})

	return entries, nil
}
