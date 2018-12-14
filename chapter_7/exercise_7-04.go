package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

type studentRecord struct {
	id    int
	name  string
	grade int
}

type studentCollection struct {
	buckets []*list.List
}

func newStudentCollection() *studentCollection {
	numBuckets := 1000
	buckets := make([]*list.List, numBuckets)

	for i := 0; i < numBuckets; i++ {
		buckets[i] = list.New()
	}

	return &studentCollection{buckets}
}

func (sc *studentCollection) addRecord(s studentRecord) {
	bucketIndex := s.id % len(sc.buckets)
	bucket := sc.buckets[bucketIndex]
	elem := bucket.Front()

	for elem != nil && elem.Value.(studentRecord).id != s.id {
		elem = elem.Next()
	}

	if elem != nil && elem.Value.(studentRecord).id == s.id {
		return
	}

	sc.buckets[bucketIndex].PushBack(s)
}

func (sc studentCollection) getRecord(id int) (studentRecord, bool) {
	bucketIndex := id % len(sc.buckets)
	elem := sc.buckets[bucketIndex].Front()

	for elem != nil {
		student := elem.Value.(studentRecord)
		if student.id == id {
			return student, true
		}
		elem = elem.Next()
	}

	return studentRecord{}, false
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
