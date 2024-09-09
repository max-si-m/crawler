package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[0]
	baseURL, err := url.Parse(rawBaseURL)
	if len(args) > 1 {
		fmt.Println("invalid base url", err)
		os.Exit(1)
	}

	config := NewConfig(baseURL, 6)
	fmt.Printf("starting crawl of: %s\n", baseURL)

	config.wg.Add(1)
	go config.crawlPage(rawBaseURL)
	config.wg.Wait()

	for normalizedURL, count := range config.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
