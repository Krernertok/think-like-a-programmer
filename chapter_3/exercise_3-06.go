package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

	rand.Seed(time.Now().Unix())

	input := "This text will be encoded."

	rand.Shuffle(len(cypher), func(i, j int) {
		cypher[i], cypher[j] = cypher[j], cypher[i]
	})

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
		encodedChars[i] = encodeChar(char)
	}
	return string(encodedChars)
}

func decode(input string) string {
	decodedChars := make([]rune, len(input))

	for i, char := range input {
		decodedChars[i] = decodeChar(char)
	}

	return string(decodedChars)
}

func encodeChar(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return cypher[char-'A']
	} else {
		return char
	}
}

func decodeChar(char rune) rune {
	if char < 'A' || char > 'Z' {
		return char
	}

	var c rune
	for i, r := range cypher {
		if char == r {
			rInt := int('A') + i
			c = rune(rInt)
			break
		}
	}

	return c
}
