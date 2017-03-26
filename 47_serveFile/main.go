package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", chien)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "tpl.html", nil)
}

func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}