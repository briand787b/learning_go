package main

import "fmt"

func DoubleCap(slice []byte) []byte {
	newSlice := make([]byte, len(slice), 2*cap(slice))
	for i, _ := range slice {
		newSlice[i] = slice[i]
	}

	return newSlice
}

func Insert(slice []int, element int) []int {
	// Fill this out soon
	return []int{55, 55, 55, 55}
}

func main() {
	var iBuffer [15]byte
	slice := iBuffer[0:15]
	fmt.Printf("len: %d,  cap: %d\n", len(slice), cap(slice))
	slice = DoubleCap(slice)
	fmt.Printf("len: %d,  cap: %d\n", len(slice), cap(slice))
}
