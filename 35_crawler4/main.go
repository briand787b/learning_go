package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"flag"
	"math"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	depthPtr := flag.Int("depth", math.MaxInt8, "the maximum depth of links to which the program will search")
	flag.Parse()
	if *depthPtr < 2 {
		log.Fatal("depth seach must be greater than 1")
	}

	linkSetsCh := make(chan []string)

	go func() { linkSetsCh <- flag.Args() }()

	for list := range linkSetsCh {
		for _, link := range list {
			go func(ln string) {
				linkSetsCh <- crawl(ln)
			}(link)
		}
	}
}
