package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Method was POST")
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error occurred parsing file from form ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		out, err := os.Create(header.Filename)
		if err != nil {
			fmt.Println("Error occurred opening file ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println("Error occurred copying uploaded file to disk ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s was completely written to disk!", header.Filename)
		return
	}

	fmt.Println("Method was not POST")
	tpl.ExecuteTemplate(w, "upload.html", nil)
}

func main() {
	http.HandleFunc("/", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
