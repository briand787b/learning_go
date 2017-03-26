package main

import "fmt"

func main() {
	defer finally()	
	hello()
	hello()
	hello()
	hello()
	hello()
	world()
}

func hello() {
	fmt.Print("hello")
}

func world() {
	fmt.Println("world")
}

func finally() {
	fmt.Println("I'm finally done with this program")
}
