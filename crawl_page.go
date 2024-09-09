package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	if !sameDomain(cfg.baseURL, rawCurrentURL) {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing current URL:", err)
		return
	}

	if !currentURL.IsAbs() {
		currentURL = cfg.baseURL.ResolveReference(currentURL)
	}

	normalizedURL, err := normalizeURL(currentURL.String())
	if err != nil {
		fmt.Println("Error normalizing current URL:", err)
		return
	}
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
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
		cfg.crawlPage(rawNextURL)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	isFirst = false

	if _, ok := cfg.pages[normalizedURL]; !ok {
		isFirst = true
	}
	cfg.mu.Lock()
	cfg.pages[normalizedURL]++
	cfg.mu.Unlock()

	return
}

func sameDomain(baseURL *url.URL, rawCurrentURL string) bool {
	currentParsed, err := url.Parse(rawCurrentURL)
	if err != nil {
		return false
	}

	return baseURL.Host == currentParsed.Host
}
