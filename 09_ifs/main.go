package main

import "fmt"

func main() {
	fmt.Println("Enter in a number")
	var res string
	_, err := fmt.Scanln(&res)
	if err != nil {
		fmt.Println("Incorrect input")
	}
}
