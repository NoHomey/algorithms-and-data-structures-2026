package main

import (
	"fmt"
	"math"
)

type Vector struct {
	x float64
	y float64
}

func NewVector(x, y float64) Vector {
	return Vector{x: x, y: y}
}

func (v Vector) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vector) String() string {
	return fmt.Sprintf("(x=%f, y=%f)", v.x, v.y)
}

func myPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("%d\n", x.(int))
	case float64:
		fmt.Printf("%f\n", x.(float64))
	case string:
		fmt.Println(x.(string))

	case fmt.Stringer:
		myPrint(x.(fmt.Stringer).String()) // myPrint(string)

	default:
		fmt.Printf("%+v\n", x)
	}
}

func main() {
	v := NewVector(3.2, 1.3)

	myPrint(v)
}
