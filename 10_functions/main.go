package main

import "fmt"

func main() {
	fmt.Println(greet("brian", "dostilio"))
	fmt.Println(greetPeople("Brian", "James", "Doug", "Matt"))

	funcExp := pracFuncExp()
	funcExp()
}

func pracFuncExp() func()  {
	return func()  {
		fmt.Println("this works")
	}
}

func greet(fname, lname string) string {
	return fmt.Sprint("hello ", fname, lname)
}

func greetPeople(names ...string) string {
	var allNames string
	for _, v := range names {
		allNames += " " + v + " "
	}
	
	return allNames
}
