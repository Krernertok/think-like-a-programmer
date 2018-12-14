package main

import (
	"container/list"
	"fmt"
)

type student struct {
	studentID int
	grade     int
}

type studentCollection struct {
	students *list.List
}

func (sc *studentCollection) addRecord(id, grade int) {
	newStudent := student{
		studentID: id,
		grade:     grade,
	}
	sc.students.PushFront(newStudent)
}

func (sc studentCollection) averageRecord() int {
	elem := sc.students.Front()

	if elem == nil {
		return 0
	}

	sum := 0
	studentCount := 0

	for elem != nil {
		sum += elem.Value.(student).grade
		studentCount++
		elem = elem.Next()
	}

	return sum / studentCount
}

func (sc *studentCollection) removeRecord(id int) {
	elem := sc.students.Front()

	for elem != nil && elem.Value.(student).studentID != id {
		elem = elem.Next()
	}

	if elem != nil {
		sc.students.Remove(elem)
	}
}

func (sc studentCollection) printCollection() {
	elem := sc.students.Front()

	fmt.Println("Printing student collection")
	for elem != nil {
		student := elem.Value.(student)
		fmt.Println("ID:", student.studentID, "\tGrade:", student.grade)
		elem = elem.Next()
	}
}

func main() {
	sc := studentCollection{
		students: list.New(),
	}
	sc.addRecord(2, 90)
	sc.addRecord(3, 95)
	sc.addRecord(4, 100)
	sc.printCollection()

	fmt.Println("Removing ID: 3")
	sc.removeRecord(3)
	sc.printCollection()

	fmt.Println("Removing ID: 4")
	sc.removeRecord(4)
	sc.printCollection()

	fmt.Println("Removing ID: 6")
	sc.removeRecord(6)
}
