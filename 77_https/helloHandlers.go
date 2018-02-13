package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func registerHelloHandlers(r *mux.Router) {
	r.HandleFunc("/hello", HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello/world", HelloWorldHandler).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/hello/world/{%s}", nameReqKey),
		HelloNamedWorldHandler).Methods(http.MethodGet)
}

// HelloHandler is the hello handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, with SSL/TLS"))
}

// HelloWorldHandler extends the hello to the entire world
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! From SSL/TLS"))
}

// HelloNamedWorldHandler handles hello-series routes that are named
func HelloNamedWorldHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)[nameReqKey]
	w.Write([]byte(fmt.Sprintf("Hello %s.  Welcome to my world", name)))
}
