package main

import (
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/ui"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/words"
	"strings"
)

type gameState struct {
	correctAnswer                  string
	maxGuesses, wrongGuesses       int
	wordlist                       []string
	guessedLetters, correctLetters []rune
}

// TODO:
//	- pattern from []rune to []bool

func main() {
	wordLength, maxGuesses := ui.Init()
	wordlist := words.GetWordlist(wordLength)

	cheatingHangman(wordlist, maxGuesses, wordLength)
}

func cheatingHangman(wordlist []string, maxGuesses, wordLength int) {
	wrongGuesses := 0
	guessedLetters := make([]rune, 0)
	correctLetters := make([]rune, wordLength)

	for wrongGuesses < maxGuesses {
		ui.UpdateUI(wrongGuesses, maxGuesses, guessedLetters, correctLetters)

		nextLetter := ui.GetNextGuess(guessedLetters)
		guessedLetters = append(guessedLetters, nextLetter)

		var letterPattern []rune
		wordlist, letterPattern = words.GetUpdatedList(wordlist, nextLetter)

		if words.SamePattern(letterPattern, []rune{}) {
			wrongGuesses++
		} else {
			correctLetters = words.MergePatterns(correctLetters, letterPattern)

			if len(wordlist) == 1 {
				regularHangman(
					wordlist[0],
					wrongGuesses,
					maxGuesses,
					guessedLetters,
					correctLetters,
				)
			}
		}

	}
	ui.PlayerLoses()
}

func regularHangman(correctAnswer string, wrongGuesses, maxGuesses int, guessedLetters, correctLetters []rune) {
	for wrongGuesses < maxGuesses {
		if string(correctLetters) == correctAnswer {
			ui.PlayerWins(correctAnswer)
		}

		ui.UpdateUI(wrongGuesses, maxGuesses, guessedLetters, correctLetters)

		nextLetter := ui.GetNextGuess(guessedLetters)
		guessedLetters = append(guessedLetters, nextLetter)

		if strings.ContainsRune(correctAnswer, nextLetter) {
			pattern := words.GetPattern(correctAnswer, nextLetter)
			correctLetters = words.MergePatterns(pattern, correctLetters)
		} else {
			wrongGuesses++
		}
	}
	ui.PlayerLoses()
}
