package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	var sumSlice []int

	for _, numbers := range numberToSum {
		sumSlice = append(sumSlice, Sum(numbers))
	}

	return sumSlice
}

