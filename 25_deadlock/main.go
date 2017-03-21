package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	fmt.Println("About to send 1 though channel")
	c <- 1
	fmt.Println("Sent 1 over the channel")
	fmt.Println(<-c)
}
