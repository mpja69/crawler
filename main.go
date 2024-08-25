package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("too few arguments provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Max nbr of concurrent routines wrong: %v", err)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Max nbr of pages wrong: %v", err)
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, int(maxConcurrency), int(maxPages))
	if err != nil {
		fmt.Println("error: %v", err)
	}
	fmt.Printf("starting crawl %s, with max %d routines, and max %d pages\n", rawBaseURL, maxConcurrency, maxPages)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}

// vim: set ts=4
