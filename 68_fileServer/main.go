package main

import (
	"net/http"
	"os"
	"fmt"
)

func main() {
	var serveDir http.Dir = "/home/brian/Documents"
	fh := http.FileServer(serveDir)
	err := http.ListenAndServe(":8080", fh)
	if err != nil {
		fmt.Println("error starting server: ", err)
		os.Exit(1)
	}
}
