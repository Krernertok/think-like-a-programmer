package main

import (
	"fmt"
	"sort"
)

var sales = [][]int{
	[]int{1856, 498, 30924, 87478, 328, 2653, 387, 3754, 387587, 2873, 276, 32},
	[]int{5865, 5456, 3983, 6464, 9957, 4785, 3875, 3838, 4959, 1122, 7766, 2534},
	[]int{23, 55, 67, 99, 265, 376, 232, 223, 4546, 564, 4544, 3434},
}

func main() {
	highestMedian := 0
	for _, salesPerPerson := range sales {
		sort.Sort(monthlySales(salesPerPerson))

		// 12 months in total, so median is the average of the middle two months
		median := (salesPerPerson[5] + salesPerPerson[6]) / 2
		if median > highestMedian {
			highestMedian = median
		}
	}
	fmt.Println("Highest median sales:", highestMedian)
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
