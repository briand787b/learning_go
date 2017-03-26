package main

import "fmt"

func main() {
	var a string
	fmt.Println("Guess my name")
	fmt.Scanln(&a)
	
	switch  a {
	case "Brian":
		fmt.Println("You are correct")
	case "Tia":
		fmt.Println("You are close")
	default:
		fmt.Println("Try again")
	}
}
