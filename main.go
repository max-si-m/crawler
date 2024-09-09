package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
)

type KeyValue struct {
	Key   string
	Value int
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("# usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	rawBaseURL := args[0]
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("invalid params passed ", err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid maxConcurrency passed ", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("invalid maxPages passed ", err)
		os.Exit(1)
	}

	config := NewConfig(baseURL, maxPages, maxConcurrency)
	fmt.Printf("starting crawl of: %s\n", baseURL)

	config.wg.Add(1)
	go config.crawlPage(rawBaseURL)
	config.wg.Wait()

	printReport(config.pages, config.baseURL.String())
}


func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Println("REPORT for", baseURL)
	fmt.Println("=============================")

	var sortedPages []KeyValue

	for key, value := range pages {
		sortedPages = append(sortedPages, KeyValue{Key: key, Value: value})
	}

	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Value > sortedPages[j].Value
	})

	for _, kv := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", kv.Value, kv.Key)
	}
}
