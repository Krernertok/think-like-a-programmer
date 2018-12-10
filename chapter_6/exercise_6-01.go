package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(iterativeSum(numbers))
	fmt.Println(recursiveSum(numbers))
}

func iterativeSum(numbers []int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func recursiveSum(numbers []int) int {
	length := len(numbers)

	if length == 1 {
		return numbers[0]
	}

	return recursiveSum(numbers[:length-1]) + numbers[length-1]
}
