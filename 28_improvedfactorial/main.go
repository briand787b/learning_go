package main

import (
	"fmt"
)

func main() {
	// supply function with numbers you want to be factorialed
	nums := []int{2, 3, 4, 5, 6, 7, 8}


	// collect the values
	for n := range factorial(gen(nums...)) {
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

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}

