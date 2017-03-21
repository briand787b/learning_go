package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Println("crawl error:", err)
	}
	return list
}

func main() {
	// resp, err := http.Get("http://google.com")

	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		fmt.Println(n)
		list := <-worklist	// chan []string
		for _, link := range list {	// ["www.google.com", "www.facebook.com"]
			if !seen[link] {	// seen["www.google.com"] => false
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)	// worklist <- ["pornhub.com", "bkdoorsluts.com"]
					fmt.Println("finishing main go routine")
				}(link)
			}
		}
	}
}
