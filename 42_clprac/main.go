package main

import (
	"os"
	"flag"
	"fmt"
	"log"
)

func main() {
	depth := flag.Int("depth", 0, "the maximum number of links to pursue")
	concurrent := flag.Bool("concurrent", true, "program will utilize concurrency if not set to false")

	if len(os.Args) < 2 {
		log.Fatal("You dun' goofed")
	}
	fmt.Println("\n")
	flag.Parse()
	fmt.Println(*depth)
	fmt.Println(*concurrent)
}


