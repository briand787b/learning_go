package main

import (
	"net/http"
	"html/template"
	"fmt"
	"io/ioutil"
)

var tpl *template.Template

func serveRoot(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("tpl.html"))
	var s string
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile: ", f, "\nheader: ", h)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	tpl.Execute(w, s)
}

func serveDog(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("tpl.html"))
	tpl.Execute(w, "ran dog")
}

func serveMe(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("tpl.html"))
	tpl.Execute(w, "hello Brian")
}

func main() {
	svmx := http.NewServeMux()
	svmx.Handle("/", http.HandlerFunc(serveRoot))
	svmx.Handle("/dog/", http.HandlerFunc(serveDog))
	svmx.Handle("/me/", http.HandlerFunc(serveMe))

	http.ListenAndServe(":8080", svmx)
}
