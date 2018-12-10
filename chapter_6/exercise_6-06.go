package main

import "fmt"

type integerNode struct {
	value int
	next  *integerNode
}

type integerList struct {
	firstNode *integerNode
}

func (l *integerList) append(value int) {
	newNode := integerNode{
		value: value,
		next:  l.firstNode,
	}
	l.firstNode = &newNode
}

func (l integerList) print() {
	node := l.firstNode

	for node != nil {
		fmt.Print(node.value)
		node = node.next
	}

	fmt.Print("\n")
}

func (l integerList) oddParity() bool {
	return determineOddParity(l.firstNode)
}

func determineOddParity(n *integerNode) bool {
	if n == nil {
		return false
	}

	oddParity := n.value%2 != 0

	if n.next == nil {
		return oddParity
	}

	if determineOddParity(n.next) {
		return !oddParity
	} else {
		return oddParity
	}
}

func main() {
	binList1 := &integerList{nil}
	binList2 := &integerList{nil}

	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			binList1.append(1)
		} else {
			binList1.append(0)
		}

	}

	for i := 1; i < 10; i++ {
		if i%4 == 0 {
			binList2.append(1)
		} else {
			binList2.append(0)
		}

	}

	binList1.print()
	binList2.print()

	fmt.Println(binList1.oddParity())
	fmt.Println(binList2.oddParity())
}
