package main

import (
	"fmt"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/wordlist"
)

// Parts:
// - main function, keeps track of game state
// - code for retrieving the word list
// - game logic
// - UI

func main() {
	wordlist := wordlist.GetWordlist()

	for i := 0; i < 10; i++ {
		fmt.Println(wordlist[i])
	}
}
