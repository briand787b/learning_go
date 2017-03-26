package main

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
)

type Server struct {
	Server string
	Ram int
	DriveSpace int
}

func main() {
	fmt.Println("Starting program")

	file, err := os.Open("/home/brian/dev/godev/src/github.com/briand787b/udemy/pracJson/config.json")
	if err != nil {
		log.Fatal(err)
	}

	var serverStats Server

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&serverStats); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s %v %v\n", serverStats.Server, serverStats.Ram, serverStats.DriveSpace)

	fmt.Println("Ending Program")

	//var serverReports Server
	//if err = json.Unmarshal(file, &serverReports); err != nil {
	//	fmt.Println(err)
	//}
}