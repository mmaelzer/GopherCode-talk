package sums // HL

func SumTwo(i, j int) int {
	return sum(i, j)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func SumAny(nums ...int) int {
	return sum(nums...)
}
