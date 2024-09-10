package main

import (
	"fmt"
	"os"
)

func main() {
	//if number of cli arguments <  1, print "no website provided"
	//if number of cli arguments > 1, print "too  many arguments provided"
	//if number of cli arguments == 1, print "starting crawl of: Base url"
	if len(os.Args[1:]) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args[1:]) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} 

	url := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", url)
	
	pages := make(map[string]int)
	pages = crawlPage(url, url, pages)

	fmt.Println(pages)
}
