package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) map[string]int {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		newError := fmt.Errorf("error parsing rawBaseURL: %s, error: %w", rawBaseURL, err)
		fmt.Println(newError)
		return pages
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		newErr := fmt.Errorf("error parsing rawCurrentURL: %s, error: %w", rawCurrentURL, err)
		fmt.Println(newErr)
		return pages
	}
	if baseURL.Host != currentURL.Host {
		return pages
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		newErr := fmt.Errorf("error normalizing url: %s, error: %w", rawCurrentURL, err)
		fmt.Println(newErr)
		return pages
	}

	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL]++
		return pages
	}
	pages[normalizedURL] = 1

	fmt.Printf("\nGetting html from %s\n", normalizedURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		newErr := fmt.Errorf("error getting the html: %w", err)
		fmt.Println(newErr)
		return pages
	}
	fmt.Printf("Got html from %s. Extracting links.\n", normalizedURL)

	newLinks, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		newErr := fmt.Errorf("error getting URLs from html: %w", err)
		fmt.Println(newErr)
		return pages
	}
	for _, link := range newLinks {
		pages = crawlPage(rawBaseURL, link, pages)
	}

	return pages
}
