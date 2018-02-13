package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	prepareRouter(r)

	// DELETE ME
	_, err := buildToken()
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println()
	_, err = GenerateJWT("junk", 5)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// n := negroni.Classic()

	// n.Use()

	go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/", 302))
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", r))
}
