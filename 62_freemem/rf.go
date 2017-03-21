package main

import (
	"os/exec"
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"os"
	"encoding/json"
)

type dbCredentials struct {
	Username string
	Password string
}

var dbUser dbCredentials

func execProg() {
	free, err := exec.Command("grep", "MemFree", "/proc/meminfo").Output()
	if err != nil {
		fmt.Println("error executing? command: ", err)
		return
	}
	fmt.Println("free memory right now is ", string(free))
	fmt.Println("\n", free)
	freeS := strings.TrimSpace(strings.TrimPrefix(string(free), "MemFree:"))
	freeS = strings.TrimSpace(strings.TrimSuffix(freeS, "kB"))
	fmt.Println("This is the free memory available: ", freeS)
	fmt.Println("Byte representation of space:", []byte(" "))
	fmt.Println("Byte representation of newline:", []byte("\n"))
	fmt.Println("Free memory available as byte slice ", []byte(freeS))

	file, err := os.Open("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewDecoder(file).Decode(&dbUser); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(192.168.1.53:3306)/gregs_list", dbUser.Username, dbUser.Password))
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database ", err)
	} else {
		fmt.Println("success connecting to database!")
	}
	_, err = db.Query(`INSERT INTO ram_usage (ram_used, date) VALUES (?, ?)`, freeS, time.Now())
	if err != nil {
		fmt.Println("error inserting into database ", err)
	}
}

func main() {
	for true {
		execProg()
		time.Sleep(1 * time.Second)
	}
}