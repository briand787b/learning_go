package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func registerLoginHandlers(r *mux.Router) {
	r.HandleFunc("/login", handleLogin).Methods("POST")
}

// 	<form action="/login" method="POST" enctype="x-www-form-urlencoded">
//		<input type="text" name="username"/>
//		<input type="text" name="password"/>
//		<input type="submit"/>
//	</form>
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "cannot parse form", 500)
	}

	un := r.Form.Get("username")
	pw := r.Form.Get("password")
	fmt.Printf("username: %s | password: %s\n", un, pw)

	token, err := generateJWT()
	if err != nil {
		http.Error(w, "could not generate jwt", 500)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
}