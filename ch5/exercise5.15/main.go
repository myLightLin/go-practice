package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	input := []int{0, 1, 2, 3, 4}
	first := input[2]
	left := input[2:]

	m, err := max(input...)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("max: %d\n", m)

	m = mustMax(first, left...)
	fmt.Printf("mustMax: %d\n", m)
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("must have at least one argument!")
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func mustMax(first int, vals ...int) int {
	max := first
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
