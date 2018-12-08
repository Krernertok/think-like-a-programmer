package main

import "fmt"

type intNode struct {
	value int
	next  *intNode
}

type intList struct {
	firstNode *intNode
}

func intToList(num int) *intList {
	var node *intNode
	var iList intList

	for num > 0 {
		mod := num % 10
		n := intNode{
			value: mod,
			next:  nil,
		}

		if node == nil {
			node = &n
		} else {
			n.next = node
			node = &n
		}

		num /= 10
	}

	iList = intList{node}

	return &iList
}

func main() {
	integer := intToList(123)
	node := integer.firstNode

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}
