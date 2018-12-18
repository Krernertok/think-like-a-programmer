package player

import (
	"github.com/krernertok/think-like-a-programmer/chapter_8/exercise_8-01/ui"
	"time"
)

// HumanPlayer represents a human player
type HumanPlayer struct{}

// CPUPlayer represents an automated player
type CPUPlayer struct{}

// Player is an interface for the different player types
type Player interface {
	NextGuess([]rune) rune
}

func (HumanPlayer) NextGuess(guessed []rune) rune {
	return ui.GetNextGuess(guessed)
}

// alphabet in order of relative frequency
var guessOrder = []rune{
	'e',
	'a',
	'r',
	'i',
	'o',
	't',
	'n',
	's',
	'l',
	'c',
	'u',
	'd',
	'p',
	'm',
	'h',
	'g',
	'b',
	'f',
	'y',
	'w',
	'k',
	'v',
	'x',
	'z',
	'j',
	'q',
}

func (CPUPlayer) NextGuess(guesses []rune) rune {
	time.Sleep(time.Second)
	index := len(guesses)
	return guessOrder[index]
}
