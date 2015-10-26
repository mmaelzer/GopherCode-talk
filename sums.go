package sums

func Sum(i, j int) int {
	return i + j
}

func SumVariadic(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
