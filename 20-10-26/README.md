# Task 1

Implement the constructor functions `Const`, `Variable`, `Sum`, `Diff`, `Mul` for the `*Exp` type so that instead of having to manually construct `Exp`ressions like the one we did in class:

```go
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
```

we can write the much nicer and simpler:

```go
	exp := Mul(Sum(Const(4), Sum(Const(5), Const(2))), Const(3))
```

# Task 2

Add support for expressions of type variable who's values should be taken from the memory at runtime when evaluating an expression.

For example if the `memory` is `[4, 5, 2]` then the following expression `Mul(Sum(Variable(0), Sum(Variable(1), Variable(2))), Const(3))` should evaluate to `33` when given that memory.
