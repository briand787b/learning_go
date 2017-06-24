package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a int) {
			fmt.Println(a)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
