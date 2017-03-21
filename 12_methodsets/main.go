package main

import (
	"fmt"
)

type car struct {
	make string
	model string
}

func (c *car) changeBadge(newMake string) {
	c.make = newMake
}

func (c car) getMake() string {
	return c.make
}

func main() {
	var brianCar = car{"Toyota", "Camry"}
	brianCar.changeBadge("Ferrari")
	fmt.Println(brianCar.getMake())
}
