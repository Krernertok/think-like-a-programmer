package main

import (
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/player"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/ui"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/words"
	"os"
	"strings"
)

func main() {
	ui.Welcome()
	initHangman()
}

func initHangman() {
	var playerType player.Player
	wordLength, maxGuesses, playerChoice := ui.Init()
	wordlist := words.GetWordlist(wordLength)

	if playerChoice == "human" {
		playerType = player.HumanPlayer{}
	} else {
		playerType = player.CPUPlayer{}
	}

	cheatingHangman(wordlist, maxGuesses, wordLength, playerType)
}

func cheatingHangman(wordlist []string, maxGuesses, wordLength int, p player.Player) {
	wrongGuesses := 0
	guessedLetters := make([]rune, 0)
	correctLetters := make([]rune, wordLength)

	for wrongGuesses < maxGuesses {
		ui.UpdateUI(wrongGuesses, maxGuesses, guessedLetters, correctLetters)

		nextLetter := p.NextGuess(guessedLetters)
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
					p,
				)
			}
		}

	}

	gameEnd(false, "")
}

func regularHangman(correctAnswer string, wrongGuesses, maxGuesses int, guessedLetters, correctLetters []rune, p player.Player) {
	for wrongGuesses < maxGuesses {
		if string(correctLetters) == correctAnswer {
			gameEnd(true, correctAnswer)
		}

		ui.UpdateUI(wrongGuesses, maxGuesses, guessedLetters, correctLetters)

		nextLetter := p.NextGuess(guessedLetters)
		guessedLetters = append(guessedLetters, nextLetter)

		if strings.ContainsRune(correctAnswer, nextLetter) {
			pattern := words.GetPattern(correctAnswer, nextLetter)
			correctLetters = words.MergePatterns(pattern, correctLetters)
		} else {
			wrongGuesses++
		}
	}

	gameEnd(false, "")
}

func gameEnd(playerWon bool, answer string) {
	if ui.EndScreen(playerWon, answer) {
		initHangman()
	}
	os.Exit(0)
}
