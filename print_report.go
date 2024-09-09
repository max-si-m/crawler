package main

import (
	"fmt"
	"sort"
)

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
