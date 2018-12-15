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

type byId []student

func printStudents(students []student) {
	for i, s := range students {
		fmt.Printf("%d: %s (ID: %d) - %d\n", i+1, s.name, s.id, s.grade)
	}
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

// Implement using binary search
func searchByID(s []student, id int) (student, bool) {
	lo := 0
	hi := len(s) - 1

	for lo != hi {
		mid := (hi + lo) / 2

		if s[mid].id > id {
			hi = mid - 1
		} else if s[mid].id < id {
			lo = mid + 1
		} else {
			return s[mid], true
		}
	}

	if s[lo].id == id {
		return s[lo], true
	}

	return student{}, false
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

	sort.Sort(byId(students))
	printStudents(students)

	ids := []int{
		10001,
		10002,
		10003,
		10004,
		10005,
		10006,
		10007,
		10008,
		10009,
		10010,
		1000000,
	}

	for _, id := range ids {
		fmt.Println("---")
		fmt.Println("Searching for ID:", id)
		student, ok := searchByID(students, id)

		if ok {
			fmt.Printf("Found student named %s with ID %d\n", student.name, student.id)
		} else {
			fmt.Println("Could not find student with ID:", id)
		}
	}

}
