package main

import (
	"fmt"

	"github.com/NoHomey/algorithms-and-data-structures-2026/dsa/ringbuffer"
)

func removeElements(arr []int, deq *ringbuffer.RingBuffer[int], i int) {
	for {
		if first, ok := deq.First(); ok && arr[first] <= arr[i] {
			deq.Dequeue()
		} else {
			break
		}
	}
	for {
		if last, ok := deq.Last(); ok && arr[last] <= arr[i] {
			_, _ = deq.Pop()
		} else {
			break
		}
	}
}

func solve(arr []int, k int) []int {
	var result []int
	deq := ringbuffer.NewRingBuffer[int](k)

	for i := 0; i < k; i++ {
		removeElements(arr, deq, i)
		deq.Enqueue(i)
	}

	first, _ := deq.First()
	result = append(result, arr[first])

	for i := k; i < len(arr); i++ {
		if first, ok := deq.First(); ok && first == i-k {
			deq.Dequeue()
		}
		removeElements(arr, deq, i)
		deq.Enqueue(i)
		first, _ := deq.First()
		result = append(result, arr[first])
	}

	return result
}

func main() {
	result := solve([]int{2, -1, 3, 2, 5, 7, 6, 8, 11}, 4)
	fmt.Println(result)
}
