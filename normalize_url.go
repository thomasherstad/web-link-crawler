package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Could not parse URL")
		return "", err
	}
	normalizedUrl := parsedUrl.Host + parsedUrl.Path
	normalizedUrl = strings.TrimRight(normalizedUrl, "/")
	normalizedUrl = strings.ToLower(normalizedUrl)
	return normalizedUrl, nil
}
