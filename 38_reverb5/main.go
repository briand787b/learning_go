package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"os"
)

func HandleConn(conn net.Conn) {
	io.WriteString(conn, "You are now connected to the server")
	go io.Copy(os.Stdout, conn)
	go io.Copy(conn, os.Stdin)
	fmt.Println("code point was reached")
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error connecting to client")
			continue
		}
		go HandleConn(conn)
	}
}
