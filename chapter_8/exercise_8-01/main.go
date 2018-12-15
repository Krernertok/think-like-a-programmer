package main

import (
	"fmt"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/ui"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/wordlist"
)

// Parts:
// - main function, keeps track of game state
// - game logic
// - UI

func main() {
	wordLength, numGuesses := ui.Init()
	wordlist := wordlist.GetWordlist(wordLength)

	for i := 0; i < 10; i++ {
		fmt.Println(wordlist[i])
	}
}
