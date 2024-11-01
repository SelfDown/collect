package collect

func RangeNumber(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i
	}
	return numbers
}
