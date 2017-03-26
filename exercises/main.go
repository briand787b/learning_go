package main

import "fmt"

func main() {
	fmt.Print("Please enter your name: ")
	var response string
	fmt.Scan(&response)
	fmt.Println("Hello, ", response)

	fmt.Print("Enter a large and a small number: ")
	var intOne, intTwo int
	fmt.Scan(&intOne, &intTwo)
	fmt.Println(intOne % intTwo)

	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
