package main

import (
	"net/url"
)
func normalizeURL(fullUrl string) (string, error) {
	parsedUrl, err := url.Parse(fullUrl)
	if err != nil {
		return "", err
	}

	return parsedUrl.Host + parsedUrl.Path, nil
}
