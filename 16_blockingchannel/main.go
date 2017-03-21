package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	fmt.Println("Initiated go routines, waiting one second")
	time.Sleep(16 * 1e9)
	fmt.Println("Waited 16 seconds in main()")
	//fmt.Println(<-ch1)
	fmt.Println("Leaving main()")
}

func pump(ch chan int) {
	fmt.Println("Entering pump(), waiting for 15 seconds")
	time.Sleep(15 * 1e9)
	fmt.Println("Waited 15 seconds")
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}