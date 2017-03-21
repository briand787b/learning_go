package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"fmt"
)

type Page struct {
	Title string
	Body []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(name string) (*Page, error) {
	filename := name + ".txt"
	xb, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title:name, Body:xb}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fn := r.URL.Path[len("/view/"):]
	pg, err := loadPage(fn)
	if err != nil {
		http.Redirect(w, r, "/edit/"+fn, http.StatusFound)
		return
	}
	renderTemplate(w, "view", pg)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fn := r.URL.Path[len("/save/"):]
	fmt.Println(fn)
	pg := &Page{Title:fn, Body:[]byte(r.PostFormValue("body"))}
	if err := pg.save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+fn, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title:title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)

	http.ListenAndServe(":8080", nil)
}
