package main

import (
	"net/http"
	"log"
	"io"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<a href='http://localhost:8080/page'>link</a>`)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<a href='http://localhost:8080/nextpage'>link1</a><br/><a href='http://localhost:8080/anotherpage'>link2</a>`)
}

func serveNext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world from serveNext function")
}

func serveAgain(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<a href='http://localhost:8080/thirdpage'>link3</a>")
}

func serveFourth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<a href='http://localhost:8080/fourthpage'>link4</a>")
}

func serveFifth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<hello world")
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/page", servePage)
	http.HandleFunc("/nextpage", serveNext)
	http.HandleFunc("/anotherpage", serveAgain)
	http.HandleFunc("/thirdpage", serveFourth)
	http.HandleFunc("/fourthpage", serveFifth)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}