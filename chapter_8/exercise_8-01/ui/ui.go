package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Init returns the number of guesses and letters to be used in the game
func Init() (letters, guesses int) {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	validLetter := false
	validGuess := false

	fmt.Printf("\nWelcome to Hangman!\n\n")

	fmt.Println("How many letters should the word contain?")

	for !validLetter {
		scanner.Scan()
		letters, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Please enter a number!")
			continue
		}

		validLetter = true
		err = nil
	}

	fmt.Println("How many wrong guesses are allowed?")

	for !validGuess {
		scanner.Scan()
		guesses, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Please enter a number!")
			continue
		}

		validGuess = true
	}

	return letters, guesses
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
func GetNextGuess() rune {
	var char rune
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Guess the next letter:")

	for true {
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		length := len(input)

		if length != 1 {
			fmt.Println("Please enter a single letter!")
			continue
		}

		char = rune(input[0])

		if char < 'a' || char > 'z' {
			fmt.Println("Please enter a letter!")
			continue
		}

		break
	}

	return char
}

// PlayerWins prints the win screen
func PlayerWins(answer string) {
	fmt.Println("Congratulations! You've won! The correct answer is:", answer)
	os.Exit(0)
}

// PlayerLoses prints the game over screen.
func PlayerLoses() {
	fmt.Println("Game over! You ran out of guesses!")
	os.Exit(0)
}
