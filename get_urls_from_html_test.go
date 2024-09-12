package main

import (
	"reflect"
	"testing"
)

// Need to add more tests
func TestGetUrlsFromHtml(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		inputURL  string
		expected  []string
	}{
		{
			name: "simple",
			inputHTML: `
				<html>
					<body>
						<a href="https://blog.boot.dev"><span>Go to Boot.dev, you React Andy</span></a>
					</body>
				</html>
				`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name: "absolute and relative URLs",
			inputHTML: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name: "html without a-tags",
			inputHTML: `
				<html>
					<body>
						<span>Boot.dev</span>
						<span>Boot.dev</span>
					</body>
				</html>
				`,
			inputURL: "",
			expected: nil,
		},
		{
			name: "relative url without '/'",
			inputHTML: `
				<html>
					<body>
						<div>
							<a href="user?id=ocean" class="hnuser">ocean</a>
						</div>
					</body>
				</html>
			`,
			inputURL: "https://news.ycombinator.com",
			expected: []string{"https://news.ycombinator.com/user?id=ocean"},
		},
	}

	for i, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(testCase.inputHTML, testCase.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: Got error: %v", i, testCase.name, err)
				return
			}

			result := reflect.DeepEqual(actual, testCase.expected)
			if result != true {
				t.Errorf("Test %v - %s FAIL: expected links: %v, actual links: %v", i, testCase.name, testCase.expected, actual)
			}
		})
	}

}
