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
	size  int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity < 1 {
		panic(fmt.Sprintf("expecting capacity to be positive integer, got: %d", capacity))
	}
	return &RingBuffer[T]{
		data:  make([]T, capacity),
		first: 0,
		last:  0,
		size:  0,
	}
}

func NewRingBufferFromSlice[T any](data []T) *RingBuffer[T] {
	return &RingBuffer[T]{
		data:  data,
		first: 0,
		last:  len(data) - 1,
		size:  len(data),
	}
}

func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.size == 0
}

func (rb *RingBuffer[T]) IsFull() bool {
	return rb.size == cap(rb.data)
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
	rb.size++
}

func (rb *RingBuffer[T]) Dequeue() (T, bool) {
	if rb.IsEmpty() {
		var tmp T
		return tmp, false
	}
	x := rb.data[rb.first]
	rb.first = (rb.first + 1) % cap(rb.data)
	rb.size--
	return x, true
}

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

type Memory struct {
	array *Array[int]
}

func Empty() Memory {
	return Memory{array: NewArray[int]()}
}

func FromValues(vals []int) Memory {
	return Memory{array: NewArrayFromSlice(vals)}
}

func (m Memory) Read(index uint) int {
	val, ok := m.array.Get(index)
	if !ok {
		panic(fmt.Sprintf("%d index is not a valid memory index", index))
	}
	return val
}

func (m Memory) Write(p Program, index uint, exp Evaluatable) {
	m.array.Set(index, exp.Eval(p, m))
}

func (m Memory) CreateVariable(p Program, exp Evaluatable) {
	m.array.Push(exp.Eval(p, m))
}

type Evaluatable interface {
	Eval(Program, Memory) int
}

type Const int

func (c Const) Eval(p Program, m Memory) int {
	return int(c)
}

type Var uint

func (v Var) Eval(p Program, m Memory) int {
	return m.Read(uint(v))
}

type Neg struct {
	exp Evaluatable
}

func NegExp(exp Evaluatable) Neg {
	return Neg{exp}
}

func (neg Neg) Eval(p Program, m Memory) int {
	return -neg.exp.Eval(p, m)
}

type BinOpSymbol byte

const (
	Plus     = BinOpSymbol('+')
	Minus    = BinOpSymbol('-')
	Multiply = BinOpSymbol('*')
)

type BinOp struct {
	op    BinOpSymbol
	left  Evaluatable
	right Evaluatable
}

func (binOp BinOp) Eval(p Program, m Memory) int {
	switch binOp.op {
	case Plus:
		return binOp.left.Eval(p, m) + binOp.right.Eval(p, m)
	case Minus:
		return binOp.left.Eval(p, m) - binOp.right.Eval(p, m)
	case Multiply:
		return binOp.left.Eval(p, m) * binOp.right.Eval(p, m)
	}
	panic(fmt.Sprintf("Unknown BinOpType: %c", binOp.op))
}

func BinOpExp(left, right Evaluatable, op BinOpSymbol) BinOp {
	return BinOp{
		op:    op,
		left:  left,
		right: right,
	}
}

func Sum(left, right Evaluatable) BinOp {
	return BinOpExp(left, right, Plus)
}

func Diff(left, right Evaluatable) BinOp {
	return BinOpExp(left, right, Minus)
}

func Mul(left, right Evaluatable) BinOp {
	return BinOpExp(left, right, Multiply)
}

type Instruction interface {
	Execute(p Program, m Memory) (val int, isReturn bool)
}

type CreateVariable struct {
	exp Evaluatable
}

func (cv CreateVariable) Execute(p Program, m Memory) (int, bool) {
	m.CreateVariable(p, cv.exp)
	return 0, false
}

type Assignment struct {
	varIndex uint
	exp      Evaluatable
}

func (a Assignment) Execute(p Program, m Memory) (int, bool) {
	m.Write(p, a.varIndex, a.exp)
	return 0, false
}

type Return struct {
	exp Evaluatable
}

func (r Return) Execute(p Program, m Memory) (int, bool) {
	return r.exp.Eval(p, m), true
}

type CodeBlock struct {
	instructions []Instruction
}

func NewCodeBlock(instructions []Instruction) CodeBlock {
	return CodeBlock{
		instructions: instructions,
	}
}

func (code CodeBlock) Execute(p Program, m Memory) int {
	var lastValue int
	instructionsCopy := make([]Instruction, len(code.instructions))
	copy(instructionsCopy, code.instructions)
	instructions := NewRingBufferFromSlice(instructionsCopy)
	for !instructions.IsEmpty() {
		instruction, ok := instructions.Dequeue()
		if !ok {
			panic("Something went wrong")
		}
		val, isReturn := instruction.Execute(p, m)
		if isReturn {
			return val
		}
		lastValue = val
	}
	return lastValue
}

type FuncCall struct {
	name string
	args []Evaluatable
}

func FuncCallExp(name string, args []Evaluatable) FuncCall {
	return FuncCall{
		name: name,
		args: args,
	}
}

func (funcCall FuncCall) Eval(p Program, scope Memory) int {
	callMemory := Empty()
	for _, arg := range funcCall.args {
		callMemory.CreateVariable(p, Const(arg.Eval(p, scope)))
	}
	return p.Call(funcCall.name, callMemory)
}

type Program struct {
	functions map[string]CodeBlock
}

func CreateProgram(functions map[string]CodeBlock) Program {
	return Program{functions: functions}
}

func (p Program) Call(name string, scope Memory) int {
	code := p.functions[name]
	return code.Execute(p, scope)
}

func (p Program) Execute() int {
	return p.Call("main", Empty())
}

type Comparator byte

const (
	LessThan    = Comparator('<')
	Equal       = Comparator('=')
	GreaterThan = Comparator('>')
)

type Condition struct {
	comp  Comparator
	left  Evaluatable
	right Evaluatable
}

func (c Condition) Eval(p Program, m Memory) int {
	switch c.comp {
	case LessThan:
		if c.left.Eval(p, m) < c.right.Eval(p, m) {
			return 1
		}
		return 0
	case Equal:
		if c.left.Eval(p, m) == c.right.Eval(p, m) {
			return 1
		}
		return 0
	case GreaterThan:
		if c.left.Eval(p, m) > c.right.Eval(p, m) {
			return 1
		}
		return 0
	default:
		panic(fmt.Sprintf("Unknown comparator %c", byte(c.comp)))
	}
}

func ConditionExp(comp Comparator, left, right Evaluatable) Condition {
	return Condition{
		comp:  comp,
		left:  left,
		right: right,
	}
}

func LessThanExp(left, right Evaluatable) Condition {
	return ConditionExp(LessThan, left, right)
}

func EqualExp(left, right Evaluatable) Condition {
	return ConditionExp(Equal, left, right)
}

func GreaterThanExp(left, right Evaluatable) Condition {
	return ConditionExp(GreaterThan, left, right)
}

type Conditional struct {
	condition Condition
	trueCase  Evaluatable
	falseCase Evaluatable
}

func (c Conditional) Eval(p Program, m Memory) int {
	if c.condition.Eval(p, m) != 0 {
		return c.trueCase.Eval(p, m)
	}
	return c.falseCase.Eval(p, m)
}

func main() {
	/*
		m := FromValues([]int{4, 5, 2, 3})
		exp := Mul(Sum(Var(0), Sum(Var(1), Var(2))), Var(3))
		fmt.Println(exp.Eval(m))
	*/
	/*
		p := NewCodeBlock([]Instruction{
			CreateVariable{exp: Const(4)},                     // x0 := 4
			CreateVariable{exp: Const(5)},                     // x1 := 5
			CreateVariable{exp: Const(2)},                     // x2 := 2
			CreateVariable{exp: Sum(Var(1), Var(2))},          // x3 := x1 + x2 = 7
			Assignment{varIndex: 3, exp: Sum(Var(0), Var(3))}, // x3 = x0 + x3 = 4 + 7 = 11
			Return{exp: Mul(Var(3), Const(3))},                // return x3 *3 ~ return 11 * 3 ~ return 33
		})

		m := Empty()
		res := p.Execute(m)
		fmt.Println(res)
		fmt.Println(m.array.data)
	*/

	/*
		func f(x, y int) int {
			return x * y + 2
		}

		func f(x0, x1) int {
			x2 := x0 * x1
			return x2 + 2
		}

		func g(x0 int) int {
			return x0 - 3
		}

		func main() int {
			x0 := 5
			x1 := 7
			x2 := f(x0, x1)
			x3 := g(x2)
			return x3
		}

	*/

	/*
		p := CreateProgram(map[string]CodeBlock{
			"f": NewCodeBlock([]Instruction{
				CreateVariable{exp: Mul(Var(0), Var(1))},
				Return{exp: Sum(Var(2), Const(2))},
			}),
			"g": NewCodeBlock([]Instruction{
				Return{exp: Diff(Var(0), Const(3))},
			}),
			"main": NewCodeBlock([]Instruction{
				CreateVariable{exp: Const(5)},
				CreateVariable{exp: Const(7)},
				CreateVariable{exp: FuncCallExp("f", []Evaluatable{Var(0), Var(1)})},
				CreateVariable{exp: FuncCallExp("g", []Evaluatable{Var(2)})},
				Return{exp: Var(3)},
			}),
		})
		fmt.Println(p.Execute())
	*/

	/*
		func f(x0) int {
			return x0 > 2 ? 1 + f(x0 - 3) : 0
		}

	*/
	p := CreateProgram(map[string]CodeBlock{
		"div3": NewCodeBlock([]Instruction{
			Return{exp: Conditional{
				condition: GreaterThanExp(Var(0), Const(2)),
				trueCase:  Sum(Const(1), FuncCallExp("div3", []Evaluatable{Diff(Var(0), Const(3))})),
				falseCase: Const(0),
			}},
		}),
		"div": NewCodeBlock([]Instruction{
			Return{exp: Conditional{
				condition: GreaterThanExp(Var(0), Diff(Var(1), Const(1))),
				trueCase:  Sum(Const(1), FuncCallExp("div", []Evaluatable{Diff(Var(0), Var(1)), Var(1)})),
				falseCase: Const(0),
			}},
		}),
		"mod": NewCodeBlock([]Instruction{
			CreateVariable{exp: FuncCallExp("div", []Evaluatable{Var(0), Var(1)})},
			CreateVariable{exp: Mul(Var(2), Var(1))},
			Return{exp: Diff(Var(0), Var(3))},
		}),
		"main": NewCodeBlock([]Instruction{
			Return{exp: FuncCallExp("mod", []Evaluatable{Const(19), Const(5)})},
		}),
	})
	fmt.Println(p.Execute())
}
