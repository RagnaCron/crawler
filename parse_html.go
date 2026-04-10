package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	reader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	return doc.Find("h1, h2").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}

	p := doc.Find("main").First().Find("p").Text()
	if p != "" {
		return p
	}

	return doc.Find("p").First().Text()
}
