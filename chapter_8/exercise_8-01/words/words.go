package words

import (
	"bufio"
	"os"
	"strings"
)

const filepath = "/mnt/c/dev/goworkspace/src/github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/data/words.txt"

// GetWordlist returns a slice of string containing valid English words
func GetWordlist(wordLength int) []string {
	wordfile, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer wordfile.Close()

	wordlist := make([]string, 0)
	scanner := bufio.NewScanner(wordfile)

	for scanner.Scan() {
		word := scanner.Text()

		if len(word) == wordLength {
			wordlist = append(wordlist, word)
		}
	}

	return wordlist
}

// FilterWords returns a slice of string containing the words that
// include the letter (include = true) or do not contain the letter
// (include = false).
func FilterWords(words []string, letter rune, include bool) []string {
	tempWords := make([]string, 0)

	for _, word := range words {
		containsLetter := strings.ContainsRune(word, letter)

		if include && !containsLetter {
			continue
		} else if !include && containsLetter {
			continue
		}

		tempWords = append(tempWords, word)
	}

	return tempWords
}

// GetCommonestPattern returns a slice of string containing the words that
// exhibit the most common pattern of the given letter in the given list
// of words, as well as the pattern as a slice of rune.
func GetCommonestPattern(originalWords []string, letter rune) ([]string, []rune) {
	var commonestPattern []rune
	var patternCount int
	words := originalWords

	for len(words) > 0 {
		count := 0
		pattern := getPattern(words[0], letter)

		for _, word := range words[1:] {
			if matchPattern(word, pattern) {
				count++
			}
		}

		if count > patternCount {
			patternCount = count
			commonestPattern = pattern
		}

		words = filterPattern(words, pattern, false)
	}

	words = filterPattern(originalWords, commonestPattern, true)

	return words, commonestPattern
}

func getPattern(s string, letter rune) []rune {
	pattern := make([]rune, 0)

	for _, r := range s {
		var char rune
		if r == letter {
			char = letter
		} else {
			char = 0
		}
		pattern = append(pattern, char)
	}

	return pattern
}

func matchPattern(s string, pattern []rune) bool {
	for i, r := range pattern {
		if r != 0 {
			if rune(s[i]) != r {
				return false
			}
		}
	}

	return true
}

func filterPattern(words []string, pattern []rune, include bool) []string {
	filteredWords := make([]string, 0)

	for _, word := range words {
		if include {
			if matchPattern(word, pattern) {
				filteredWords = append(filteredWords, word)
			}
		} else {
			if !matchPattern(word, pattern) {
				filteredWords = append(filteredWords, word)
			}
		}

	}

	return filteredWords
}
