package main

import (
	"fmt"
)

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	t, ok := val.(string)
	if ok {
		fmt.Printf("%T\n", t)
	} else {
		fmt.Println("Invalid assertion: ", ok)
	}
}
