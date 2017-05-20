package main

import (
	"github.com/briand787b/pgInit"
	"fmt"
	"os"
)

func main() {
	db, err := pgInit.ConnectDefault("test")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := db.Query(`SELECT md5_sum FROM new_table;`)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var sums []int
	var tmp int
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		sums = append(sums, tmp)
	}

	fmt.Println(sums)
}
