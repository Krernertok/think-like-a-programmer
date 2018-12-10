package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

// n should be a valid pointer to a node struct
func isBST(n *node) bool {
	if n.left != nil {
		if !isBST(n.left) || n.left.value > n.value {
			return false
		}
	}

	if n.right != nil {
		if !isBST(n.right) || n.right.value < n.value {
			return false
		}
	}
	return true
}

// duplicates are dropped
func bstInsert(value int, bst *node) {
	if value == bst.value {
		return
	}

	var parent *node
	var lessThan bool

	if value < bst.value {
		parent = bst.left
		lessThan = true
	} else {
		parent = bst.right
		lessThan = false
	}

	if parent != nil {
		bstInsert(value, parent)
		return
	}

	n := &node{
		value: value,
		left:  nil,
		right: nil,
	}

	if lessThan {
		bst.left = n
	} else {
		bst.right = n
	}
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

	values := []int{0, 2, 9, 17}
	for _, v := range values {
		fmt.Println("Adding value", v, "into binary search tree.")
		bstInsert(v, bst)
		fmt.Println("Still a binary search tree:", isBST(bst))
		fmt.Println("---")
	}
	fmt.Println("Left child of 1:", node1.left.value)
	fmt.Println("Right child of 1:", node1.right.value)
	fmt.Println("Left child of 10:", node10.left.value)
	fmt.Println("Right child of 14:", node14.right.value)
}
