package main

import "fmt"

func iterativeCount(target int, numbers []int) int {
	count := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			count++
		}
	}

	return count
}

func recursiveCount(target int, numbers []int) int {
	length := len(numbers)
	count := 0

	if numbers[length-1] == target {
		count = 1
	}

	if length == 1 {
		return count
	}

	return recursiveCount(target, numbers[:length-1]) + count
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	target1 := 0
	target2 := 10

	fmt.Println(iterativeCount(target1, numbers))
	fmt.Println(iterativeCount(target2, numbers))

	fmt.Println(recursiveCount(target1, numbers))
	fmt.Println(recursiveCount(target2, numbers))
}
