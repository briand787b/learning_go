package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("-----------------")

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
