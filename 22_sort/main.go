package main

import (
	"sort"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Persons []Person

func (p Persons) Len() int { return len(p) }
func (p Persons) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	sl := []string{"Dan", "Matt", "Bob", "Jenna", "Anders", "Anderson"}
	il := []int{25, 28, 28, 93, 44, 29}
	fmt.Println(sl)
	sort.Strings(sl)
	fmt.Println(sl)

	var perSlice Persons
	if (len(il) == len(sl)) {
		for i := range il {
			perSlice = append(perSlice, Person{sl[i], il[i]})
		}
	}

	fmt.Println(perSlice)
	sort.Sort(perSlice)
	fmt.Println(perSlice)

	slicePtr := new([]byte)
	fmt.Printf("Type of slicePtr: %T\n", slicePtr)
	*slicePtr = append(*slicePtr, []byte("strings!!")...)
	fmt.Println(*slicePtr)
	fmt.Printf("len: %v  cap: %v\n", len(*slicePtr), cap(*slicePtr))
}

