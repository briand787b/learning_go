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
	"flag"
	"errors"
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
	Id int
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

type dbConnection struct {
	Location string
	IP string
	Port string
	Database string
}

var tpl *template.Template
var db *sql.DB
var dbUser dbCredentials
var dbConxns []dbConnection
var srvrLocation string

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))

	credFile, err := os.Open("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}

	defer credFile.Close()
	if err = json.NewDecoder(credFile).Decode(&dbUser); err != nil {
		log.Fatal(err)
	}

	conxnFile, err := os.Open("./dbConnections.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.NewDecoder(conxnFile).Decode(&dbConxns); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&srvrLocation, "location", "home", "consult dbConnections.json for predefined " +
		"database locations to connect to e.g. home, work")
}

// This method searches the database for the session id
func sessionHelper(w http.ResponseWriter, r *http.Request) error {
	ck, err := r.Cookie("sessionId")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return errors.New(" could not find sessionId in cookie, redirecting to login ")
	}

	rows, err := db.Query(`SELECT u.username, u.password FROM sessions s INNER JOIN users u ON s.user_id = u.id WHERE s.id = ?`, ck.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return errors.New(" error querying database for username based on session id from cookie ")
	}
	var un, pw string
	for rows.Next() {
		err = rows.Scan(&un, &pw)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return errors.New("error scanning username into variable");
		}
	}

	return nil
}

func viewAllHandler(w http.ResponseWriter, r *http.Request) {
	if err := sessionHelper(w, r); err != nil {
		fmt.Println(err)
		return
	}

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
	if err := sessionHelper(w, r); err != nil {
		fmt.Println(err)
		return
	}

	drinkName := r.URL.Path[len("/drink/"):]
	rows, err := db.Query("SELECT d.drink_name, d.main, d.amount1, d.second, d.amount2, d.directions FROM easy_drinks d WHERE d.drink_name = ?", drinkName)
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

	if len(drinks) < 1 {
		fmt.Println("Error: number of drinks returned from database is less than one")
		http.Redirect(w, r, "/drinks", http.StatusNotFound)
	}

	tpl.ExecuteTemplate(w, "drink.html", drinks[0])
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("method on login was POST")

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
			// user is authenticated, give them a cookie
			fmt.Println("Nonzero records returned, username password pair found in database")
			sessionId := uuid.NewV4()
			ck := http.Cookie { Name: "sessionId", Value: sessionId.String(), Path: "/" }
			http.SetCookie(w, &ck)

			// store cookie and user association in database
			var userId int
			if err = rows.Scan(&userId); err != nil {
				// grant user access anyway since row->struct scan could just be server error
				fmt.Print(" error scanning row from db into user struct ")
			} else {
				_, err = db.Query(`INSERT INTO sessions (id, user_id, createdDateTime) VALUES (?, ?, ?)`, sessionId, userId, time.Now())
			}

			http.Redirect(w, r, "/drinks/", http.StatusFound)
			return
		} else {
			fmt.Println(" There were no rows in the database corresponding to that username password pair ")

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

	fmt.Println("Method on login was NOT POST")

	ck, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Print(" could not find sessionId in cookie, rendering template ")
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

	// user has a valid token, send them to the main page
	http.Redirect(w, r, "/drinks", http.StatusSeeOther)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Entering signUpHandler method with method: NOT POST")
		tpl.ExecuteTemplate(w, "signup.html", nil)
		return
	}

	fmt.Println("Entering signUpHandler with method POST")

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
	flag.Parse()
	if srvrLocation == "" {
		fmt.Println("Warning, you are choosing the default database location.  If you have connection issues," +
			" please consult dbConnections.json")
	}

	// find the database location which matches the name in srvrLocation flag
	// since i don't expect there to be many db locations, i will use a simple search
	var dbConxn dbConnection
	for _, dbc := range dbConxns {
		if dbc.Location == srvrLocation {
			dbConxn = dbc
		}
	}

	cnxnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser.Username,
		dbUser.Password,
		dbConxn.IP,
		dbConxn.Port,
		dbConxn.Database)
	fmt.Println("Connecting to ", cnxnString)
	db, _ = sql.Open("mysql", cnxnString)
	if err := db.Ping(); err != nil {
		log.Fatal(" error connecting to database ", err)
	}

	defer db.Close()

	http.HandleFunc("/signup/", signUpHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/drinks/", viewAllHandler)
	http.HandleFunc("/drink/", viewSingleHandler)
	http.Handle("/favicon.ico/", http.NotFoundHandler())
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}