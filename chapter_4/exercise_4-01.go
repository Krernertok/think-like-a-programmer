/*Since slices are dynamical, I will implement the string and linked
list structures instead.*/
package main

import "fmt"

type arrayString []rune

func (a *arrayString) append(r rune) {
	*a = append(*a, r)
}

func (a arrayString) runeAt(index int) rune {
	return a[index]
}

func (a *arrayString) concatenate(b arrayString) {
	*a = append(*a, b...)
}

func main() {
	a := arrayString{'h', 'e', 'l', 'l', 'o'}
	b := arrayString{'w', 'o', 'r', 'l', 'd'}

	fmt.Println(string(a), string(b))

	a.append(',')
	a.append(' ')
	fmt.Println(string(a))

	a.concatenate(b)
	fmt.Println(string(a))

	fmt.Println(string(a.runeAt(0)))
}
