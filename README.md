## WebCrawler
A lightweight, concurrent web crawler written in Go. It traverses internal links of a given base URL.

### Features

Concurrent crawling with configurable limits
Simple HTML link extraction
Generates a report of discovered internal pages and their link counts

### Usage
``` go run . <baseURL> <maxConcurrency> <maxPages> ```

### Attribution
This project was developed as part of a guided project from Boot.dev.