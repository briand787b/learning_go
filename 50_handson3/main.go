package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

type foo func(name int) int


func Handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "I see you connected")

	scnr := bufio.NewScanner(conn)
	for scnr.Scan() {
		fmt.Println(scnr.Text())
		io.WriteString(conn, "message received")
	}
	fmt.Println("connection closed")
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go Handle(c)
	}
}
