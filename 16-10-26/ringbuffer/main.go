package main

import "fmt"

type RingBuffer[T any] struct {
	data  []T
	first int
	last  int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity < 1 {
		panic(fmt.Sprintf("expecting capacity to be positive integer, got: %d", capacity))
	}
	return &RingBuffer[T]{
		data:  make([]T, capacity),
		first: 0,
		last:  0,
	}
}

func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.first == rb.last && !rb.IsFull()
}

func (rb *RingBuffer[T]) IsFull() bool {
	return (rb.first+rb.last+1)%cap(rb.data) == 0
}

func (rb *RingBuffer[T]) Enqueue(x T) {
	if rb.IsFull() {
		newCapacity := 2 * cap(rb.data)
		tmp := make([]T, newCapacity)
		copy(tmp, rb.data[rb.first:cap(rb.data)])
		if rb.first != 0 {
			copy(tmp[rb.first+1:], rb.data[0:rb.first+1])
		}
		rb.last = cap(rb.data)
		rb.data = tmp
		rb.first = 0
	}
	rb.data[rb.last] = x
	rb.last = (rb.last + 1) % cap(rb.data)
}

func (rb *RingBuffer[T]) Dequeue() (T, bool) {
	if rb.IsEmpty() {
		var tmp T
		return tmp, false
	}
	x := rb.data[rb.first]
	rb.first = (rb.first + 1) % cap(rb.data)
	return x, true
}

func main() {
	q := NewRingBuffer[int](5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	// [1, 2, 3, 4, 0]
	fmt.Println(q.data)
	x, ok := q.Dequeue()
	// 1, true, [1, 2, 3, 4, 0]
	fmt.Println(x, ok, q.data)
	x, ok = q.Dequeue()
	// 2, true, [1, 2, 3, 4, 0]
	fmt.Println(x, ok, q.data)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	// [6, 7, 3, 4, 5]
	fmt.Println(q.data)
	q.Enqueue(8)
	fmt.Println(q.data)
}
