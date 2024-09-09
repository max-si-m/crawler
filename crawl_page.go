package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !sameDomain(rawBaseURL, rawCurrentURL) {
		return
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing current URL:", err)
		return
	}

	if !currentURL.IsAbs() {
		currentURL = baseURL.ResolveReference(currentURL)
	}

	key, err := normalizeURL(currentURL.String())
	if err != nil {
		fmt.Println("Error normalizing URL:", err)
		return
	}

	pages[key]++

	if pages[key] > 1 {
		fmt.Printf("Already crawled: %s\n", key)
		return
	}

	currentHTML, err := getHTML(currentURL.String())
	if err != nil {
		return
	}

	urls, err := getURLsFromHTML(currentHTML, rawCurrentURL)
	if err != nil {
		return
	}

	for _, rawNextURL := range urls {
		crawlPage(rawBaseURL, rawNextURL, pages)
	}
}

func sameDomain(baseURL, rawCurrentURL string) bool {
	baseParsed, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	currentParsed, err := url.Parse(rawCurrentURL)
	if err != nil {
		return false
	}

	return baseParsed.Host == currentParsed.Host
}
