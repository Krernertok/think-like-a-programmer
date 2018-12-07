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

func (a *arrayString) replaceString(target, replaceText arrayString) {
	source := *a
	sourceLength := len(source)
	replaceIndexes := make([]int, 0)
	targetLength := len(target)

	// special case: empty arrayString as both source and target
	if sourceLength == 0 && targetLength == 0 {
		*a = replaceText
		return
	}

Outer:
	for i := 0; i < sourceLength; i++ {
		for j := 0; j < targetLength; j++ {
			outOfBounds := i+j >= sourceLength
			mismatch := source[i+j] != target[j]

			if outOfBounds || mismatch {
				continue Outer
			}
		}
		replaceIndexes = append([]int{i}, replaceIndexes...)
	}

	for _, i := range replaceIndexes {
		end := append(arrayString{}, source[i+targetLength:]...)

		source = append(source[:i], replaceText...)
		source = append(source, end...)
	}

	*a = source
}

func main() {
	a := arrayString{'h', 'e', 'l', 'l', 'o'}
	target := arrayString{'l'}
	replace := arrayString{'s', 'z'}

	a.replaceString(target, replace)
	fmt.Println(string(a))
	fmt.Println("---")

	a = arrayString{'h', 'e', 'l', 'l', 'o'}
	target = arrayString{'l'}
	replace = arrayString{}

	a.replaceString(target, replace)
	fmt.Println(string(a))
	fmt.Println("---")

	a = arrayString{}
	target = arrayString{}
	replace = arrayString{'h', 'e', 'l', 'l', 'o'}

	a.replaceString(target, replace)
	fmt.Println(string(a))
}
