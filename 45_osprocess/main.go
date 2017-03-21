package main

import (
	"os/exec"
	"fmt"
	"log"
	"time"
	"bufio"
)

func main() {
	cmd := exec.Command("firefox", "http://google.com")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	pid := cmd.Process.Pid
	fmt.Println(pid)
	time.Sleep(1*time.Second)
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("There was an error")
		log.Println(err)
	}

	for {
		exec.Command("free", "-h")
		time.Sleep(1*time.Second)
		fmt.Println("\r")
	}
	bufio.ScanLines()
}
