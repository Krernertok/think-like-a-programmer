package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

// n should be a valid pointer to a node struct
func isHeap(n *node) bool {
	if n.left != nil {
		if !isHeap(n.left) || n.left.value > n.value {
			return false
		}
	}

	if n.right != nil {
		if !isHeap(n.right) || n.right.value > n.value {
			return false
		}
	}

	return true
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
	node1 := &node{1, nil, nil}
	node4 := &node{4, nil, nil}
	node7 := &node{7, nil, nil}
	node6 := &node{6, node4, node7}
	node3 := &node{3, node1, node6}
	node13 := &node{13, nil, nil}
	node14 := &node{14, node13, nil}
	node10 := &node{10, nil, node14}
	bst := &node{8, node3, node10}

	// Binary tree (BT)
	//     5
	//    / \
	//   1   4
	bt := &node{5, node1, node4}

	fmt.Println("BST is a heap:", isHeap(bst))
	fmt.Println("BT is a heap:", isHeap(bt))
}
