package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	fmt.Println("entering main()")
	out := make(chan int)
	fmt.Println("sending 2 on out channel")
	out <- 2
	fmt.Println("initiating go routine f1")
	go f1(out)
}