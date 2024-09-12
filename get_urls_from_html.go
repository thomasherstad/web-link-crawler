package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("problem parsing html: %w", err)
	}

	var toVisit []*html.Node
	var links []string
	_, _, links = traverseHTML(doc, rawBaseURL, toVisit, links)

	return links, nil
}

func getLinkFromNode(node *html.Node, rawBaseURL string) (string, error) {
	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			newURL, err := url.Parse(attribute.Val)
			if err != nil {
				return "", fmt.Errorf("error parsing url: %s, error: %w", attribute.Val, err)
			}

			baseURL, err := url.Parse(rawBaseURL)
			if err != nil {
				return "", fmt.Errorf("error parsing baseURL: %s, error: %w", rawBaseURL, err)
			}

			resolvedURL := baseURL.ResolveReference(newURL)

			return resolvedURL.String(), nil
		}
	}
	return "", nil
}

func traverseHTML(node *html.Node, rawBaseURL string, toVisit []*html.Node, links []string) (*html.Node, []*html.Node, []string) {

	if node == nil {
		return nil, toVisit, links
	}

	//If a-tag, get link
	if node.Type == html.ElementNode && node.Data == "a" {
		link, err := getLinkFromNode(node, rawBaseURL)
		if err != nil {
			fmt.Println(err)
			return nil, toVisit, links
		}
		links = append(links, link)
	}

	// if next sibling: Add next sibling to the stack
	if node.NextSibling != nil {
		toVisit = append(toVisit, node.NextSibling)
	}
	// if first child: Add first child to the stack
	if node.FirstChild != nil {
		toVisit = append(toVisit, node.FirstChild)
	}

	// traverse the next node from the toVisit stack
	l := len(toVisit)
	if l < 1 {
		return nil, toVisit, links
	}
	newNode := toVisit[l-1]
	toVisit = toVisit[:l-1]

	return traverseHTML(newNode, rawBaseURL, toVisit, links)

}
