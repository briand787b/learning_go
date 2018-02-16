package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	privKeyPath = "app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

// remove this function if code ever goes
// to production.
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

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
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
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
	token, err := generateJWT()
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get token: %s", err), 500)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer: %s", token))
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

// bool indicates if token from passed request is valid
func validateToken(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	fmt.Println("auth header: ", authHeader)

	if authHeader != "" {
		return true
	}

	return false
}

func generateJWT() (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = jwt.StandardClaims{
		Subject:   "bar",
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
	}

	tokenString, err := t.SignedString(signKey)
	fmt.Println(tokenString)
	return tokenString, err
}
