package main

import "fmt"

func fexp() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	wrapper := fexp()
	fmt.Println(wrapper())
	fmt.Print(wrapper())
	fmt.Println()
	wrapper2 := fexp()
	fmt.Println(wrapper2())
	fmt.Println(wrapper2())
	fmt.Println()
}
