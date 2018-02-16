package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func registerHelloHandlers(r *mux.Router) {
	r.HandleFunc("/hello", hello).Methods("GET")
	r.HandleFunc("/hello/secure", helloSecured).Methods("GET")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func helloSecured(w http.ResponseWriter, r *http.Request) {
	tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		fmt.Println("error parsing token: ", err)
		return
	}

	ret := struct {
		Bearer string
		Token  *jwt.Token
	}{
		tokenString,
		token,
	}

	if err := json.NewEncoder(w).Encode(&ret); err != nil {
		http.Error(w, fmt.Sprintf("could not encode to json: %s", err), 500)
	}
}
