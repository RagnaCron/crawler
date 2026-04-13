package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)

	doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("src")
		if !exists {
			return
		}

		u, ok := normalize(baseURL, href)
		if !ok {
			return
		}

		urls = append(urls, u.String())
	})

	return urls, nil
}
