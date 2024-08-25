package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	var keys []string
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return pages[keys[i]] < pages[keys[j]]
	})
	for _, k := range keys {
		fmt.Printf("Found %d internal links to %s\n", pages[k], k)
	}
}

// vim: set ts=4
