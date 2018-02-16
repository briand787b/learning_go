package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	authBase := mux.NewRouter()
	r.PathPrefix("/auth").Handler(negroni.New(
		// negroni.HandlerFunc(authmiddleware),
		negroni.Wrap(authBase),
	))

	auth := authBase.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", login)
	auth.HandleFunc("/logout", logout)

	postsBase := mux.NewRouter()
	r.PathPrefix("/posts").Handler(negroni.New(
		negroni.HandlerFunc(authmiddleware),
		negroni.Wrap(postsBase),
	))

	posts := postsBase.PathPrefix("/posts").Subrouter()
	posts.HandleFunc("/", getPosts)

	homeBase := mux.NewRouter()
	r.PathPrefix("/").Handler(negroni.New(
		negroni.HandlerFunc(authmiddleware),
		negroni.Wrap(homeBase),
	))

	home := homeBase.PathPrefix("/").Subrouter()
	home.HandleFunc("/home", homeRoute)

	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", r)
}

func authmiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("in auth middleware")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "error parsing form", 500)
	}

	if r.Form.Get("password") != "123" {
		w.WriteHeader(401)
		return
	}

	next(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "logout")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getting posts")
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
