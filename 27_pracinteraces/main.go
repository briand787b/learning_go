package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	age int
}

func (d Dog) Speak() string {
	return "woof"
}

func (d Dog) Bite() string {
	return "Ouch"
}


func main() {
	var monkey Dog
	// var anml Animal
	anml := Animal(monkey)
	fmt.Println(anml.Speak())
	bigMeans := Dog(anml)
	fmt.Println(bigMeans.Bite())
}