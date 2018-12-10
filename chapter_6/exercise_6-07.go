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

func (l integerList) getCount(target int) int {
	return recursiveCount(l.firstNode, target)
}

func recursiveCount(n *integerNode, target int) int {
	if n == nil {
		return 0
	}

	count := 0

	if n.value == target {
		count = 1
	}

	if n.next == nil {
		return count
	}

	return recursiveCount(n.next, target) + count
}

func main() {
	numList1 := &integerList{nil}
	numList2 := &integerList{nil}
	numList3 := &integerList{nil}
	target := 9

	for i := 1; i < 10; i++ {
		numList1.append(i)
	}

	for i := 1; i < 10; i++ {
		numList2.append(i % 3)
	}

	fmt.Println("Occurences of", target, "in:")
	numList1.print()
	fmt.Println(numList1.getCount(target))

	fmt.Println("Occurences of", target, "in:")
	numList2.print()
	fmt.Println(numList2.getCount(target))

	fmt.Println("Occurences of", target, "in:")
	numList3.print()
	fmt.Println(numList3.getCount(target))
}
