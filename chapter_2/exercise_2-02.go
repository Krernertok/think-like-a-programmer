package main

import (
	"fmt"
)

// Produce the following output with nothing but the hash mark (#), the space
// ( ), and the newline character (\n):
//
//    ##
//   ####
//  ######
// ########
// ########
//  ######
//   ####
//    ##
func main() {
	lineLength := 8
	halfLine := lineLength / 2

	for i := 1; i <= halfLine; i++ {
		for j := 0; j < halfLine-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2 * i); k++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}

	for i := 0; i < halfLine; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (lineLength - (2 * i)); k++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
