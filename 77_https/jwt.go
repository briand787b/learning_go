package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// location of the files used for signing and verification
const (
	privKeyPath = "app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
	secret      = "unsafesecret"
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

// func buildToken() {
// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"foo": "bar",
// 		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	hmacSampleSecret := "this_is_my_unsecuresecret"
// 	tokenString, err := token.SignedString(hmacSampleSecret)

// 	fmt.Println(tokenString, err)
// }

func buildToken() (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = jwt.StandardClaims{
		Subject:   "bar",
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
	}

	tokenString, err := t.SignedString(signKey)
	fmt.Println(tokenString)
	return tokenString, err
}

// func JWTMiddleware() negroni.HandlerFunc {
// 	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

// 	})
// }

//GenerateJWT generates JWT web token
func GenerateJWT(user string, customStruct interface{}) (string, error) {

	mySigningKey := []byte(secret)

	// Create the Claims
	claims := jwt.MapClaims{
		"Subject": "bar",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	fmt.Println(tokenString)
	return tokenString, err
}

func ParseToken(tokenString string) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return signKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
