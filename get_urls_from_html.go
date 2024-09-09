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

	var toVisit []*html.Node
	var links []string
	_, _, links = traverseHTML(doc, rawBaseURL, toVisit, links)

	fmt.Println(links)

	return links, nil
}

func getLinkFromNode(node *html.Node, rawBaseURL string) string {
	link := ""
	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			link = attribute.Val
			if link[:4] != "http" {
				link = rawBaseURL + link
			}
			return link
		}
	}
	return link
}

func traverseHTML(node *html.Node, rawBaseURL string, toVisit []*html.Node, links []string) (*html.Node, []*html.Node, []string) {

	if node == nil {
		return nil, toVisit, links
	}

	//If a-tag, get link
	if node.Type == html.ElementNode && node.Data == "a" {
		link := getLinkFromNode(node, rawBaseURL)
		if link != "" {
			links = append(links, link)
		}
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
