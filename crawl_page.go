package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.ch <- struct{}{}
	defer func() {
		<-cfg.ch
		cfg.wg.Done()
	}()

	//NOTE: Avsluta om vi nått gränsen
	if cfg.size() >= cfg.maxPages {
		return
	}

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing currentUrl, %v\n", err)
		return
	}
	if cfg.baseURL.Hostname() != currentUrl.Hostname() {
		fmt.Println("error: not an internal url")
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing %s - %v", rawCurrentURL, err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html for %s - %v\n", rawCurrentURL, err)
		return
	}

	links, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Errorf("error getting links from %s - %v\n", rawCurrentURL, err)
	}
	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

}

// vim: ts=4
