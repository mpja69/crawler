package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages    map[string]int
	baseURL  *url.URL
	mux      *sync.Mutex
	ch       chan struct{}
	wg       *sync.WaitGroup
	maxPages int
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse the website url")
	}
	return &config{
		pages:    make(map[string]int),
		baseURL:  baseURL,
		mux:      &sync.Mutex{},
		ch:       make(chan struct{}, maxConcurrency), // NOTE: Skapa x antal buffert, för att begränsa antal go-routines
		wg:       &sync.WaitGroup{},
		maxPages: int(maxPages),
	}, nil
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mux.Lock()
	defer cfg.mux.Unlock()
	_, ok := cfg.pages[normalizedURL]
	cfg.pages[normalizedURL]++
	return !ok
}
func (cfg *config) size() int {
	cfg.mux.Lock()
	defer cfg.mux.Unlock()
	return len(cfg.pages)
}
