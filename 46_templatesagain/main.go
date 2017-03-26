package main

import (
	"net/http"
	"text/template"
	"fmt"
)

var tpl *template.Template
var fn = template.FuncMap{ "ft": firstThree }

type Person struct {
	Name string
	Age int
}

func init() {
	template.Must(template.New("something").Funcs(fn).ParseGlob("templates/*"))

}

func firstThree(str string) string {
	return str[0:3]
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	Person := []Person{{"Brian", 22}, {"Tia", 18}, {"Chris", 33}}
	// tpl := template.Must(template.ParseGlob("templates/*"))
	fighter := struct {
		Person
		Weight int
	}{Person[0], 240}
	fmt.Println(fighter)
	tpl.ExecuteTemplate(w, "tpl.gohtml", fighter)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
