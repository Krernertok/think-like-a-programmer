package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input:")

	scanner.Scan()
	input := scanner.Text()

	fmt.Println("---")

	words := strings.Split(input, " ")
	fmt.Println("Number of words in input:", len(words))

	var longestWord string
	var maxChars int

	for _, w := range words {
		wordLength := len(w)
		if wordLength > maxChars {
			longestWord = w
			maxChars = wordLength
		}
	}

	fmt.Println("Longest word:", longestWord, "Total characters:", maxChars)
	fmt.Println("---")
}
