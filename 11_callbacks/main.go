package main

import "fmt"

func main() {
	var button = 0;
	clickButton(button, clickButtonEventHandler)
}

func clickButton(sender int, callback func(int)) {
	sender += 1
	callback(sender)
}

func clickButtonEventHandler(buttonNum int) {
	fmt.Println("The button that was clicked was ", buttonNum)
}
