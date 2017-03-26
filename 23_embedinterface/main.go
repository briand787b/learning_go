package main

import (
	"io"
	"fmt"
	"os"
)

type echoer struct {
	io.Reader
}

func (e *echoer) Read(p []byte) (int, error) {
	return 5, *new(error)
}

func ReadUpToN(r io.Reader, n int) {
	fmt.Println("Read up to n was called")
}

func main() {
	ReadUpToN(os.Stdin, 5)

	var ec echoer
	ec.Reader = os.Stdin
	ReadUpToN(&ec, 4)
}