package main

import "fmt"

func iterativeOddParity(b []int) bool {
	count := 0

	for i := 0; i < len(b); i++ {
		if b[i] == 1 {
			count++
		}
	}

	return count%2 == 1
}

func recursiveOddParity(b []int) bool {
	length := len(b)

	if length == 1 {
		return b[0] == 1
	}

	parityOfRest := recursiveOddParity(b[:length-1])

	// pOR() is true and b[l-1] is 1: false
	// pOR() is true and b[l-1] is 0: true
	// pOR() is false and b[l-1] is 1: true
	// pOR() is false and b[l-1] is 0: false
	// => XOR (no XOR operator in Go?)

	if parityOfRest {
		return b[length-1] != 1
	} else {
		return b[length-1] == 1
	}
}

func main() {
	binary1 := []int{0, 0, 0, 1, 0, 1, 0, 1}
	binary2 := []int{0, 0, 0, 0, 0, 1, 0, 1}

	fmt.Println(iterativeOddParity(binary1))
	fmt.Println(iterativeOddParity(binary2))

	fmt.Println(recursiveOddParity(binary1))
	fmt.Println(recursiveOddParity(binary2))
}
