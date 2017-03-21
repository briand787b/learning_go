package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go factorial(4)
	go factorial(5)
	go factorial(6)

	wg.Wait()

	// Going to attempt to do this a different way
	outChan := make(chan int)
	go factorialChan(4, outChan)
	go factorialChan(5, outChan)
	go factorialChan(6, outChan)

	for i := 0; i < 3; i++ {
		fmt.Println(<- outChan)
	}
}

func factorial(start int) {
	total := 1
	for i := start; i > 0; i-- {
		total *= i
	}

	fmt.Println(total)
	wg.Done()
}

func factorialChan(start int, ch chan int) {
	total := 1
	for i := start; i > 0; i-- {
		total *= i
	}

	ch <- total
}