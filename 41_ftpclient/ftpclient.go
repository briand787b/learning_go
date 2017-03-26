package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(conn, os.Stdin)
}