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

const shapeWidth = 14
const shapeHeight = 8

func main() {

	// upper half of the shape
	for i := 1; i <= shapeHeight/2; i++ {
		printPatternLine(i)
	}

	// lower half of the shape
	for i := shapeHeight / 2; i > 0; i-- {
		printPatternLine(i)
	}
}

func printPatternLine(i int) {
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
