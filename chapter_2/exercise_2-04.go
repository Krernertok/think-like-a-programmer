package main

import (
	"fmt"
)

// Own shape
// Produce the following output with nothing but the hash mark (#), the space
// ( ), and the newline character (\n):
//
// #      #
//  #    #
//   #  #
//    ##
//   #  #
//  #    #
// #      #
func main() {
	width := 8
	height := 7

	for i := 0; i < height; i++ {
		var j int
		for j = 0; j < (3 - abs(3-i)); j++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
		for k := 0; k < (width - 2*j - 2); k++ {
			fmt.Print(" ")
		}
		fmt.Print("#\n")
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
