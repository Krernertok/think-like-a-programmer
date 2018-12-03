package main

import (
	"fmt"
	"sort"
)

type student struct {
	grade int
	id    int
	name  string
}

func main() {
	students := []student{
		{87, 10001, "Fred"},
		{28, 10002, "Tom"},
		{100, 10003, "Alistair"},
		{78, 10004, "Sasha"},
		{84, 10005, "Erin"},
		{98, 10006, "Belinda"},
		{75, 10007, "Leslie"},
		{70, 10008, "Candy"},
		{81, 10009, "Aretha"},
		{68, 10010, "Veronica"},
	}

	sort.Sort(byGrade(students))
	numStudents := len(students)

	q1 := students[(numStudents / 4)].grade
	q2 := students[(numStudents / 2)].grade
	q3 := students[((numStudents * 3) / 4)].grade

	fmt.Println("To be as good or better than 25% of students, score at least:", q1)
	fmt.Println("To be as good or better than 50% of students, score at least:", q2)
	fmt.Println("To be as good or better than 75% of students, score at least:", q3)
}

type byGrade []student

func (g byGrade) Len() int {
	return len(g)
}

func (g byGrade) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g byGrade) Less(i, j int) bool {
	return g[i].grade < g[j].grade
}
