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

type byGrade []student
type byId []student

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
	printStudents(students)

	fmt.Println("---")

	sort.Sort(byId(students))
	printStudents(students)

}

func printStudents(students []student) {
	for i, s := range students {
		fmt.Printf("%d: %s (ID: %d) - %d\n", i+1, s.name, s.id, s.grade)
	}
}

func (g byGrade) Len() int {
	return len(g)
}

func (g byGrade) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g byGrade) Less(i, j int) bool {
	return g[i].grade > g[j].grade
}

func (b byId) Len() int {
	return len(b)
}

func (b byId) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byId) Less(i, j int) bool {
	return b[i].id < b[j].id
}
