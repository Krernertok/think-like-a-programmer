/*Since slices are dynamical, I will implement the string and linked
list structures instead.*/
package main

import "fmt"

type listNode struct {
	studentID int
	grade     int
	next      *listNode
}

type studentCollection struct {
	headPointer *listNode
}

func (sc *studentCollection) addRecord(id, grade int) {
	newNode := listNode{
		studentID: id,
		grade:     grade,
		next:      sc.headPointer,
	}
	sc.headPointer = &newNode
}

func (sc studentCollection) averageRecord() int {
	node := sc.headPointer

	if node == nil {
		return 0
	}

	sum := 0
	nodeCount := 0

	for node != nil {
		sum += node.grade
		nodeCount++
		node = node.next
	}

	return sum / nodeCount
}

func (sc *studentCollection) removeRecord(id int) {
	var prev *listNode
	node := sc.headPointer

	for node.studentID != id {
		prev = node
		node = node.next

		if node == nil {
			return
		}
	}

	if prev == nil {
		sc.headPointer = node.next
	} else {
		prev.next = node.next
	}
}

func (sc studentCollection) printCollection() {
	node := sc.headPointer

	fmt.Println("Printing student collection")
	for node != nil {
		fmt.Println("ID:", node.studentID, "\tGrade:", node.grade)
		node = node.next
	}
}

func main() {
	node := listNode{1, 80, nil}
	sc := studentCollection{&node}
	sc.addRecord(2, 90)
	sc.addRecord(3, 100)
	sc.addRecord(4, 100)
	sc.printCollection()

	sc.removeRecord(3)
	sc.printCollection()
	sc.removeRecord(4)
	sc.printCollection()

	sc.removeRecord(6)
}
