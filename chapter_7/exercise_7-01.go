package main

import "fmt"

// constants for choosing first student policy
const (
	GRADE = "grade"
	ID    = "id"
	NAME  = "name"
)

type studentRecord struct {
	name  string
	id    int
	grade int
}

type studentNode struct {
	studentData studentRecord
	next        *studentNode
}

type studentCollection struct {
	headPointer        *studentNode
	firstStudentPolicy func(r1, r2 studentRecord) bool
}

var firstStudentPolicies = map[string]func(r1, r2 studentRecord) bool{
	GRADE: func(r1, r2 studentRecord) bool {
		return r2.grade > r1.grade
	},
	ID: func(r1, r2 studentRecord) bool {
		return r2.id < r1.id
	},
	NAME: func(r1, r2 studentRecord) bool {
		return r2.name < r1.name
	},
}

func (sc *studentCollection) addRecord(newStudent studentRecord) {
	newNode := studentNode{
		studentData: newStudent,
		next:        sc.headPointer,
	}
	sc.headPointer = &newNode
}

func (sc *studentCollection) getFirstStudent() *studentNode {
	if sc.firstStudentPolicy == nil {
		return nil
	}

	node := sc.headPointer
	firstStudent := node

	for node != nil && node.next != nil {
		if sc.firstStudentPolicy(node.studentData, node.next.studentData) {
			firstStudent = node.next
		}
		node = node.next
	}

	return firstStudent
}

func (sc *studentCollection) recordWithID(id int) *studentNode {
	node := sc.headPointer

	for node.studentData.id != id {
		node = node.next

		if node == nil {
			return nil
		}
	}

	return node
}

func (sc *studentCollection) removeRecord(id int) {
	var prev *studentNode
	node := sc.headPointer

	for node.studentData.id != id {
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

func (sc *studentCollection) setFirstStudentPolicy(policy string) {
	sc.firstStudentPolicy = firstStudentPolicies[policy]
}

func main() {
	records := []studentRecord{
		studentRecord{"John", 1, 73},
		studentRecord{"Albert", 2, 75},
		studentRecord{"Jill", 3, 83},
		studentRecord{"Bob", 4, 86},
		studentRecord{"Alice", 5, 90},
		studentRecord{"Eve", 6, 97},
		studentRecord{"Phil", 7, 98},
		studentRecord{"Donna", 8, 100},
	}
	sc := studentCollection{}

	for _, r := range records {
		sc.addRecord(r)
	}

	fmt.Println("First student if policy hasn't been set:", sc.getFirstStudent())

	for _, policy := range []string{GRADE, ID, NAME} {
		sc.setFirstStudentPolicy(policy)
		s := sc.getFirstStudent()
		fmt.Printf(
			"First student by %s: %s (ID: %d, grade: %d)\n",
			policy,
			s.studentData.name,
			s.studentData.id,
			s.studentData.grade,
		)
	}
}
