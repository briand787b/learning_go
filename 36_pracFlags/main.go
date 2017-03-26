package main

import (
	"flag"
	"math"
	"fmt"
)

func main() {
	flag.Int("depth", math.MaxInt8, "the maximum number of links to search through")
	flag.Parse()

	fmt.Println(flag.Args())
}