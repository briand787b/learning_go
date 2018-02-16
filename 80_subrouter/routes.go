package main

// import (
// 	"github.com/gorilla/mux"
// 	"github.com/urfave/negroni"
// 	"net/http"
// )

// func registerRoutes() *negroni.Negroni {
// 	r := mux.NewRouter()

// 	acc := r.PathPrefix("/account").Subrouter()
// 	acc.HandleFunc("/login", handleLogin)

// 	hlo := r.PathPrefix("/hello").Subrouter()
// 	hlo.HandleFunc("/secure", helloSecured)

// }

// func terminateUnauthenticatedRoutes(hf http.HandlerFunc) negroni.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 		w.Write([]byte("in the terminate unauthenticated routes function"))
// 		hf(w, r)
// 		return
// 	}
// }
