package main

import "fmt"

func main() {
	myVar := 1

	b := func() {
		fmt.Println(myVar) // prints out 0
	}

	myVar = 0

	a := func() {
		myVar := 2
		fmt.Printf("myVar is %v in a\n", myVar)
		b()
	}
	a()
}
