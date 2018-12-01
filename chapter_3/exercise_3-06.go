package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var cypher [26]rune

func main() {
	rand.Seed(time.Now().Unix())

	input := "This text will be encoded."

	randomizeCypher()

	encoded := encode(input)
	decoded := decode(encoded)

	fmt.Println(input)
	fmt.Println(encoded)
	fmt.Println(decoded)
}

func randomizeCypher() {
	for i := 0; i < 26; i++ {
		var index int
		for {
			r := rand.Int()
			index = r % 26
			if index != i && cypher[index] == 0 {
				break
			}
		}
		cypher[index] = rune(int('A') + i)
	}
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
