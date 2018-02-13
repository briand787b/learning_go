package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	registerRoutes(r)

	n := negroni.Classic() // receive in variable when ready
	n.Use(negroni.HandlerFunc(practiceMiddleware))
	n.UseHandler(r)

	go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/", http.StatusFound))
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", n))
}
