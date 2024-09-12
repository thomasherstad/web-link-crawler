package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {

	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	startURL := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", startURL)

	var c config
	c.pages = make(map[string]int)

	var err error
	c.baseURL, err = url.Parse(startURL)
	if err != nil {
		fmt.Printf("error parsing url: %s, error: %v", startURL, err)
		return
	}

	c.mu = &sync.Mutex{}
	maxConcurrency := 1
	c.concurrencyControl = make(chan struct{}, maxConcurrency)
	c.wg = &sync.WaitGroup{}

	c.wg.Add(1)
	c.concurrencyControl <- struct{}{}
	go c.crawlPage(startURL)

	c.wg.Wait()
	fmt.Println(c.pages)
}
