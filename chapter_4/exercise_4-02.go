package main

import "fmt"

type arrayString []rune

func (a *arrayString) append(r rune) {
	*a = append(*a, r)
}

func (a *arrayString) runeAt(index int) rune {
	return (*a)[index]
}

func (a *arrayString) concatenate(b arrayString) {
	*a = append(*a, b...)
}

func (a *arrayString) substring(start, length int) *arrayString {
	b := make(arrayString, length)
	copy(b, (*a)[start:start+length])
	return &b
}

func main() {
	a := arrayString{'h', 'e', 'l', 'l', 'o'}
	b := a.substring(1, 3)
	fmt.Println(string(a))
	fmt.Println(string(*b))
}
