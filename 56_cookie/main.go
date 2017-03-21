package main

import (
	"net/http"
	"io"
	"fmt"
	"github.com/satori/go.uuid"
)

var uniqueVisitors = make(map[string]int)

func cookieListen(w http.ResponseWriter, r *http.Request) {
	var msg string

	cFoo, err := r.Cookie("Foo")
	if err != nil {
		io.WriteString(w, "Foo cookie was not set")
		return
	}
	uniqueVisitors[cFoo.String()]++
	fmt.Println(uniqueVisitors[cFoo.String()])
	msg += fmt.Sprintln("Foo cookie: ", cFoo.Value)

	cSession, err := r.Cookie("Session")
	if err != nil {
		io.WriteString(w, "Session was not set")
		return
	}
	msg += fmt.Sprintln("Session cookie: ", cSession.Value)

	io.WriteString(w, msg)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewV4()
	http.SetCookie(w, &http.Cookie{Name:"Session", Value:id.String()})
	io.WriteString(w, "Hey the cookie should be set")
}

func main() {
	http.HandleFunc("/listen/", cookieListen)
	http.HandleFunc("/", setCookie)

	http.ListenAndServe(":8080", nil)
}
