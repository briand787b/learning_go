package main

import "fmt"

type path []byte

func (p *path) ToUpper() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func ShortenSlice(slicePtr *[]int) {
	// returns a shortened slice
	*slicePtr = (*slicePtr)[0:len(*slicePtr)-1]
}

func main() {
	pathName := path("/usr/bin/tso")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
	fmt.Print("herro worrd!\n")

	var arr [10]int
	for i := range arr {
		arr[i] = i
	}

	slice := arr[0:10]
	fmt.Println("before shorten function: ", slice)
	ShortenSlice(&slice)
	fmt.Println("after shorten function: " , slice)

	var bytes []byte = []byte("string")
	fmt.Printf("%s\n", bytes)
	fmt.Println("length: ", len(bytes))
	fmt.Println("Capacity", cap(bytes), "\n")

	// composite literal
	items := []int{ 0, 1, 2, 3, 4 }
	items = append(items, 5)
	fmt.Println("Length: ", len(items))
	fmt.Println("Capacity: ", cap(items))
	items = items[:8]
	fmt.Println("Length: ", len(items))
	fmt.Println("Capacity: ", cap(items))
	items = items[:10]
	fmt.Println("Length: ", len(items))
	fmt.Println("Capacity: ", cap(items))

	str1 := "string"
	// str1 = append(str1, 'n')
	fmt.Println(str1, "\n")

	var byteSlice []byte
	fmt.Println("byteslice length: ", len(byteSlice))
	fmt.Println("byteslice capacity: ", cap(byteSlice))
	byteSlice = append(byteSlice, 'A')
	fmt.Println("byteslice array ptr: ", &(byteSlice[0]))

	slash := "/usr/local/bin/go"[0]
	fmt.Printf("String representation: %q\n", slash)
	fmt.Println(string(slash), "\n")
}
