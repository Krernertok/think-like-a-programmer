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

func (sc studentCollection) recordsWithinRange(low, high int) *studentCollection {
	records := &studentCollection{}
	node := sc.headPointer

	for node != nil {
		if node.grade >= low && node.grade <= high {
			records.addRecord(node.studentID, node.grade)
		}
		node = node.next
	}

	return records
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
	sc := studentCollection{}
	sc.addRecord(1, 73)
	sc.addRecord(2, 75)
	sc.addRecord(3, 83)
	sc.addRecord(5, 86)
	sc.addRecord(6, 90)
	sc.addRecord(7, 97)
	sc.addRecord(8, 100)
	sc.printCollection()

	records := sc.recordsWithinRange(75, 90)
	records.printCollection()
}
