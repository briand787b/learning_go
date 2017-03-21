package main

import (
	"fmt"
)

func pump() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func suction(c chan int) {
	for i := range c {
		
	}
}

func main() {
	fourChan := pump()
	// do stuff()

}