package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	fmt.Printf("starting crawl of: %s\n", baseURL)
	fmt.Println(getHTML(baseURL))
}

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("can not open %s, %v", rawURL, err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("can make GET request to %s, %v", rawURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode > 400 {
		return "", fmt.Errorf("invalid response status: %v", res.StatusCode)
	}

	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("invalid content type: %v", res.Header.Get("Content-Type"))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("can not read request body: %v", err)
	}

	return string(body), nil
}
