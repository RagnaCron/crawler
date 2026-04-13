package main

import (
	"fmt"
	"os"
)

func main() {
	argLen := len(os.Args[1:])
	if argLen < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if argLen > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseURL)
}
