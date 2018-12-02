package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	inputLength := 50
	input := make([]int, inputLength)

	for i := 0; i < inputLength; i++ {
		input[i] = rand.Int() % 10
	}

	counts := make(map[int]int)
	for _, n := range input {
		counts[n]++
	}

	mode := []int{}
	modeCount := 1
	for k, v := range counts {
		if v == modeCount {
			mode = append(mode, k)
		} else if v > modeCount {
			mode = []int{k}
			modeCount = v
		}
	}

	fmt.Println("Input:", input)

	// if all items in the list only appeared once, there is no mode
	if modeCount == 1 {
		fmt.Println("There is no mode.")
		return
	}

	if len(mode) == 1 {
		m := mode[0]
		fmt.Printf("The mode is: %v. It occurred %d times.\n", m, modeCount)
	} else {
		fmt.Printf("The modes are: %v. They occurred %d times.\n", mode, modeCount)
	}
}
