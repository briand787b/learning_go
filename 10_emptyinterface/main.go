package main

import (
	"fmt"
)

type cat interface {}

// empty object literal works in this function for interface
func doSomething(a interface{}) {
	fmt.Println("You called doSomething")
}

// does empty object literal work for concrete types
func doSomethingElse(c struct{}) {
	fmt.Println("You called doSomethingElse")
	fmt.Println("c: ", c)
}
// no!!!

func main() {
	var meow struct{}
	var mewo2 struct{}
	doSomething(meow)
	doSomethingElse(meow)
	doSomethingElse(mewo2)
}
