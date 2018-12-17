package main

import (
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/ui"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/words"
	"strings"
)

func main() {
	wrongGuesses := 0
	wordLength, maxGuesses := ui.Init()
	wordlist := words.GetWordlist(wordLength)
	guessedLetters := make([]rune, 0)
	correctGuesses := make([]rune, wordLength)

	for i := 0; i < wordLength; i++ {
		correctGuesses[i] = 0
	}

	for true {
		if wrongGuesses == maxGuesses {
			ui.PlayerLoses()
		}

		var pattern []rune
		ui.UpdateUI(wrongGuesses, maxGuesses, guessedLetters, correctGuesses)

		nextLetter := ui.GetNextGuess()
		guessedLetters = append(guessedLetters, nextLetter)

		wordlist, pattern = words.GetUpdatedList(wordlist, nextLetter)
		emptyPattern := []rune{}
		numWords := len(wordlist)

		if numWords == 1 {
			break
		}

		if words.SamePattern(pattern, emptyPattern) {
			wrongGuesses++
			continue
		}

		correctGuesses = words.MergePatterns(pattern, correctGuesses)
	}

	correctAnswer := wordlist[0]

	for wrongGuesses < maxGuesses {
		if string(correctGuesses) == correctAnswer {
			ui.PlayerWins(correctAnswer)
		}

		ui.UpdateUI(wrongGuesses, maxGuesses, guessedLetters, correctGuesses)

		nextLetter := ui.GetNextGuess()
		guessedLetters = append(guessedLetters, nextLetter)

		if strings.ContainsRune(correctAnswer, nextLetter) {
			pattern := words.GetPattern(correctAnswer, nextLetter)
			correctGuesses = words.MergePatterns(pattern, correctGuesses)
			continue
		}

		wrongGuesses++
	}

	ui.PlayerLoses()
}
