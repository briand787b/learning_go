package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	err := ioutil.WriteFile("newfile.txt", []byte("content"), 0600)
	fmt.Println(err)
}
