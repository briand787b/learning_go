package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(1 * 1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func receiveData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s ", input)
	}
	fmt.Println("leaving receiveData()")
}
