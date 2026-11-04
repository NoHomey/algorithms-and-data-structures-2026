package main

import (
	"fmt"

	"github.com/NoHomey/algorithms-and-data-structures-2026/dsa/stack"
)

func solve(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	s := stack.NewStack[int]()
	for i := 2*n - 1; i >= 0; i-- {
		j := i % n
		current := arr[j]
		for {
			if t, ok := s.Top(); ok && t < current {
				s.Pop()
			} else {
				break
			}

		}
		if t, ok := s.Top(); ok && t > current {
			res[j] = t
		} else {
			res[j] = -1
		}
		s.Push(current)
	}
	return res
}

func main() {
	result := solve([]int{2, 3, 1, 5, 6, 4})
	fmt.Println(result)
}
