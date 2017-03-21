package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"time"
	"os"
	"encoding/json"
)

type dbCredentials struct {
	Username string
	Password string
}

var dbUser dbCredentials

func main() {
	file, err := os.Open("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewDecoder(file).Decode(&dbUser); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(192.168.1.53:3306)/gregs_list", dbUser.Username, dbUser.Password))
	err = db.Ping()
	if err != nil {
		log.Fatal("database is unreachable: ", err)
	}

	rows, err := db.Query(`INSERT INTO users (username, password, lastLogin) VALUES (?, ?, ?)`, "brian", "b8003", time.Now())
	if err != nil {
		fmt.Println("error inserting rows into database: ", err)
	}

	fmt.Println(rows.Next())

	var un, pw string
	var ll time.Time

	for rows.Next() {
		rows.Scan(&un, &pw, &ll)
		fmt.Println(un, pw, ll)
	}
}
