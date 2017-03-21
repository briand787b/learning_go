package main

import (
	"fmt"
	"os"
	"syscall"
	"os/exec"
)

type path []byte

type Person struct {
	first string
	age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func pracPointer(loc string) *path {
	p := make(path, len(loc), 2 * len(loc))
	copy(p, path(loc))
	return &p
}

func pracPersonPointer(name string, age int) *Person {
	p := new(Person)
	p.first = name
	p.age = age
	return p
}

func pracPerson(p *Person) string {
	return p.first
}

func personPrac(p Person) (age int, name string) {
	age = p.age
	name = p.first
	return
}

func main() {
	fmt.Println("querying hostname")
	hostName, _ := os.Hostname()
	fmt.Println(hostName)

	fmt.Println("getting current working directory")
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	fmt.Println("changing direcotry")
	err := os.Chdir("~")
	if os.IsNotExist(err) {
		fmt.Println("directory not found")
		fmt.Println("error: ", err)
	}
	fmt.Println("getting current working directory")
	pwd, _ = os.Getwd()
	fmt.Println(pwd)

	fmt.Println("The temp directory: ")
	temp := os.TempDir()
	fmt.Println(temp)

	// create a new file in the /tmp/junk folder
	err = os.Chdir("/tmp/junk")
	if err != nil {
		fmt.Println("Error in changing directories: ", err)
	}
	file, errTouch := os.Create("fromGoGo")
	if errTouch != nil {
		fmt.Println("error in creating file: ", errTouch)
	}
	fmt.Println("file location: ", &file)
	fmt.Println(file.Name())
	file.Write([]byte("this is the content"))
	fileInfo, statErr := os.Stat(file.Name())
	if statErr != nil {
		fmt.Println(statErr)
	}
	fmt.Printf("fileInfo: %+v", fileInfo)
	err = file.Chown(1000, 1000)
	if err != nil {
		fmt.Println("Mode was not changed")
	}
	fileInfo, _ = os.Stat(file.Name())
	fmt.Println("file mode: ", fileInfo.Mode())
	fmt.Println("file's owner is: ", fileInfo.Sys().(*syscall.Stat_t).Uid)

	// use the os/exec package
	path, err := exec.LookPath("/home/brian/dev/ruby/fifty_times.rb")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("path: ", path)
	var cmd *exec.Cmd = exec.Command(path)
	fmt.Println("pointer to command: ", cmd)

	// *******************************
	// Practice pointers since I obviously suck at them
	// *******************************
	fileLoc := pracPointer("stuff")
	fmt.Println(fileLoc)

	somebodyPtr := pracPersonPointer("brian", 27)
	fmt.Printf("%T\n", *somebodyPtr)
	fmt.Println("field on somebodyPtr: ", (*somebodyPtr).first)
	fmt.Println("field on somebodyPtr: ", somebodyPtr.first)

	var somebody Person
	somebody.first = "nairb"
	somebody.age = 62
	fmt.Printf("%T\n", somebody)
	fmt.Println(somebody.first)

	fmt.Println(pracPerson(&somebody))
	fmt.Println(personPrac(*somebodyPtr))

	// *******************************
	// Practice Structs
	// *******************************
	p1 := DoubleZero{
		Person: Person{
			first: "James",
			age: 20,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			first: "brian",
			age: 26,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)

}