/*Since slices are dynamical, I will implement the string and linked
list structures instead.*/
package main

import "fmt"

type arrayString []rune

func (a *arrayString) append(r rune) {
	*a = append(*a, r)
}

func (a arrayString) runeAt(index int) rune {
	return a[index]
}

func (a *arrayString) concatenate(b arrayString) {
	*a = append(*a, b...)
}

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

func main() {
	a := arrayString{'h', 'e', 'l', 'l', 'o'}
	b := arrayString{'w', 'o', 'r', 'l', 'd'}

	fmt.Println(string(a), string(b))

	a.append(',')
	a.append(' ')
	fmt.Println(string(a))

	a.concatenate(b)
	fmt.Println(string(a))

	fmt.Println(string(a.runeAt(0)))

	fmt.Println("---")

	node := listNode{1, 80, nil}
	sc := studentCollection{&node}
	sc.addRecord(2, 90)
	sc.addRecord(3, 100)
	sc.addRecord(4, 100)

	fmt.Println("Average:", sc.averageRecord())

	empty := studentCollection{nil}
	fmt.Println("Average of empty:", empty.averageRecord())
}
