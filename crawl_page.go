package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (c *config) crawlPage(rawCurrentURL string) {
	defer func() { <-c.concurrencyControl }()
	defer c.wg.Done()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		newErr := fmt.Errorf("error parsing rawCurrentURL: %s, error: %w", rawCurrentURL, err)
		fmt.Println(newErr)
		return
	}

	if c.baseURL.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		newErr := fmt.Errorf("error normalizing url: %s, error: %w", rawCurrentURL, err)
		fmt.Println(newErr)
		return
	}

	//Lock the pages map
	c.mu.Lock()
	_, ok := c.pages[normalizedURL]
	if ok {
		c.pages[normalizedURL]++
		c.mu.Unlock() //Unlock if url is in the map
		return
	}
	c.pages[normalizedURL] = 1
	c.mu.Unlock() //Unlock if url is not already in the map

	fmt.Printf("\nGetting html from %s\n", normalizedURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		newErr := fmt.Errorf("error getting the html: %w", err)
		fmt.Println(newErr)
		return
	}
	fmt.Printf("Got html from %s. Extracting links.\n", normalizedURL)

	newLinks, err := getURLsFromHTML(html, c.baseURL.String())
	if err != nil {
		newErr := fmt.Errorf("error getting URLs from html: %w", err)
		fmt.Println(newErr)
		return
	}
	for _, link := range newLinks {
		c.wg.Add(1)
		go c.crawlPage(link)
	}

}
