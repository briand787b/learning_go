package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"os"
)

type FirstName struct {
	Name string
}

const jsonStream = `
	[
		{ "name": "Brian" },
		{ "name": "Brian" }
	]
`

func main() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	var names []FirstName
	err := dec.Decode(&names)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(names)
}
