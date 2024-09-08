package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error){
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)
	var trav func(n *html.Node)

	trav = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					url := a.Val
					if !isAbsoluteUrl(url, rawBaseURL) {
						url = rawBaseURL + url
					}

					urls = append(urls, url)
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			trav(c)
		}
	}

	trav(doc)

	return urls, nil
}

func isAbsoluteUrl(url, baseUrl string) bool {
	return (strings.Contains(url, "http://") || strings.Contains(url, "https://")) || strings.Contains(url, baseUrl)
}
