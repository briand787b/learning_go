package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"fmt"
	"encoding/json"
)

// username will not be used in this program
type credentials struct {
	Username string
	Password string
}

func init() {

}

func main() {
	var creds credentials

	f, err := os.Open("../credentials.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&creds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dsn := fmt.Sprintf("user=postgres dbname=test password=%s", creds.Password)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	rows, err := db.Query(`SELECT md5_sum FROM new_table;`)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	var sums []int
	defer rows.Close()
	for rows.Next() {
		var tmp int
		err = rows.Scan(&tmp)
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
		sums = append(sums, tmp)
	}

	fmt.Println(sums)
}
