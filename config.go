package main

import (
	"net/url"
	"sync"
)

type config struct {
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

func NewConfig(baseURL *url.URL, maxConcurrency int) *config {
	return &config{
		baseURL: baseURL,
		pages: make(map[string]int),
		concurrencyControl: make(chan struct{}, maxConcurrency),
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}
}
