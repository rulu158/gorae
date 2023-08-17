package scraper

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const base_url = "https://dle.rae.es/"

func Scrape(word string) (*goquery.Document, error) {
	if word == "" || strings.Contains(word, " ") {
		return nil, errors.New("not a valid word")
	}
	req, err := http.NewRequest("GET", base_url+word, nil)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: remove headers
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/116.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS13,
		},
	}

	client := &http.Client{Transport: transport}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("status code is not 200")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.New("error getting document from body")
	}

	return doc, nil
}
