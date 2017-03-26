package main

import "net/http"

func serveMain(w http.ResponseWriter, r *http.Request) {

}

func serveSecond(w http.ResponseWriter, r *http.Request) {

}

func serveThird(w http.ResponseWriter, r *http.Request) {

}

func serveFourth(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", serveMain)
}
