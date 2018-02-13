package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func registerHomeHandlers(r *mux.Router) {
	r.HandleFunc("/", handleHome).Methods(http.MethodGet)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome home!"))
}
