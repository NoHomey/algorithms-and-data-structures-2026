package main

type ArrayInt struct {
	data []int
}

func NewArrayInt() *ArrayInt {
	return &ArrayInt{data: nil}
}

func NewArrayIntWithCap(capacity int) *ArrayInt {
	a := ArrayInt{data: make([]int, 0, capacity)}
	// Go automatically moves the pointer to the Heap.
	// This is not like C where we would return a pointer to memory on the stack which will be an error or a bug!
	return &a
}

func (a *ArrayInt) IsEmpty() bool {
	return a.Len() == 0 // len((*a).data) == 0
}

func (a *ArrayInt) Len() int {
	return len(a.data)
}

func (a *ArrayInt) Cap() int {
	return cap(a.data)
}

func (a *ArrayInt) Push(x int) {
	// Exercise implement `Push` without using the built-in `append` function.
	// For this use the built-in fuction `make` to allocate larger slice and a for loop to copy the elements.
	// Next time we will use `make` and `copy` and make a Generic Array and start writing and using our own packages.
	a.data = append(a.data, x)
}
