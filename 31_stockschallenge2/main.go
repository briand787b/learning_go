package main

import (
	"fmt"
	"math/rand"
	"math"
)

type stockSale struct {
	price int
	index int
}

func main() {
	stockPrices := [1000]int{}
	for i := 0; i < len(stockPrices); i++ {
		stockPrices[i] = rand.Int()>>55
	}

	fmt.Println(stockPrices)

	fmt.Println("First value of stock prices: ", stockPrices[0])
	// bitwise shift value
	stockPrices[0] = stockPrices[0]>>49
	fmt.Println("First value of stock prices: ", stockPrices[0])

	fmt.Println("Max value of int64: ", math.MaxInt64)
	fmt.Println("Max value of 64 shifted int61: ", math.MaxInt64>>61)
}

func findHighestFromIndex(index int, slice []int) <-chan stockSale{
	out := make(chan stockSale)
	go func() {
		// var min int
		for i := index; i < len(slice)-index; i++ {

		}
	}()
	return out
}
