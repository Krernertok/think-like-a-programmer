package main

import (
	"fmt"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/ui"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/words"
)

// Parts:
// - main function, keeps track of game state
// - filtering -> wordlist
// - UI

func main() {
	wrongGuesses := 0
	wordLength, maxGuesses := ui.Init()
	wordlist := words.GetWordlist(wordLength)

	// guessedLetters := make([]rune, 0)

	for wrongGuesses < maxGuesses {
		// update UI

		// get next guess from ui
		// filter wordlist based on next guess
		// if wordlist > 1
		//	   update ui with guessed letters
		//     loop
		// else proceed to "normal" hangman loop
		wrongGuesses++
	}

	filteredWords := words.FilterWords(wordlist, 'l', true)
	for i := 0; i < len(filteredWords); i++ {
		fmt.Println(filteredWords[i])
	}

}
