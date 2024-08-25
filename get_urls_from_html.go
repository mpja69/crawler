package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error couldn't parse base url, %w", err)
	}

	body := strings.NewReader(htmlBody)
	root, err := html.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse html, %w", err)
	}

	var urls []string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("error parsing href, %v", err)
						continue
					}
					resolvedURL := baseURL.ResolveReference(href)
					if resolvedURL.Host != baseURL.Host {
						continue
					}
					urls = append(urls, resolvedURL.String())
					break
				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(root)
	return urls, nil
}

// vim: ts=4
