package main

import "fmt"

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {
	var iBuffer [150]int
	slice := iBuffer[0:0]
	fmt.Println(len(slice))
	for i := 0; i < 20; i++ {
		slice = Extend(slice, i)
		fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	}

	asianLetters := "日本語"
	for i := 0; i < len(asianLetters); i++ {
		fmt.Printf("%q", asianLetters[i])
	}
	fmt.Println()
	for _, v := range asianLetters {
		fmt.Printf("%q", v)
	}
	fmt.Println()
}
