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

func NewConfig(baseURL *url.URL) *config {
	return &config{
		baseURL: baseURL,
		pages: make(map[string]int),
		concurrencyControl: make(chan struct{}, 2),
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}
}
