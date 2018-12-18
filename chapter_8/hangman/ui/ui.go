package ui

import (
	"bufio"
	"fmt"
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/config"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

// Welcome prints the welcome message
func Welcome() {
	fmt.Printf("\nWelcome to Hangman!\n\n")
}

// Init returns the number of guesses and letters to be used in the game
func Init() (letters, guesses int, player string) {
	letters = getWordLength()
	guesses = getNumber("How many wrong guesses are allowed? ")
	player = getPlayerSelection()

	return letters, guesses, player
}

func getNumber(prompt string) int {
	var number int
	var err error

	for number == 0 {
		fmt.Printf(prompt)

		scanner.Scan()
		number, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Please enter a number!")
			continue
		}
	}

	return number
}

func getWordLength() int {
	var wordLength int
	prompt := "How many letters should the word contain? "

	for true {
		wordLength = getNumber(prompt)

		if wordLength < config.MAXWORDLENGTH {
			break
		}

		fmt.Println("Maximum word length is 22 characters!")
	}

	return wordLength
}

func getPlayerSelection() string {
	var playerSelection int
	prompt := "Do you want to player yourself (1) or do you want the computer (2) to play? "

	for true {
		playerSelection = getNumber(prompt)

		if playerSelection == 1 || playerSelection == 2 {
			break
		}

		fmt.Println("Please enter 1 for playing yourself or 2 for having the computer play!")
	}

	if playerSelection == 1 {
		return "human"
	}

	return "cpu"
}

// UpdateUI updates the user interface.
func UpdateUI(wrongGuesses, maxGuesses int, guessedLetters, pattern []rune) {
	var word []rune

	for _, r := range pattern {
		var char rune
		if r == 0 {
			char = '*'
		} else {
			char = r
		}
		word = append(word, char)
	}

	fmt.Println("---")
	fmt.Println("Word:\t", string(word))
	fmt.Println("Already guessed letters:", string(guessedLetters))
	fmt.Println("Wrong guesses left:", (maxGuesses - wrongGuesses))
}

// GetNextGuess returns the next character the player guesses.
func GetNextGuess(guessed []rune) rune {
	var r rune

	for true {
		fmt.Printf("Guess the next letter: ")

		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		if len(input) != 1 {
			fmt.Println("Please enter a single letter!")
			continue
		}

		r = rune(input[0])

		if r < 'a' || r > 'z' {
			fmt.Println("Please enter a letter!")
			continue
		}

		if runeInSlice(r, guessed) {
			fmt.Println("You've already guessed that letter! Try another one.")
			continue
		}

		break
	}

	return r
}

// EndScreen prints the final result and prompts the user to play again
func EndScreen(playerWon bool, answer string) bool {
	if playerWon {
		fmt.Println("Congratulations, you've won! The correct answer was:", answer)
	} else {
		fmt.Println("Game over! You ran out of guesses!")
	}

	fmt.Printf("Play again (y/n): ")
	var playAgain bool

	for true {
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		if len(input) != 1 {
			fmt.Println("Please enter y or n!")
			continue
		}

		r := rune(input[0])

		if r == 'y' {
			playAgain = true
			break
		} else if r == 'n' {
			playAgain = false
			break
		} else {
			fmt.Println("Please enter y or n!")
			continue
		}
	}

	return playAgain
}

func runeInSlice(r rune, sr []rune) bool {
	for _, rune := range sr {
		if rune == r {
			return true
		}
	}

	return false
}
