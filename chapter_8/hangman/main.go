package main

import (
	"github.com/krernertok/think-like-a-programmer/chapter_8/hangman/player"
	"github.com/krernertok/think-like-a-programmer/chapter_8/hangman/ui"
	"github.com/krernertok/think-like-a-programmer/chapter_8/hangman/words"
	"os"
	"strings"
)

func main() {
	ui.Welcome()
	initHangman()
}

func initHangman() {
	var playerType player.Player
	wordLength, numGuesses, playerChoice := ui.Init()
	wordlist := words.GetWordlist(wordLength)

	if playerChoice == "human" {
		playerType = player.HumanPlayer{}
	} else {
		playerType = player.CPUPlayer{}
	}

	cheatingHangman(wordlist, numGuesses, wordLength, playerType)
}

func cheatingHangman(wordlist []string, numGuesses, wordLength int, p player.Player) {
	guessedLetters := make([]rune, 0)
	correctLetters := make([]rune, wordLength)

	for numGuesses > 0 {
		ui.UpdateUI(numGuesses, guessedLetters, correctLetters)

		nextLetter := p.NextGuess(guessedLetters)
		guessedLetters = append(guessedLetters, nextLetter)

		var letterPattern []rune
		wordlist, letterPattern = words.GetUpdatedList(wordlist, nextLetter)

		if words.SamePattern(letterPattern, []rune{}) {
			numGuesses--
		} else {
			correctLetters = words.MergePatterns(correctLetters, letterPattern)

			if len(wordlist) == 1 {
				regularHangman(
					wordlist[0],
					numGuesses,
					guessedLetters,
					correctLetters,
					p,
				)
			}
		}

	}

	gameEnd(false, "")
}

func regularHangman(correctAnswer string, numGuesses int, guessedLetters, correctLetters []rune, p player.Player) {
	for numGuesses > 0 {
		if string(correctLetters) == correctAnswer {
			gameEnd(true, correctAnswer)
		}

		ui.UpdateUI(numGuesses, guessedLetters, correctLetters)

		nextLetter := p.NextGuess(guessedLetters)
		guessedLetters = append(guessedLetters, nextLetter)

		if strings.ContainsRune(correctAnswer, nextLetter) {
			pattern := words.GetPattern(correctAnswer, nextLetter)
			correctLetters = words.MergePatterns(pattern, correctLetters)
		} else {
			numGuesses--
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
