package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type charCount struct {
	char  string
	count int
}

type byCount []charCount

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input:")

	scanner.Scan()
	input := scanner.Text()

	fmt.Println("---")

	words := strings.Split(input, " ")
	longestWord, maxChars := longestWordStats(words)
	charCounts := charCountStats(words)
	totalChars := 0

	for _, charCount := range charCounts {
		totalChars += charCount.count
	}

	fmt.Println("Number of words in input:", len(words))
	fmt.Println("Longest word:", longestWord)
	fmt.Println("Total characters in longest word:", maxChars)
	fmt.Println("Total characters in input (excluding spaces):", totalChars)
	fmt.Println("---")

	for i, charCount := range charCounts {
		fmt.Printf("%d. Char %s appeared %d times.\n", i, charCount.char, charCount.count)
	}

	fmt.Println("---")
}

func longestWordStats(words []string) (string, int) {
	var longestWord string
	var maxChars int

	for _, w := range words {
		wordLength := len(w)
		if wordLength > maxChars {
			longestWord = w
			maxChars = wordLength
		}
	}

	return longestWord, maxChars
}

func charCountStats(words []string) []charCount {
	var counts []charCount
	hist := make(map[string]int)

	for _, word := range words {
		for _, char := range word {
			hist[string(char)]++
		}
	}

	for k, v := range hist {
		counts = append(counts, charCount{k, v})
	}

	sort.Sort(byCount(counts))

	return counts
}

func (b byCount) Len() int {
	return len(b)
}

func (b byCount) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byCount) Less(i, j int) bool {
	return b[i].count > b[j].count
}
