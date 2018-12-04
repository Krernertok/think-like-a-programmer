package main

import (
	"fmt"
	"sort"
)

var salesPerPerson = [][]int{
	// []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 387587},
	// []int{5865, 5456, 3983, 6464, 9957, 4785, 3875, 3838, 4959, 1122, 7766, 2534},
	[]int{1856, 498, 30924, 87478, 328, 2653, 387, 3754, 387587, 2873, 276, 32},
	[]int{5865, 5456, 3983, 6464, 9957, 4785, 3875, 3838, 4959, 1122, -1, -1},
	[]int{23, 55, 67, 99, 265, 376, 232, 223, 4546, 564, 4544, 3434},
}

func main() {
	highestMedian := 0
	for _, sales := range salesPerPerson {
		actualSales := getActualSales(sales)
		median := getMedian(actualSales)

		if median > highestMedian {
			highestMedian = median
		}
	}
	fmt.Println("Highest median sales:", highestMedian)
}

func getActualSales(sales []int) []int {
	actualSales := []int{}
	for _, month := range sales {
		if month == -1 {
			continue
		}
		actualSales = append(actualSales, month)
	}
	return actualSales
}

func getMedian(sales []int) int {
	var median int
	length := len(sales)
	midpoint := length / 2

	sort.Sort(monthlySales(sales))
	fmt.Println(sales)

	if length%2 == 0 {
		median = (sales[midpoint] + sales[midpoint-1]) / 2
	} else {
		median = sales[midpoint]
	}

	return median
}

type monthlySales []int

func (m monthlySales) Len() int {
	return len(m)
}

func (m monthlySales) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m monthlySales) Less(i, j int) bool {
	return m[i] < m[j]
}
