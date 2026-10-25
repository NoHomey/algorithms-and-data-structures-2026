# Task

Find what is the problem in the code. It is very tiny and kinda hidden but it is also simple to fix :)

Hint 1: Look into the definition of the `type Program` and then look into the `type CodeBlock`
and think what happens once the `Execute` method for `CodeBlock` gets called.

Hint 2: What happens with the `instructions` of a `CodeBlock` once the `Execute` method gets called?

The above is about the actual problem a bug in the code.

There is another problem - an unwanted side effect that leads to an infinite recursuion:

The following print evaluates both cases of the conditional expression
```go
fmt.Println("Cond", c.condition.Eval(p, m), c.trueCase.Eval(p, m), c.falseCase.Eval(p, m))
```

If you remove it the following program:
```go
p := CreateProgram(map[string]CodeBlock{
	"div3": NewCodeBlock([]Instruction{
		Return{exp: Conditional{
			condition: GreaterThanExp(Var(0), Const(2)),
			trueCase:  Sum(Const(1), FuncCallExp("div3", []Evaluatable{Diff(Var(0), Const(3))})),
			falseCase: Const(0),
		}},
	}),
	"main": NewCodeBlock([]Instruction{
		Return{exp: FuncCallExp("div3", []Evaluatable{Const(16)})},
	}),
})
fmt.Println(p.Execute())
```
will print 5 wich is correct because when doing an integer division of 16 by 3 the result is 5.

For that you must also implement a `type Condition` and support at least `>`, `<`, `==` as conditional operators!

And of course change the `type Conditional` to:
```go
type Conditional struct {
	condition Condition
	trueCase  Evaluatable
	falseCase Evaluatable
}
```

## For exercise
Instead of `div3` which hard-codes on what we divide make a true integer division function `div` that does an integer division of it's first argument by it's second.

Hint the program should be:
```go
p := CreateProgram(map[string]CodeBlock{
	"div": NewCodeBlock([]Instruction{
		Return{exp: Conditional{
			condition: GreaterThanExp(Var(0), Diff(Var(1), Const(1)))
			trueCase:  Sum(Const(1), FuncCallExp("div", []Evaluatable{Diff(Var(0), Var(1)), Var(1)})),
			falseCase: Const(0),
		}},
	}),
	"main": NewCodeBlock([]Instruction{
		Return{exp: FuncCallExp("div", []Evaluatable{Const(16), Const(3)})},
	}),
})
fmt.Println(p.Execute())
```
And it should again print out 5 as above.
