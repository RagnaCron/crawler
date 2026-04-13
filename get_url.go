package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var urls []string

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
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

func normalize(base *url.URL, raw string) (*url.URL, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, false
	}

	u, err := url.Parse(raw)
	if err != nil {
		return nil, false
	}

	// Skip unwanted schemes
	if u.Scheme == "mailto" || u.Scheme == "javascript" {
		return nil, false
	}

	// Case 1: absolute URL
	if u.IsAbs() {
		return u, true
	}

	return base.ResolveReference(u), true
}
