package main

import (
	"fmt"
)

// Produce the following output with nothing but the hash mark (#) and a newline character (\n).
// ########
//  ######
//   ####
//    ##
func main() {
	lineLength := 8

	for i := 0; i < lineLength/2; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (lineLength - (2 * i)); k++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
