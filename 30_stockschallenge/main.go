package main

import (
	"fmt"
)

func main() {
	prices := []int{38, 24, 13, 59, 12, 83, 6, 69, 100, 84, 80, 98, 81, 35, 51, 44, 54, 56, 77, 87, 78, 73, 55, 9, 3, 40, 95, 46, 36, 39}
	minPrice := prices[0]
	maxPrice := prices[1]
	maxDiff := 0
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - minPrice; diff > maxDiff {
			maxPrice = prices[i]
			maxDiff = diff
		}

		if prices[i] < minPrice {
			fmt.Println("price is lower than buying price ",
				prices[i])
			minPrice = prices[i]
		}
	}

	fmt.Printf("Buy Price: %v, Sell Price: %v, Max Diff: %v\n", minPrice,
		maxPrice, maxDiff)
}
