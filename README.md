# crawler

A boot.dev project in Go

## Description

Run the web crawler by
`./crawl URL maxConcurrency maxPages`

Where `URL` is the web site you want to crawl, and `maxConcurrency`is the max number of goroutines you want to use, and `maxPages` is the max number pages you want to collect.

It counts the internal links on the web site.
