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

func (l *integerList) reverse() {
	l.firstNode = reverseNode(l.firstNode, nil)
}

// Not really easier to solve recursively
func reverseNode(node, prev *integerNode) *integerNode {
	temp := node.next
	node.next = prev
	if temp == nil {
		return node
	}
	return reverseNode(temp, node)
}

func (l integerList) printNodes() {
	node := l.firstNode

	fmt.Println("Printing nodes:")
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func main() {
	intList := &integerList{nil}

	for i := 1; i < 10; i++ {
		intList.append(i)
	}

	intList.printNodes()

	intList.reverse()
	intList.printNodes()
}
