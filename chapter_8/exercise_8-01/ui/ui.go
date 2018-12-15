package ui

import "fmt"

// Init returns the number of guesses and letters to be used in the game
func Init() (letters, guesses int) {
	fmt.Println("Welcome to Hangman!")

	// TODO: Check that input is valid
	fmt.Println("How many letters should the word contain?")
	fmt.Scanln(&letters)

	fmt.Println("How many wrong guesses are allowed?")
	fmt.Scanln(&guesses)

	return letters, guesses
}
