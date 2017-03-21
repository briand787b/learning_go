package main

import (
	"text/template"
	"net/http"
	"fmt"
)

var tpl *template.Template
var fn = template.FuncMap{
	"uc": ToUpper,
	"ft": FirstThree,
}

type Job struct {
	Title string
	Salary int
	Company string
}

func ToUpper(str string) string {
	var newStr []byte
	for i := range str {
		newStr = append(newStr, byte(int(str[i]) + 32))
	}
	return string(newStr)


}

func FirstThree(str string) string {
	return str[0:3]
}

func ServeTemplates(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.New("tmplt").Funcs(fn).ParseGlob("templates/*"))
	tpl.ExecuteTemplate(w, "tpl2.html", struct {
			Name string
			Age int
			Job
		}{
			"Brian",
			26,
			Job { "Programmer", 50000, "Elegant Software Solutions" },
	})
}

func main() {
	http.HandleFunc("/", ServeTemplates)
	http.ListenAndServe(":8080", nil)
}
