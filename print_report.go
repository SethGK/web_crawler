package main

import (
	"fmt"
	"sort"
)

type pageInfo struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	var pageList []pageInfo
	for url, count := range pages {
		pageList = append(pageList, pageInfo{url: url, count: count})
	}

	sort.Slice(pageList, func(i, j int) bool {
		if pageList[i].count == pageList[j].count {
			return pageList[i].url < pageList[j].url
		}
		return pageList[i].count > pageList[j].count
	})

	for _, p := range pageList {
		fmt.Printf("Found %d internal links to %s\n", p.count, p.url)
	}

}
