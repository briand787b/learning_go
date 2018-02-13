package main

import (
	"github.com/gorilla/mux"
)

const (
	// request keys for mux vars
	nameReqKey = "name"
)

func prepareRouter(r *mux.Router) {
	registerHelloHandlers(r)
	registerHomeHandlers(r)
}
