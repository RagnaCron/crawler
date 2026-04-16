package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.getPageLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	nCurURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if visited := cfg.addPageVisit(nCurURL); !visited {
		return
	}

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	page := extractPageData(html, rawCurrentURL)
	cfg.setPage(nCurURL, page)

	for _, url := range page.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}

	return true
}

func (cfg *config) setPage(nURL string, page PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[nURL] = page
}

func (cfg *config) getPageLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}
