package pracpac

import "fmt"

func PrintMultiples(printable string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print(printable)
	}
}
