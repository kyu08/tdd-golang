package array

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int
	for _, n := range numbersToSum {
		result = append(result, Sum(n))
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int
	for _, n := range numbersToSum {
		if len(n) == 0 {
			result = append(result, 0)
			continue
		}
		result = append(result, Sum(n[1:]))
	}
	return result
}
