package main

import (
	"net/http"
	"text/template"
)

func serveBrianPage(w http.ResponseWriter, r *http.Request) {
	lngs := []string{"perl", "python", "go", "ruby", "c++"}

	tmp := template.Must(template.ParseGlob("templates/***"))
	tmp.ExecuteTemplate(w, "tpl.gohtml", lngs)
}

func main() {
	http.HandleFunc("/", serveBrianPage)

	http.ListenAndServe(":8080", nil)
}
