package main

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
	"log"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"os"
	"encoding/json"
)

type user struct {
	UserName string
	Password string
}

type dbConnection struct {

}

type credentials struct {
	Username string
	Password string
}

// Map sessions to validated users
var usrSessn = make(map[string]*user)
var sIndex int
var db *sql.DB

// Global template variable
var tpl *template.Template

func init() {
	creds, err := os.Open("../credentials.json")
	if err != nil {
		fmt.Println("Error opening credentials.json: ", err)
		os.Exit(1)
	}
	defer creds.Close()

	var dbCreds credentials
	dec := json.NewDecoder(creds)
	err = dec.Decode(&dbCreds)
	if err != nil {
		fmt.Println("Error parsing the credentials.json file: ", err)
		os.Exit(1)
	}

	// Connect to the database
	dsn := fmt.Sprintf("%s:%s/second_db", dbCreds.Username, dbCreds.Password)
	db, _ = sql.Open("mysql", dsn)
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	// Parse templates
	tpl = template.Must(template.ParseGlob("*.html"))
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// Print the ip address of the user to the console
	userIP := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Println(userIP)
	rows, err := db.Query(`INSERT INTO request_ips (ip, date) VALUES (?, ?)`, userIP, time.Now())
	if err != nil {
		fmt.Println("error querying database: ", err)
	} else {
		fmt.Println("rows from db: ", rows)
	}

	// Work on form
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hello %s", c.Value)
	un, ok := usrSessn[c.Value]
	if ok {
		fmt.Fprintf(w, "Hello %s", un)
	} else {
		fmt.Fprint(w, "Hey, we couldn't find you in the map")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		c, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Error reading cookie ", err)
		} else {
			fmt.Println("cookie: ", c.Value)
		}
		err = tpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			fmt.Print("error: ", err)
		}
	} else {
		r.ParseForm()
		fmt.Println("This is the better way to get the value from the form, since url.Values is a map of strings to slice of strings", r.Form.Get("username"))
		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form.Get("password"))
		sIndex++
		fmt.Println("This is the value of sIndex: ", string(sIndex))
		http.SetCookie(w, &http.Cookie{Name:"session", Value:strconv.Itoa(sIndex)})
		usrSessn[string(sIndex)] = &user{UserName:r.Form["username"][0], Password:r.Form["password"][0]}
	}
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled favicon request")
}

func main() {
	defer db.Close()
	http.HandleFunc("/login", login)
	http.HandleFunc("/favicon.ico/", handleFavicon)
	http.HandleFunc("/", sayhelloName)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
