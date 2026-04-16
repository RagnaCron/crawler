package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing url: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n", baseURL)

	cfg := config{
		pages:              make(map[string]PageData),
		baseURL:            parsedURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}
}
