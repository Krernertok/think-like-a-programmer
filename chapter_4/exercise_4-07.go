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

func main() {

	cl := &charList{}
	cl.append('c')
	cl.append('a')
	cl.append('t')

	cl2 := &charList{}
	cl.append('n')
	cl.append('a')
	cl.append('p')

	cl.concatenate(*cl2)

	fmt.Println(cl.output())

}
