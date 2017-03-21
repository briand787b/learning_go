package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonIn = `{
		"First": "Ed",
		"Last": "Same",
		"Age": 28
	}`

	type Person struct {
		First string
		Last  string
		Age   int
	}

	var p Person
	err := json.NewDecoder(strings.NewReader(jsonIn)).Decode(&p)
	if err != nil {
		fmt.Println("You fucked up")
	}
	fmt.Println(p)

	const jsonStream = `
    {"Name": "HHH", "Text": "Knock KNock"}
    {"Name": "Ed", "Text": "Who's there?"}
    {"Name": "HHH", "Text": "The HHH"}
    {"Name": "Ed", "Text": "Uh Oh"}
  `
	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
