package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "normalized input",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme https",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme http",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing /",
			inputURL: "blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme https and trailing /",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheeme http and trailing /",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "longer path",
			inputURL: "https://blog.boot.dev/path/is/longer/",
			expected: "blog.boot.dev/path/is/longer",
		},
		{
			name:     "Capital letters",
			inputURL: "https://blog.BOOT.Dev/PATH",
			expected: "blog.boot.dev/path",
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := normalizeURL(testCase.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, testCase.name, err)
				return
			}
			if actual != testCase.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, testCase.name, testCase.expected, actual)
			}
		})
	}
}
