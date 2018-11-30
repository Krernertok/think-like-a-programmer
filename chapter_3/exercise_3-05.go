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
	encoded := encode(input)
	decoded := decode(encoded)

	fmt.Println(input)
	fmt.Println(encoded)
	fmt.Println(decoded)
}

func encode(input string) string {
	str := strings.ToUpper(input)
	encodedChars := make([]rune, len(input))

	for i, char := range str {
		encodedChars[i] = getCharCypher(char)
	}
	return string(encodedChars)
}

func decode(input string) string {
	decodedChars := make([]rune, len(input))

	for i, char := range input {
		decodedChars[i] = getCharCypher(char)
	}

	return string(decodedChars)
}

func getCharCypher(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return cypher[char-'A']
	} else {
		return char
	}
}
