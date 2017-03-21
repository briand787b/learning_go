package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

type message struct {
	Name string
	Body string
	time int64
}

func main() {
	// ************************
	// Marshal the data
	// ************************
	fmt.Println("Program initiated")

	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	a := []byte(`{"Name":"Alice","Body":"Hello","time":1294706395881547000}`)
	fmt.Println(a)
	fmt.Println(b)

	c, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	// ***************************
	// Unmarshal the data
	// ***************************
	msg := message{Name: "Bob", Body: "blah", time: 2313141}
	err = json.Unmarshal(a, &msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)

	// **************************
	// Generic JSON
	// **************************
	var i interface{}
	fmt.Printf("type of interface before assignment: %T\n", i)
	i = "a string"
	i = 2011
	i = 2.77
	fmt.Printf("type of interface i: %T\n", i)
	r := i
	fmt.Printf("type of var copied from i: %T\n", r)
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is ", v * 2)
	case float64:
		fmt.Println("the reciprocal of i is ", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is ", v[:h]+v[:h])
	default:
		// i isn't one of the types above
	}
}