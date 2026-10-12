As an excerice implement [Queue](https://en.wikipedia.org/wiki/Queue_(abstract_data_type)) data structure backed by a [Ring Buffer](https://en.wikipedia.org/wiki/Circular_buffer) as follows:

```go
type RingBuffer[T any] struct {
	// you decide what data members you need
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T]

func (rb *RingBuffer[T]) IsEmpty() bool

func (rb *RingBuffer[T]) IsFull() bool

func (rb *RingBuffer[T]) Enqueue(x T)

func (rb *RingBuffer[T]) Dequeue() (T, bool)
```
