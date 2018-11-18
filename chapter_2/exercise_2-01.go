package main

import (
	"fmt"
	"strings"
)

// Produce the following output with nothing but the hash mark (#) and a newline character (\n).
// ########
//  ######
//   ####
//    ##
func main() {
	lineLength := 8
	var output []string

	for i := 0; i < lineLength/2; i++ {
		for j := 0; j < i; j++ {
			output = append(output, " ")
		}
		for k := 0; k < (lineLength - (2 * i)); k++ {
			output = append(output, "#")
		}
		output = append(output, "\n")
	}

	fmt.Print(strings.Join(output, ""))
}
