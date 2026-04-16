package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	pageData := make([]PageData, 0, len(pages))
	for _, k := range keys {
		pageData = append(pageData, pages[k])
	}

	data, err := json.MarshalIndent(pageData, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding json: %v", err)
	}

	err = os.WriteFile(filename, data, 0o644)
	if err != nil {
		return fmt.Errorf("couldn't not write file: %v", err)
	}

	return nil
}
