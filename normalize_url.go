package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {
	sURL, err := url.Parse(u)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	return fmt.Sprintf("%s%s", sURL.Host, strings.TrimSuffix(sURL.Path, "/")), nil
}
