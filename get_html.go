package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		return "", fmt.Errorf("error: status code: %v", resp.StatusCode)
	}
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("response body is not html. got: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	return string(body), nil
}
