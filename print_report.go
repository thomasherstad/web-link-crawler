package main

import (
	"fmt"
	"sort"
)

type item struct {
	webpage string
	number  int
}

func printReport(pages map[string]int, baseURL string) {

	var pagesSlice []item
	for page, n := range pages {
		var itm item
		itm.webpage = page
		itm.number = n
		pagesSlice = append(pagesSlice, itm)
	}

	sortedPages := sortPages(pagesSlice)

	printHeader(baseURL)

	for _, itm := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", itm.number, itm.webpage)
	}

}

func sortPages(pages []item) []item {
	sort.Slice(pages, func(i, j int) bool { return pages[i].number > pages[j].number })
	return pages
}

func printHeader(baseURL string) {
	separator := "=================================================="
	fmt.Printf("\n\n\n")
	fmt.Println(separator)
	fmt.Println("INTERNAL LINKS REPORT FOR " + baseURL)
	fmt.Println(separator)
}
