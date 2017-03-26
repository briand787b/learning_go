package main

import "fmt"

var a, c, d string = "foo", "bar", "baz"

func main() {
	b := "bar"

	fmt.Printf("%v", b)

	acc()
}

func acc() {
	fmt.Println(d)
}
