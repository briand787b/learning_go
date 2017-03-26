package main

import (
	"os"
	"log"
	"fmt"
)

type File struct {
	*int
}

func NewFilePtr() *File {
	a := 4
	return &File{&a}
}


func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fd := f
	fmt.Println(f)
	fmt.Println(&fd)
	fmt.Println(*f)

	fmt.Println("************************************\n")

	FilePtr := NewFilePtr()
	fmt.Println(*FilePtr)
	fmt.Println(&FilePtr)

	fmt.Println("************************************\n")

	type person struct {}
	var tom *person = &person{}
	fmt.Printf("%p\n", tom)
	fmt.Printf("%p\n", *tom)
	fmt.Printf("%p\n", &tom)


}
