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

func (cl *charList) concatenate(cl2 charList) {
	// Note! Infinite loop if concatenating cl with itself
	node := cl2.headPointer
	for node != nil {
		cl.append(node.char)
		node = node.next
	}
}

func (cl charList) output() string {
	chars := make([]rune, 0)

	for n := cl.headPointer; n != nil; n = n.next {
		chars = append(chars, n.char)
	}

	return string(chars)
}

func (cl *charList) removeChars(start, numChars int) {
	i := 0
	node := cl.headPointer
	var prev *charNode

	for i != start && node != nil && node.next != nil {
		prev = node
		node = node.next
		i++
	}

	if node == nil {
		return
	}

	for i < start+numChars {
		if node != nil {
			node = node.next
		}
		i++
	}

	prev.next = node
}

func main() {
	cl := &charList{}
	cl.append('k')
	cl.append('e')
	cl.append('t')
	cl.append('c')
	cl.append('h')
	cl.append('u')
	cl.append('p')

	cl.removeChars(1, 2)

	fmt.Println(cl.output())
}
