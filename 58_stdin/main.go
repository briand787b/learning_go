package main

import (
	"fmt"
	"os"
	"os/exec"
)

func so() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "2").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		fmt.Println("I got the byte", b, "("+string(b)+")")
	}
}

func brian() {
	var b []byte = make([]byte, 100)
	count := 0

	for {
		os.Stdin.Read(b)
		// fmt.Println(b)
		count++
		if count > 10 {
			fmt.Println(b)
		}
	}

}

func main() {
	brian()
	// so()
}