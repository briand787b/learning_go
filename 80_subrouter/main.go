package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	acc := r.PathPrefix("/account").Subrouter()
	acc.HandleFunc("/login", handleLogin)

	hlo := r.PathPrefix("/hello").Subrouter()
	hlo.HandleFunc("/secure", helloSecured)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", r)
	httpMux.Handle("/account/", negroni.New(
		negroni.Wrap(r),
	))
	httpMux.Handle("/hello/", negroni.New(
		negroni.Handler(&practiceMiddleware{}),
		negroni.Wrap(r),
	))

	n := negroni.Classic()
	n.UseHandler(httpMux)

	go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/", http.StatusFound))
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", n))
}
