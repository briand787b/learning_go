package main

import (
	"net"
	"bufio"
	"log"
	"fmt"
	"io"
	"os"
)

func HandleConn(conn net.Conn) {
	fmt.Println("client connected")
	s := bufio.NewScanner(conn)
	for s.Scan(){
		io.WriteString(os.Stdout, s.Text())
	}
}

func main() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleConn(c)
	}
}

