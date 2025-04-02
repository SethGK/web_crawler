package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					parsedURL, err := url.Parse(attr.Val)
					if err != nil {
						absoluteURL := baseURL.ResolveReference(parsedURL).String()
						urls = append(urls, absoluteURL)
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = n.NextSibling {
			extractLinks(c)
		}
	}
	extractLinks(doc)
	return urls, nil
}
