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

func (l integerList) sum() int {
	return recursiveSum(l.firstNode)
}

func recursiveSum(n *integerNode) int {
	if n == nil {
		return 0
	}

	if n.next == nil {
		return n.value
	}

	return recursiveSum(n.next) + n.value
}

func main() {
	intList := &integerList{nil}

	for i := 1; i < 10; i++ {
		intList.append(i)
	}

	fmt.Println(intList.sum())
}
