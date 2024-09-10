package main

import (
	"fmt"
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

	url := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", url)

	//TODO
	var c config
	c.pages = make(map[string]int)
	c.mu = &sync.Mutex{}
	maxConcurrency := 1
	c.concurrencyControl = make(chan struct{}, maxConcurrency)
	c.wg = &sync.WaitGroup{}
	c.crawlPage(url)

	fmt.Println(c.pages)
}
