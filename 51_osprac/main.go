package main

import (
	"os"
	"fmt"
)

func main() {
	_, err := os.Create("newfile.txt")
	if err != nil {
		fmt.Println(err)
	}


}
