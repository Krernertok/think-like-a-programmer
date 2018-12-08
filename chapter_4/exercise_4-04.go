package main

import "fmt"

// everything about this feels wrong

type arrayString []rune

func (a *arrayString) append(r rune) {
	*a = append((*a), r)
	(*a)[0]++
}

func (a arrayString) runeAt(index int) rune {
	return a[index+1]
}

func (a *arrayString) concatenate(b arrayString) {
	*a = append(*a, b[1:]...)
	// update the length
	(*a)[0] = rune(int((*a)[0]) + len(b) - 1)
}

func (a arrayString) output() string {
	return string(a[1:])
}

func main() {
	a := arrayString{rune(5), 'h', 'e', 'l', 'l', 'o'}
	(&a).append('!')
	printOutput(a)

	(&a).concatenate(a)
	printOutput(a)
}

func printOutput(a arrayString) {
	fmt.Println(a[0])
	fmt.Println((&a).output())
	fmt.Println(string((&a).runeAt(0)))
}
