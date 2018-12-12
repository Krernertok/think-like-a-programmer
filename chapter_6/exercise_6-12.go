package main

import (
	"fmt"
	"sort"
)

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

func median(n *intNode) int {
	values := intList(n)
	sort.Ints(values)
	midpoint := len(values) / 2

	if len(values)%2 == 0 {
		return (values[midpoint] + values[midpoint-1]) / 2
	} else {
		return values[midpoint]
	}
}

func intList(n *intNode) []int {
	values := []int{n.value}

	if n.left != nil {
		values = append(values, intList(n.left)...)
	}

	if n.right != nil {
		values = append(values, intList(n.right)...)
	}

	return values
}

// n should not be nil
func mode(n *intNode) int {
	counts := recursiveCount(n)
	value, count := 0, 0

	for k, v := range counts {
		if v > count {
			value = k
			count = v
		}
	}

	return value
}

func recursiveCount(n *intNode) map[int]int {
	counts := make(map[int]int)
	counts[n.value]++

	if n.left != nil {
		for k, v := range recursiveCount(n.left) {
			counts[k] += v
		}
	}

	if n.right != nil {
		for k, v := range recursiveCount(n.right) {
			counts[k] += v
		}
	}

	return counts
}

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

	fmt.Println("Average:", average(bst)) // 7

	node6.value = 1
	node13.value = 2
	node14.value = 2
	node10.value = 2
	fmt.Println("Mode:", mode(bst)) // 2
	node6.value = 6
	node13.value = 13
	node14.value = 14
	node10.value = 10

	fmt.Println("Median:", median(bst)) // 7
	// node14.left = nil
	// fmt.Println("Median:", median(bst)) // 6

}
