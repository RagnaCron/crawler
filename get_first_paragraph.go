package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
