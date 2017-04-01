package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	rF, err := os.Open("old.pdf")
	if err != nil {
		fmt.Println("Error opening file ", err)
		os.Exit(1)
	}

	defer rF.Close()

	wF, err := os.Create("new.pdf")
	if err != nil {
		fmt.Println("Error creating new.pdf ", err)
		os.Exit(1)
	}

	defer wF.Close()

	io.Copy(wF, rF)
}

