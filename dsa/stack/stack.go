package stack

import array "github.com/NoHomey/algorithms-and-data-structures-2026/dsa/array"

type Stack[T any] struct {
	array *array.Array[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{array: array.NewArray[T]()}
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

func (s Stack[T]) Top() (T, bool) {
	return s.array.Last()
}

func (s Stack[T]) Len() int {
	return s.array.Len()
}

func (s Stack[T]) Cap() int {
	return s.array.Cap()
}
