package main

import (
	"fmt"
	"strings"
)

var cypher = []rune{
	'Z',
	'Y',
	'X',
	'W',
	'V',
	'U',
	'T',
	'S',
	'R',
	'Q',
	'P',
	'O',
	'N',
	'M',
	'L',
	'K',
	'J',
	'I',
	'H',
	'G',
	'F',
	'E',
	'D',
	'C',
	'B',
	'A',
}

func main() {
	input := "This text will be encoded."
	str := strings.ToUpper(input)
	encodedChars := make([]rune, len(input))

	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			encodedChars[i] = cypher[char-'A']
		} else {
			encodedChars[i] = char
		}
	}

	fmt.Println(input)
	fmt.Println(string(encodedChars))
}
