package gorae

import (
	"fmt"
	"log"
	"strings"

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

func QueryWordJSON(word string, minifyJSON ...bool) (string, error) {
	definition, err := QueryWord(word)
	if err != nil {
		return "", err
	}

	minify := false
	if len(minifyJSON) > 0 && minifyJSON[0] {
		minify = true
	}

	definitionJSON, err := convertWordDefinitionToJSON(definition, minify)
	if err != nil {
		return "", err
	}
	return string(definitionJSON), nil
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
		var types, definition, examples string

		entrySelection.Each(func(i int, chunk *goquery.Selection) {
			// entry number
			if chunk.Find("span.n_acep").Text() != "" {
				n, err := sanitizeStrNum(chunk.Find("span.n_acep").Text())
				if err != nil {
					log.Fatal(err)
				}

				num = n
			}

			// types of the entry
			chunk.Find("abbr.d, abbr.g").Each(func(i int, title *goquery.Selection) {
				typeOfEntry, _ := title.Attr("title")
				types += typeOfEntry + ", "
			})
			types = strings.Trim(types, " ,") // remove last ' ' and ','

			// examples of entry
			chunk.Find("span.h").Each(func(i int, exampleSelection *goquery.Selection) {
				examples += exampleSelection.Text() + " "
			})
			examples = strings.Trim(examples, " ") // Remove last ' '

			// remove number, types and examples from DOM so we can get a clean definition
			chunk.Find("span.n_acep").Remove()
			chunk.Find("abbr.d, abbr.g").Remove()
			chunk.Find("span.h").Remove()
			definition = strings.Trim(chunk.Text(), " ")
		})

		entry := &WordDefinitionEntry{
			Num:        num,
			Types:      types,
			Definition: definition,
			Examples:   examples,
		}

		entries = append(entries, entry)
	})

	return entries, nil
}
