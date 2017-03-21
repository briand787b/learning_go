package main

import (
	"fmt"
	"sync"
)

func main() {
	// set up the pipeline
	c := gen(2, 3)

	out1 := sq(c)
	out2 := sq(c)

	// consume the output
	for n := range merge(out1, out2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int{
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the input goroutines are
	// done.  This must start after the wg.Add call
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}