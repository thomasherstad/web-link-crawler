package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Println("Problem parsing html: %w", err)
		return nil, err
	}

	//Because the OG one is actually ""
	doc = doc.FirstChild

	// Just to make it not complain when I compile
	if rawBaseURL == "" {
		return nil, nil
	}
	getLink(doc)

	//Just to make it not complain when I compile
	return []string{}, nil
}

// In-work
func getLink(node *html.Node) {
	if node.Data == "a" {
		return
	}
}

// Next step: use the getLink function to get just the link from an a-tag. Then make another crawl function that traverses the tree and that can be recursively called
