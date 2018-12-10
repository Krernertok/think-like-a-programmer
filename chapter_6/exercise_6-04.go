package main

import "fmt"

func recursiveReverse(nums []int) []int {
	length := len(nums)

	if length == 0 || length == 1 {
		return nums
	}
	return append([]int{nums[length-1]}, recursiveReverse(nums[:length-1])...)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
	fmt.Println(recursiveReverse(numbers))

}
