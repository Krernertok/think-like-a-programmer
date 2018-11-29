package main

import "fmt"

func main() {
	unsorted := []int{3, 4, 1, 2, 5}
	sorted := []int{1, 2, 3, 4, 5}
	one := []int{0}
	empty := []int{}

	fmt.Println("Expected false:", isSorted(unsorted))
	fmt.Println("Expected true:", isSorted(sorted))
	fmt.Println("Expected true:", isSorted(one))
	fmt.Println("Expected true:", isSorted(empty))
}

func isSorted(ints []int) bool {
	itemCount := len(ints)

	if itemCount == 0 || itemCount == 1 {
		return true
	}

	for i := 1; i < itemCount; i++ {
		if ints[i] < ints[i-1] {
			return false
		}
	}

	return true
}
