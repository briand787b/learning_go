package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//func (c circle) area() float64 {
//	return c.radius * c.radius * math.Pi
//}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	// try to use value method receiver
	c := circle{5}
	info(&c)

	fmt.Println(c.area())
	fmt.Println((&c).area())
}