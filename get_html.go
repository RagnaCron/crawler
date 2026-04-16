package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("couldn't create request: %s", err)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")
	req.Header.Set("Accept", "text/html")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("couldn't make request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("wrong content type, got: %v, expected: text/html", contentType)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %s", err)
	}

	return string(data), nil
}
