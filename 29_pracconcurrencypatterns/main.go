package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(1, 2, 3, 4, 5)

	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int{
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}