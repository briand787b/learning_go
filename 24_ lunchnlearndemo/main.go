//This program is a really good demonstration of closures as well.
// Change expensiveCalculation(index) to expensiveCalculation(i) to
// demonstrate this.

package main

import (
	"fmt"
	"time"
)

func main() {
	sl := make([]int, 10)
	ch := make(chan bool)
	for i := 0; i < len(sl); i++ {
		go func(slice []int, index int, c chan bool) {
			slice[index] = expensiveCalculation(index)
			slice[index]++
			ch <- true
		}(sl, i, ch)
	}

	for i := 0; i < len(sl); i++ {
		<- ch
	}

	fmt.Println(sl)
}

func expensiveCalculation(index int) int{
	time.Sleep(1 * 1e9)
	return index + 42
}
