package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
	"os"
	"log"
)

type dbCredentials struct {
	Username string
	Password string
}

var db *sql.DB
var dbUser dbCredentials

func main() {
	file, err := os.Open("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewDecoder(file).Decode(&dbUser); err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(192.168.1.53:3306)/foo", dbUser.Username, dbUser.Password))
	if err != nil {
		fmt.Println("There was an error: ", err)
	} else {
		fmt.Println("Connection succeeded")
	}

	// check(err)
	err = db.Ping()
	if err != nil {
		fmt.Println("error connecting to database: ", err)
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM states;`)
	if err != nil {
		fmt.Println("error querying database: ", err)
	}

	var s, c1, c2, c3  string
	s = "RETRIEVED RECORDS: \n"
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3)
		if err != nil {
			fmt.Println("error scanning row", err)
		}
		s += c1 + c2 + c3 + "\n"
	}
	fmt.Println(s)

	stmt, err := db.Prepare(`INSERT INTO states (state, population) VALUES ("Indiana", "6633053");`)
	if err != nil {
		fmt.Println("error inserting into the database: ", err)
	}
	r, err := stmt.Exec()
	if err != nil {
		fmt.Println("error in executing insert: ", err)
	}
	n, err := r.RowsAffected()
	if err != nil {
		fmt.Println("i don't know what happened: ", err)
	}
	fmt.Println("INSERTED RECORD: ", n)
}
