package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	totalLength := len(numberToSum)
	sumArray := make([]int, totalLength)

	for i, numbers := range numberToSum {
		sumArray[i] = Sum(numbers)
	}

	return sumArray
}

