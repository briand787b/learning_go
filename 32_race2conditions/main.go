package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	sl := make([]int, 10, 20)
	wg.Add(len(sl))
	for i := 0; i < len(sl); i++ {
		go expensiveCalculation(sl, i)
		sl[i] += 5
	}

	wg.Wait()
	fmt.Println(sl)
}

func expensiveCalculation(slice []int, index int) {
	time.Sleep(1 * 1e9)
	slice[index] = 42
	wg.Done()
}
