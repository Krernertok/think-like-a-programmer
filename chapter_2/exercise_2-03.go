package main

import "fmt"

// Produce the following output with nothing but the hash mark (#), the space
// ( ), and the newline character (\n):
//
// #            #
//  ##        ##
//   ###    ###
//    ########
//    ########
//   ###    ###
//  ##        ##
// #            #

func main() {
	shapeWidth := 14
	shapeHeight := 8

	// upper half of the shape
	for i := 1; i <= shapeHeight/2; i++ {
		var j, k int
		for j = 0; j < i-1; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < i; k++ {
			fmt.Print("#")
		}
		for l := 0; l < (shapeWidth - 2*j - 2*k); l++ {
			fmt.Print(" ")
		}
		for m := 0; m < i; m++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}

	// lower half of the shape
	for i := shapeHeight / 2; i > 0; i-- {
		var j, k int
		for j = 0; j < i-1; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < i; k++ {
			fmt.Print("#")
		}
		for l := 0; l < (shapeWidth - 2*j - 2*k); l++ {
			fmt.Print(" ")
		}
		for m := 0; m < i; m++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
