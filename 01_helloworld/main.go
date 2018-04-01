package main

import "fmt"


const q = "42"

func main() {
	fmt.Printf("%x - %b - %#x", 42, 42, 42)
	fmt.Println()
	fmt.Println(q)
	fmt.Printf("fuck ni**ers " + q + "\n")
	fmt.Println("q - fuck ni**ers", q)
	var a, b, c string
	n, err := fmt.Scanln(&a, &b, &c)
	fmt.Println(a)
	fmt.Println("number of items scanned - ", n)
	fmt.Println("error that was encountered - ", err)
}
