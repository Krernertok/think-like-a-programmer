package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getBinary(359)) // 101100111
	fmt.Println(getBinary(3))   // 11
	fmt.Println(getBinary(774)) // 1100000110

	fmt.Println(getDecimal("101100111"))  // 359
	fmt.Println(getDecimal("11"))         // 3
	fmt.Println(getDecimal("1100000110")) // 774
}

// getBinary returns the binary representation of an integer as a string
func getBinary(i int) string {
	binary := []string{}

	for i > 0 {
		var bit string
		if i%2 == 0 {
			bit = "0"
		} else {
			bit = "1"
		}
		binary = append([]string{bit}, binary...)
		i /= 2
	}

	return strings.Join(binary, "")
}

// getDecimal returns an integer from a binary number represented by a string
func getDecimal(b string) int {
	decimal := 0

	for i, v := range b {
		bit := v - '0'
		exp := len(b) - i - 1
		if bit == 1 {
			decimal += power(2, exp)
		}
	}

	return decimal
}

// power calculates an integer x to its nth power
func power(x, n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	result := x
	for i := 1; i < n; i++ {
		result *= x
	}
	return result
}
