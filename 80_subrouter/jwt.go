package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// remove this function if code ever goes
// to production.
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
