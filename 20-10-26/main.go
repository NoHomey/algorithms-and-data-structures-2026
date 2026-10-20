package main

import "fmt"

/*
import "fmt"

func fib() func() int {
	prev := 0
	curr := 1
	return func() int {
		next := prev + curr
		prev = curr
		curr = next
		return next
	}
}

type fibGenerator struct {
	prev int
	curr int
}

func NewFibGenerator() fibGenerator {
	return fibGenerator{
		prev: 0,
		curr: 1,
	}
}

func (g *fibGenerator) next() int {
	next := g.prev + g.curr
	g.prev = g.curr
	g.curr = next
	return next
}

func main() {
		nextFib := fib()
		for i := 0; i < 5; i++ {
			fmt.Println(nextFib())
		}
	g := NewFibGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(g.next())
	}
}
*/

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

type Memory = Array[int]

/*
func (m *Memory) Read(index int) int {
	return
}
*/

type ExpType = uint

const (
	ValueType = ExpType(iota)
	VariableType
	NegateType
	BinOpType
)

type Exp struct {
	expType  ExpType
	val      int
	variable uint
	negate   *Exp
	binOp    BinOp
}

func (exp Exp) Eval(m *Memory) int {
	switch exp.expType {
	case ValueType:
		return exp.val
	case NegateType:
		return -exp.negate.Eval(m)
	case BinOpType:
		return exp.binOp.Eval(m)
	}
	panic(fmt.Sprintf("Unknown expression type: %d", exp.expType))
}

type BinOpSymbol byte

const (
	Plus     = BinOpSymbol('+')
	Minus    = BinOpSymbol('-')
	Multiply = BinOpSymbol('*')
)

type BinOp struct {
	op    BinOpSymbol
	left  *Exp
	right *Exp
}

func (binOp BinOp) Eval(m *Memory) int {
	switch binOp.op {
	case Plus:
		return binOp.left.Eval(m) + binOp.right.Eval(m)
	case Minus:
		return binOp.left.Eval(m) - binOp.right.Eval(m)
	case Multiply:
		return binOp.left.Eval(m) * binOp.right.Eval(m)
	}
	panic(fmt.Sprintf("Unknown BinOpType: %c", binOp.op))
}

func Const(x int) *Exp {
	// implement
	return nil
}

func Variable(index uint) *Exp {
	// implement
	return nil
}

func Neg(exp *Exp) *Exp {
	// implement
	return nil
}

func Sum(left, right *Exp) *Exp {
	// implement
	return nil
}

func Diff(left, right *Exp) *Exp {
	// implement
	return nil
}

func Mul(left, right *Exp) *Exp {
	// implement
	return nil
}

func main() {
	exp := BinOp{
		op: Multiply,
		left: &Exp{
			expType: BinOpType,
			binOp: BinOp{
				op: Plus,
				left: &Exp{
					expType: ValueType,
					val:     4,
				},
				right: &Exp{
					expType: BinOpType,
					binOp: BinOp{
						op: Plus,
						left: &Exp{
							expType: ValueType,
							val:     5,
						},
						right: &Exp{
							expType: ValueType,
							val:     2,
						},
					},
				},
			},
		},
		right: &Exp{
			expType: ValueType,
			val:     3,
		},
	}

	res := Mul(Sum(Const(4), Diff(Const(5), Const(2))), Const(3)).Eval(nil)

	// 33 21
	fmt.Println(exp.Eval(nil), res)
}
