package main

import (
	"fmt"
	"time"
)

func main() {

	sl := make([]int, 10, 20)
	fmt.Println(len(sl))

	mainDone := make(chan bool)
	for i := 0; i < len(sl); i++ {
		fmt.Println("Entering first for loop: ", i)
		fmt.Println("Length of sl: ", len(sl))
		if i < 10 {
			go func(comp chan bool) {
				fmt.Println("Entering go routine: ", i)
				done := make(chan bool)
				expensiveCalculation(sl, i, done)
				<-done
				sl[i] += 5
				comp <- true
			}(mainDone)
		}
	}

	for i := 0; i < len(sl); i++ {
		<- mainDone
	}
	fmt.Println(sl)
}

func expensiveCalculation(slice []int, index int, completion chan bool) {
	time.Sleep(1 * 1e9)
	slice[index] = 42
	completion <- true
}
