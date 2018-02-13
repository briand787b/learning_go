package main

import (
	"fmt"
	"net/http"
)

func practiceMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("in the practice middleware")

	if r.URL.Query().Get("password") == "123" {
		next(w, r)
	} else {
		http.Error(w, "nah-ah-ah, you didn't say the magic word", 401)
	}
}
