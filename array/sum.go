package array

func Sum(numbers []int) int {

	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}



func SumAll(numbersToSum ... []int) []int {
	var sums []int

	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
