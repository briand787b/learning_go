package main

import (
	"fmt"
	"io"
	"os"
)

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
	io.Writer
}

type Request interface {
	io.Reader
}

type ConcreteHandler struct {}

func (c ConcreteHandler) handler(resWriter ResponseWriter, reqPtr *Request) {
	fmt.Println("The function was called")
}

func newHandler() *Handler {
	return new(Handler)
}

func CallHandler(h Handler) {

}

func main() {
	readerSatisfier := os.Stdin
	writerSatisfier := os.Stdout

	handlwer := new(ConcreteHandler)

}
