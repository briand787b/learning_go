package main

import (
	"os"
	"io"
	"encoding/json"
	"fmt"
	"bytes"
)

type person struct {
	Name string
	Age int
}

// the purpose of this file is to practice using stdin and stdout
func main() {
	var jsonInput person
	err := json.NewDecoder(os.Stdin).Decode(&jsonInput)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(os.Stdout, "something\n")
	fmt.Println(jsonInput)

	// r & w, i assume, must be assigned after this function runs
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Name())
	fmt.Println(w.Name())
	fmt.Println("\n")
	fmt.Println(r.Stat())

	fmt.Println("Starting buffer practice")
	var b bytes.Buffer
	fmt.Println("b's capacity", b.Cap())
	sl := []byte("this is the string")
	fmt.Println("length of string: ", len(sl))
	b.Write(sl)
	fmt.Println("b's capacity", b.Cap())



}
