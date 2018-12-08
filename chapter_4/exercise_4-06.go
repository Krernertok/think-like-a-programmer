package main

import "fmt"

type charNode struct {
	char rune
	next *charNode
}

type charList struct {
	headPointer *charNode
}

func (cl *charList) append(r rune) {
	node := cl.headPointer
	newChar := charNode{
		char: r,
		next: nil,
	}

	if node == nil {
		cl.headPointer = &newChar
		return
	}

	for node.next != nil {
		node = node.next
	}

	node.next = &newChar
}

func (cl charList) characterAt(index int) rune {
	i := 0
	node := cl.headPointer

	for index != i {
		node = node.next
		i++
	}

	return node.char
}

func (cl charList) output() string {
	chars := make([]rune, 0)

	for n := cl.headPointer; n != nil; n = n.next {
		chars = append(chars, n.char)
	}

	return string(chars)
}

func main() {

	cl := &charList{}
	cl.append('c')
	cl.append('a')
	cl.append('t')

	fmt.Println(cl.output())
	c := cl.characterAt(1)
	fmt.Println(string(c))
}
