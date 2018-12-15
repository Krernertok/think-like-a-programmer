package wordlist

import (
	"bufio"
	"os"
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

		if len(word) <= wordLength {
			wordlist = append(wordlist, word)
		}
	}

	return wordlist
}
