package wordlist

import (
	"bufio"
	"os"
)

const filepath = "/mnt/c/dev/goworkspace/src/github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/data/words.txt"

// GetWordlist returns a slice of string containing valid English words
func GetWordlist() []string {
	// - Open wordlist file
	wordfile, err := os.Open(filepath)
	check(err)
	defer wordfile.Close()

	wordlist := make([]string, 0)
	scanner := bufio.NewScanner(wordfile)

	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	return wordlist
}

// helper for checking errors as seen on www.gobyexample.com
func check(e error) {
	if e != nil {
		panic(e)
	}
}
