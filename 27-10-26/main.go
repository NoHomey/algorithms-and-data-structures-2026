package main

import (
	"fmt"
	//"exec/dsa/queue"
	ringbuffer "github.com/NoHomey/algorithms-and-data-structures-2026/dsa/ringbuffer"
)

type Range struct {
	Beg int
	End int
}

func solve(arr []int) []Range {
	q := ringbuffer.NewRingBuffer[*Range](16)
	for i, x := range arr {
		if q.IsEmpty() {
			q.Enqueue(&Range{Beg: i, End: i})
			continue
		}
		last, ok := q.Last()
		if !ok {
			panic("Something unexpected happened")
		}
		if arr[last.End]+1 == x {
			// (*last).End = i
			last.End = i
		} else {
			q.Enqueue(&Range{Beg: i, End: i})
		}
	}
	data := q.Data()
	res := make([]Range, len(data))
	for i := range data {
		res[i] = *data[i]
	}
	return res
}

func main() {
	arr := []int{2, 1, 2, 4, 5, 6, 2, 3, 5}
	res := solve(arr)
	fmt.Println(arr)
	fmt.Println(res)
}
