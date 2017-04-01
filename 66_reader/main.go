package main

import (
	"os"
	"fmt"
	"io"
)

// This file demonstrates that subsequent reads on a stream
// (in this case the contents of a file) resume at the last
// location the read happened.

func main() {
	fd, err := os.Open("tempfile.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
	}


	defer fd.Close()

	buf := make([]byte, 4096)
	tmp := make([]byte, 256)

	for {
		n, err := fd.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error occurred before reading all of file: ", err)
			}

			break
		}

		buf = append(buf, tmp[:n]...)
	}

	fmt.Println("all text: ", string(buf))
}
