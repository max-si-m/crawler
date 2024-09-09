package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
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

	config := config{
		baseURL: baseURL,
		pages: make(map[string]int),
		concurrencyControl: make(chan struct{}, 2),
		mu: &sync.Mutex{},
	}
	fmt.Printf("starting crawl of: %s\n", baseURL)
	config.crawlPage(baseURL.String())

	fmt.Println("pages hash: ", config.pages)
}
