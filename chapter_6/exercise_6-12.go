package main

import "fmt"

type intNode struct {
	value int
	left  *intNode
	right *intNode
}

func average(n *intNode) int {
	sum, numTerms := btAverage(n)
	return sum / numTerms
}

func btAverage(n *intNode) (int, int) {
	sum := n.value
	numTerms := 1

	if n.left != nil {
		tempSum, tempNumTerms := btAverage(n.left)
		sum += tempSum
		numTerms += tempNumTerms
	}
	if n.right != nil {
		tempSum, tempNumTerms := btAverage(n.right)
		sum += tempSum
		numTerms += tempNumTerms
	}

	return sum, numTerms
}

// func btMedian(n *intNode) int {}

// func btMode(n *intNode) int {}

func main() {
	// Binary search tree (BST)
	//               8
	//             /   \
	//           3      10
	//          / \       \
	//         1   6       14
	//	          / \     /
	//           4   7  13
	node1 := &intNode{1, nil, nil}
	node4 := &intNode{4, nil, nil}
	node7 := &intNode{7, nil, nil}
	node6 := &intNode{6, node4, node7}
	node3 := &intNode{3, node1, node6}
	node13 := &intNode{13, nil, nil}
	node14 := &intNode{14, node13, nil}
	node10 := &intNode{10, nil, node14}
	bst := &intNode{8, node3, node10}

	fmt.Println("Average:", average(bst))
}
