package iterative

func Filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, v := range numbers {
		if callback(v) {
			xs = append(xs, v)
		}
	}

	return xs
}
