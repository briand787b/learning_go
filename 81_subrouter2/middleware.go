package main

import (
	"fmt"
	"net/http"
)

type practiceMiddleware struct{}

// negroni
func (pm *practiceMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("in the practice middleware")

	if r.URL.Query().Get("password") == "123" {
		next(w, r)
	} else {
		http.Error(w, "nah-ah-ah, you didn't say the magic word", 401)
		// http.Redirect(w, r, "/account/login", http.StatusFound)
	}
}

// non-negroni
func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("password") == "123" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "invalid password", http.StatusUnauthorized)
		}
	})
}
