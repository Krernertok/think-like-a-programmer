package main

import (
	"fmt"
	"math/rand"
)

type studentRecord struct {
	id    int
	name  string
	grade int
}

type studentCollection struct {
	students map[int]studentRecord
}

func newStudentCollection() *studentCollection {
	sc := studentCollection{
		students: make(map[int]studentRecord),
	}
	return &sc
}

func (sc *studentCollection) addRecord(s studentRecord) {
	sc.students[s.id] = s
}

func (sc studentCollection) getRecord(id int) (studentRecord, bool) {
	s, ok := sc.students[id]
	return s, ok
}

func main() {
	sc := newStudentCollection()

	for i := 0; i < 10000; i++ {
		student := studentRecord{
			id:    i,
			grade: rand.Int() % 100,
		}
		sc.addRecord(student)
	}

	for j := 0; j < 20; j++ {
		id := rand.Int() % 10000

		if id%2 == 0 {
			id += 10000
		}

		student, ok := sc.getRecord(id)

		fmt.Println("Looking for ID:", id)

		if ok {
			fmt.Println("Got record with ID:", student.id)
		} else {
			fmt.Println("Could not find record with ID:", id)
		}
	}
}
