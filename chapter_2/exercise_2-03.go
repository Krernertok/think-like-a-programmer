package main

import "fmt"

// Produce a custom output with nothing but the hash mark (#), a space ( ) and a newline character (\n):
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
		for m := 1; m <= i; m++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
