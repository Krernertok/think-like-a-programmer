package words

import (
	"bufio"
	"os"
	"strings"
)

const filepath = "/mnt/c/dev/goworkspace/src/github.com/krernertok/think-like-a-programmer/chapter_8/hangman/data/words.txt"

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
func FilterWords(words []string, r rune, include bool) []string {
	tempWords := make([]string, 0)

	for _, word := range words {
		containsLetter := strings.ContainsRune(word, r)

		if include && containsLetter {
			tempWords = append(tempWords, word)
		} else if !include && !containsLetter {
			tempWords = append(tempWords, word)
		}
	}

	return tempWords
}

// GetUpdatedList returns a slice of strings with an updated
// word list based on the given letter and the letter pattern
// used to filter the list.
func GetUpdatedList(words []string, letter rune) ([]string, []rune) {
	wordsWithoutLetter := FilterWords(words, letter, false)
	wordsWithLetter := FilterWords(words, letter, true)
	wordsWithCommonestPattern, pattern := getCommonestPattern(wordsWithLetter, letter)

	if len(wordsWithoutLetter) >= len(wordsWithCommonestPattern) {
		return wordsWithoutLetter, []rune{}
	}

	return wordsWithCommonestPattern, pattern
}

// getCommonestPattern returns a slice of string containing the words that
// exhibit the most common pattern of the given letter in the given list
// of words, as well as the most common pattern as a slice of rune.
func getCommonestPattern(originalWords []string, letter rune) ([]string, []rune) {
	var commonestPattern []rune
	var patternCount int
	words := originalWords

	for len(words) > 0 {
		count := 1
		pattern := GetPattern(words[0], letter)

		for _, word := range words[1:] {
			if matchPattern(word, pattern, letter) {
				count++
			}
		}

		if count > patternCount {
			patternCount = count
			commonestPattern = pattern
		}

		words = filterPattern(words, pattern, letter, false)
	}

	words = filterPattern(originalWords, commonestPattern, letter, true)

	return words, commonestPattern
}

// GetPattern returns the pattern of where the given rune appears in the
// string.
func GetPattern(s string, letter rune) []rune {
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

func matchPattern(s string, pattern []rune, letter rune) bool {
	for i, r := range pattern {
		if r != 0 {
			if rune(s[i]) != r {
				return false
			}

		} else {
			if rune(s[i]) == letter {
				return false
			}
		}
	}

	return true
}

func filterPattern(words []string, pattern []rune, letter rune, include bool) []string {
	filteredWords := make([]string, 0)

	for _, word := range words {
		if include {
			if matchPattern(word, pattern, letter) {
				filteredWords = append(filteredWords, word)
			}
		} else {
			if !matchPattern(word, pattern, letter) {
				filteredWords = append(filteredWords, word)
			}
		}

	}

	return filteredWords
}

// SamePattern returns true if patterns are the same, false if not.
func SamePattern(p1, p2 []rune) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i, r := range p1 {
		if r != p2[i] {
			return false
		}
	}

	return true
}

// MergePatterns adds the runes in p1 into p2 according to their position
// in p1. p1 may not be longer than p2
func MergePatterns(p1, p2 []rune) []rune {
	for i, r := range p1 {
		if r != 0 {
			p2[i] = r
		}
	}
	return p2
}
