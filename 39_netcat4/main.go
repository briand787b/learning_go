package main

import (
	"net"
	"log"
	"io"
	"os"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("connection cannot be established: ", err)
	}
	done := make(chan struct{})
	doneWriting := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct {}{} // signal the main goroutine
	}()
	go func() {
		mustCopy(conn, os.Stdin)
		var tcpConn *net.TCPConn = conn.(*net.TCPConn)
		if err := tcpConn.CloseWrite(); err != nil {
			fmt.Println(err)
		}
		doneWriting<- struct{}{}
	}()
	<-doneWriting
	<-done
	conn.Close()

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
