package main

import "fmt"

func main() {
	var slice = []int{}
	for i, _ := range slice {
		slice[i] = i
	}

	for i, _ := range slice {
		fmt.Println(slice[i])
	}

	fmt.Printf("Slice has type: %T", slice)
}
