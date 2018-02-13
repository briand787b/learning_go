package main

import (
	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	registerHelloHandlers(r)
	registerLoginHandlers(r)
}
