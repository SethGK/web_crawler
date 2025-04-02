package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing base url: %s: %v\n", rawBaseURL, err)
		return
	}

	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing current URL %s: %v\n", rawCurrentURL, err)
		return
	}

	if base.Host != current.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	if count, ok := pages[normalizedURL]; ok {
		pages[normalizedURL] = count + 1
		return
	}

	pages[normalizedURL] = 1

	fmt.Printf("crawling: %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error fetching HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("Error extracting URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}

}
