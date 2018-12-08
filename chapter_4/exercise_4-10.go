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

func intListSum(a, b intList) *intList {
	var answer *intList
	var currNode *intNode
	var carryOver int

	reverseList(&a)
	reverseList(&b)
	currA, currB := a.firstNode, b.firstNode

	for currA != nil || currB != nil {
		valA := getNodeValue(currA)
		valB := getNodeValue(currB)

		tempSum := valA + valB + carryOver
		carryOver = tempSum / 10

		node := intNode{
			value: tempSum % 10,
			next:  nil,
		}

		if answer != nil {
			currNode.next = &node
			currNode = &node
		} else {
			answer = &intList{&node}
			currNode = &node
		}

		if currA != nil {
			currA = currA.next
		}

		if currB != nil {
			currB = currB.next
		}
	}

	if carryOver > 0 {
		node := intNode{
			value: carryOver,
			next:  nil,
		}
		currNode.next = &node
	}

	reverseList(answer)
	return answer
}

func reverseList(l *intList) {
	var prev, curr, next *intNode
	curr = l.firstNode

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.firstNode = prev
}

func getNodeValue(n *intNode) int {
	if n != nil {
		return n.value
	}
	return 0
}

func main() {
	integer := intToList(9222)
	integer2 := intToList(11777)

	sum := intListSum(*integer, *integer2)
	node := sum.firstNode

	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}
