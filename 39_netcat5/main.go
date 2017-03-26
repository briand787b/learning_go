package main

import (
	"net"
	"log"
	"io"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.22:8080")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(conn, "hello")
	if err != nil {
		fmt.Println("Error writing to connection")
	}

	ch := make(chan struct{})
	go func() {
		io.Copy(conn, os.Stdin)
		ch <- struct{}{}
	}()
	go func() {
		io.Copy(os.Stdout, conn)
		ch <- struct{}{}
	}()
	<-ch
}
