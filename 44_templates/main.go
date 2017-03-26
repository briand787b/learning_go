package main

import (
	"text/template"
	"net/http"
)

func displayHtml(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("tpl.gohtml")
	tpl.Execute(w, nil)
}

func main() {
	// tpl, _ := template.ParseFiles("tpl.gohtml")
	// tpl.Execute(os.Stdout, nil)
	http.HandleFunc("/", displayHtml)

	http.ListenAndServe(":8080", nil)
}
