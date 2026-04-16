package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	maxPages           int
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing url: %v\n", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error parsing maxPages: %v\n", err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error parsing maxConcurrency: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n", baseURL)

	cfg := config{
		maxPages:           maxPages,
		pages:              make(map[string]PageData),
		baseURL:            parsedURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}

	err = writeJSONReport(cfg.pages, "report.json")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
