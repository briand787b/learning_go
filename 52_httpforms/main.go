package main

import (
	"net/http"
	"html/template"
	"fmt"
)

var tpl *template.Template

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("tpl.html"))
	q := r.FormValue("q")
	tpl.Execute(w, "Your query"+q)
	// fmt.Fprint(w, "Your query: "+q)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:	"my-cookie",
		Value:	"SOME-VALUE",
	})
	fmt.Fprint(w, "cookie was set")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE: ", ck)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/setcookie/", setCookie)
	http.HandleFunc("/readcookie/", readCookie)
	http.ListenAndServe(":8080", nil)
}
