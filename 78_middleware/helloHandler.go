package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func registerHelloHandlers(r *mux.Router) {
	r.HandleFunc("/", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
