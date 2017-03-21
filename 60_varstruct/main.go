package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"time"
	"log"
	"github.com/satori/go.uuid"
	"os"
	"encoding/json"
)

type drink struct {
	DrinkName string
	Main string
	Amount float64
	Second string
	Amount2 float64
	Directions string
}

type user struct {
	UserName string
	Password string
	LastLoggedIn time.Time
}

type session struct {
	SessionId int
	TimeIssued time.Time
}

type dbCredentials struct {
	Username string
	Password string
}

var tpl *template.Template
var db *sql.DB
var dbUser dbCredentials

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))

	file, err := os.Open("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewDecoder(file).Decode(&dbUser); err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(192.168.1.53:3306)/gregs_list", dbUser.Username, dbUser.Password))
	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to database", err)
	}
}

// This method searches the database for the session id
func sessionHelper() {

}

// This method
func validationHelper() {

}

func viewAllHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM easy_drinks`)
	if err != nil {
		fmt.Print("error: ", err)
	}

	drinks := []drink{}
	var rec drink
	for rows.Next() {
		err = rows.Scan(&rec.DrinkName, &rec.Main, &rec.Amount, &rec.Second, &rec.Amount2, &rec.Directions)
		if err != nil {
			fmt.Println("error: ", err, ".  ")
		}
		drinks = append(drinks, rec)
	}

	tpl.ExecuteTemplate(w, "drinks.html", drinks)
}

func viewSingleHandler(w http.ResponseWriter, r *http.Request)  {
	ck, err := r.Cookie("session_id")
	if err != nil {
		fmt.Print("error reading cookie: ", err)
		http.Redirect(w, r, "/login/", http.StatusForbidden)
		return
	}

	rows, err := db.Query(`SELECT u.username FROM sessions s INNER JOIN users u ON s.user_id = u.id WHERE s.id = ?`, ck.Value)
	if err != nil {
		fmt.Printf("error selecting username from uuid %s: %s", ck.Value, err)
		http.Redirect(w, r, "/login/", http.StatusForbidden)
		return
	}

	var uname string
	for rows.Next() {
		err = rows.Scan(&uname)
		if err != nil {
			fmt.Printf("error converting username from database to string variable: %s", err)
			http.Redirect(w, r, "/login/", http.StatusForbidden)
			return
		}

	}

	drinkName := r.URL.Path[len("/drink/"):]
	rows, err = db.Query("SELECT * FROM easy_drinks WHERE drink_name = ?", drinkName)
	if err != nil {
		fmt.Println("error querying drink name from database: ", err)
		http.Error(w, "Error, drink not found", http.StatusNotFound)
	}

	drinks := []drink{}
	var rec drink
	for rows.Next() {
		err = rows.Scan(&rec.DrinkName, &rec.Main, &rec.Amount, &rec.Second, &rec.Amount2, &rec.Directions)
		if err != nil {
			fmt.Println("error scanning row into struct: ", err)
		}
		drinks = append(drinks, rec)
	}

	tpl.ExecuteTemplate(w, "drinks.html", drinks)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Print("method on login was POST")

		uname := r.FormValue("username")
		if uname == "" {
			fmt.Print(" error getting username from form")
			tpl.ExecuteTemplate(w, "login.html", struct{
				username string
				password string
				err string
			}{
				"",
				"",
				"error getting username from form",
			})

			// http.Error(w, "Error retrieving username from form", http.StatusInternalServerError)
			return
		}

		pword := r.FormValue("password")
		if pword == "" {
			fmt.Print(" error getting password from form")
			tpl.ExecuteTemplate(w, "login.html", struct{
				Username string
				Password string
				Err string
			}{
				"",
				"",
				"error getting password from form",
			})

			//http.Error(w, "Error retrieving values from form", http.StatusInternalServerError)
			return
		}

		rows, err := db.Query(`SELECT u.id FROM users u WHERE u.username = ? AND u.password = ?`, uname, pword)
		if err != nil {
			fmt.Print(" error retrieving username and password from, not sure if the error condition just means no results though")
			http.Error(w, "Error connecting to database", http.StatusInternalServerError)

			// See what happens for this return after http error
			// I don't know how http errors work
			return
		}

		if rows.Next() {
			fmt.Print("nonzero records returned, username password pair found in database")
			http.Redirect(w, r, "/drinks/", http.StatusFound)
			return
		} else {
			fmt.Print(" There were no rows in the database corresponding to that username password pair ")

			// This and the block below it are a huge mess that i don't understand
			// i need to fix this/understand it because it makes no sense why it works
			// the way it is now, but doesn't work otherwise
			// http.Redirect(w, r, "/login", http.StatusFound)

			tpl.ExecuteTemplate(w, "login.html", struct{
				Username string
				Password string
				Err string
			}{
				uname,
				pword,
				"Sorry, we couldn't find you in the database",
			})
			return
		}
	}

	fmt.Print("method on login was NOT POST")

	ck, err := r.Cookie("session_id")
	if err != nil {
		fmt.Print(" could not find session_id in cookie, rendering template ")
		tpl.ExecuteTemplate(w, "login.html", nil)
		return
	}

	rows, err := db.Query(`SELECT u.username, u.password FROM sessions s INNER JOIN users u ON s.user_id = u.id WHERE s.id = ?`, ck.Value)
	if err != nil {
		fmt.Print("error querying database for username based on session id from cookie", err)
		tpl.ExecuteTemplate(w, "login.html", nil)
		return
	}
	var un, pw string
	for rows.Next() {
		err = rows.Scan(&un, &pw)
		if err != nil {
			fmt.Print("error scanning username into variable", err)
			tpl.ExecuteTemplate(w, "login.html", nil)
			return
		}
	}

	tpl.ExecuteTemplate(w, "login.html", struct{
		username string
		password string
		err string
	}{
		un,
		pw,
		"",
	})
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("entering signUpHandler method with method: NOT POST")
		tpl.ExecuteTemplate(w, "signup.html", nil)
		return
	}

	fmt.Println("entering signUpHandler with method POST")

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form: ", err)
		tpl.ExecuteTemplate(w, "signup.html", err)
	}

	un := r.FormValue("username")
	pw := r.FormValue("password")

	// Check that username does not exist in the database already
	rows, err := db.Query(`SELECT username FROM users WHERE username = ?`, un)
	if err != nil {
		log.Fatal("errors running query on databse", err)
	}
	if rows.Next() {
		tpl.ExecuteTemplate(w, "signup.html", "username is already taken")
		return
	}

	_, err = db.Query(`INSERT INTO users (username, password, createdDateTime) VALUES (?, ?, ?)`, un, pw, time.Now())
	if err != nil {
		fmt.Println("error inserting row into databse")
		tpl.ExecuteTemplate(w, "signup.html", "error storing user")
		return
	}

	uuid := uuid.NewV4()
	_, err = db.Query(`INSERT INTO sessions (id, user_id, createdDateTime) VALUES (?, (SELECT u.id FROM users u WHERE u.username = ?), ?)`, uuid.String(), un, time.Now())
	if err != nil {
		fmt.Print("error inserting session row into sessions table", err)
		tpl.ExecuteTemplate(w, "signup.html", "error storing user")
	}
	ck := &http.Cookie{ Name: "session_id", Value: uuid.String(), Path: "/"}
	http.SetCookie(w, ck)

	http.Redirect(w, r, "/drinks", http.StatusFound)
}

func main() {
	defer db.Close()

	http.HandleFunc("/signup/", signUpHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/drinks/", viewAllHandler)
	http.HandleFunc("/drink/", viewSingleHandler)
	http.Handle("/favicon.ico/", http.NotFoundHandler())
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}