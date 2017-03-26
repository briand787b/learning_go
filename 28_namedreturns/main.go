package main

import (
	"fmt"
)

type Thread struct {
	name string
	replies int
}

func pracReturns() (intSlice  []int) {
	intSlice = append(intSlice, []int{1, 2, 3}...)
	return
}

func (t *Thread) ChangeName(newName string) {
	t.name = newName
}

func AddReply(t *Thread) {
	t.replies++
}

func main() {
	intSl := pracReturns()
	fmt.Println(intSl)
	newThread := Thread{"favorite food", 0}
	fmt.Println(newThread)
	(&newThread).ChangeName("Least favorite food")
	fmt.Println(newThread.name)
	newThread.name = "Favorite food"
	fmt.Println(newThread.name)
	threadPtr := &newThread
	AddReply(threadPtr)
	fmt.Println(newThread.replies)
}