package main

import (
	"log"
	"gopl.io/ch5/links"
	"fmt"
	"flag"
)

// counting semaphore
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // return token
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	depth := flag.Int("depth", 1000000, "the maximum number of links to be fetched")
	flag.Parse()

	linksCh := make(chan []string)
	seen := make(map[string]bool)
	var n, c int

	n++
	go func(){ linksCh <- flag.Args() }()

	for ; n > 0; n-- {
		page := <- linksCh
		for _, link := range page {
			if !seen[link] && c < *depth{
				c++
				n++
				seen[link] = true
				go func(ln string) {
					linksCh <- crawl(ln)
				}(link)
			}
		}
	}
}