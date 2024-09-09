package main

import (
	"net/url"
	"sync"
)

type config struct {
	maxPages		   int
	pages              map[string]int
	concurrencyControl chan struct{}
	baseURL            *url.URL
	mu                 *sync.Mutex
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	isFirst = false

	cfg.mu.Lock()
	if _, ok := cfg.pages[normalizedURL]; !ok {
		isFirst = true
	}
	cfg.pages[normalizedURL]++
	cfg.mu.Unlock()

	return isFirst
}

func (cfg *config) reachLimit() bool {
	cfg.mu.Lock()
	reached := len(cfg.pages) >= cfg.maxPages
	cfg.mu.Unlock()

	return reached
}

func NewConfig(baseURL *url.URL, maxPages, 	maxConcurrency int) *config {
	return &config{
		baseURL: baseURL,
		maxPages: maxPages,
		pages: make(map[string]int),
		concurrencyControl: make(chan struct{}, maxConcurrency),
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}
}
