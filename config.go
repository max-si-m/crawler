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
