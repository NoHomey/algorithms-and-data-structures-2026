package main

import "fmt"

type Array[T any] struct {
	data []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{data: nil}
}

func NewArrayIntWithCap[T any](capacity int) *Array[T] {
	a := Array[T]{data: make([]T, 0, capacity)}
	return &a
}

func (a *Array[T]) IsEmpty() bool {
	return a.Len() == 0 // len((*a).data) == 0
}

func (a *Array[T]) IsFull() bool {
	return a.Len() == a.Cap()
}

func (a *Array[T]) Len() int {
	return len(a.data)
}

func (a *Array[T]) Cap() int {
	return cap(a.data)
}

func (a *Array[T]) Push(x T) {
	if a.IsFull() {
		newCap := 2 * a.Cap()
		if newCap == 0 {
			newCap = 32
		}
		tmp := make([]T, a.Len(), newCap)
		/*
			for i, x := range a.data {
				tmp[i] = x
			}
		*/
		// full copy of a data
		copy(tmp, a.data)
		a.data = tmp
	}
	a.data = append(a.data, x)
}

func (a *Array[T]) Last() T {
	if a.IsEmpty() {
		panic("Calling Last on an empty Array")
	}
	return a.data[a.Len()-1]
}

func (a *Array[T]) First() T {
	if a.IsEmpty() {
		panic("Calling First on an empty Array")
	}
	return a.data[0]
}

func (a *Array[T]) Pop() (T, bool) {
	if a.IsEmpty() {
		var zero T
		return zero, false
	}
	last := a.Last()
	a.data = a.data[:a.Len()-1]
	return last, true
}

type Stack[T any] struct {
	array *Array[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{array: NewArray[T]()}
}

func (s Stack[T]) IsEmpty() bool {
	return s.array.IsEmpty()
}

func (s Stack[T]) IsFull() bool {
	return s.array.IsFull()
}

func (s Stack[T]) Push(x T) {
	s.array.Push(x)
}

func (s Stack[T]) Pop() (T, bool) {
	return s.array.Pop()
}

func (s Stack[T]) Len() int {
	return s.array.Len()
}

func (s Stack[T]) Cap() int {
	return s.array.Cap()
}

func main() {
	s := NewStack[int]()

	s.Push(2)
	s.Push(3)
	s.Push(4)
	x, ok := s.Pop()
	s.Push(5)

	// 4 true 3 32
	fmt.Println(x, ok, s.Len(), s.Cap())
}
