// Package arrays provides functions to work with slices of integers.
package arrays

// Sum returns the sum of all elements in a slice of integers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAll returns a slice of integers containing the sum of each slice of integers passed as arguments.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns a slice of integers containing the sum of all elements except the first one in each slice of integers passed as arguments.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(x[1:]))
	}
	return Reduce(numbers, sumTail, []int{})
}
