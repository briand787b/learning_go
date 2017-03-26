package main

import (
	"fmt"
	iterative "github.com/briand787b/udemy/pracpac2"
)

func main() {
	xz := iterative.Filter([]int{1,2,3,4}, greaterThanOne)
	fmt.Println(xz)
}

func greaterThanOne(comparison int) bool {
	return comparison > 1
}
