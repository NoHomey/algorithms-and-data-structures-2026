package array

type Array[T any] struct {
	data []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{data: nil}
}

func NewArrayFromSlice[T any](data []T) *Array[T] {
	return &Array[T]{data}
}

func NewArrayWithCap[T any](capacity int) *Array[T] {
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

func (a *Array[T]) Get(index uint) (T, bool) {
	if index < uint(a.Len()) {
		return a.data[index], true
	}
	var t T
	return t, false
}

func (a *Array[T]) Set(index uint, val T) {
	if index >= uint(a.Len()) {
		var t T
		for i := uint(0); i < index-uint(a.Len()); i++ {
			a.Push(t)
		}
	}
	a.data[index] = val
}
