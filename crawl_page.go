package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	parsedURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	if currentURL.Hostname() != parsedURL.Hostname() {
		return
	}

	nCurURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	_, ok := pages[nCurURL]
	if ok {
		pages[nCurURL]++
		return
	}
	pages[nCurURL] = 1

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("%s\n")
		return
	}

	urls, err := getURLsFromHTML(html, parsedURL)

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
