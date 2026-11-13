package main

import (
	"Nov_10_26/list"
	"fmt"
)

func main() {
	l := list.NewList[int]()

	// []
	fmt.Println(l.Data())
	l.AddFront(3)
	// [3]
	fmt.Println(l.Data())
	l.AddFront(2)
	// [2, 3]
	fmt.Println(l.Data())
	l.AddFront(1)
	// [1, 2, 3]
	fmt.Println(l.Data())

	l.AddBack(4)
	// [1, 2, 3, 4]
	fmt.Println(l.Data())

	iter := l.Last()
	// 4
	fmt.Println(iter.Element())
	iter = iter.Next()
	// 3
	fmt.Println(iter.Element())
	l.AddBefore(iter, 5)
	// [1, 2, 5, 3, 4]
	fmt.Println(l.Data())

	l.AddAfter(iter, 6)
	// [1, 2, 5, 3, 6, 4]
	fmt.Println(l.Data())

	elem, ok := l.RemoveFront()
	// 1, true
	fmt.Println(elem, ok)
	// [2, 5, 3, 6, 4]
	fmt.Println(l.Data())

	elem, ok = l.RemoveBack()
	// 4, true
	fmt.Println(elem, ok)
	// [2, 5, 3, 6]
	fmt.Println(l.Data())

	elem = l.RemoveNode(iter)
	// 3
	fmt.Println(elem)
	// [2, 5, 6]
	fmt.Println(l.Data())

	// 3
	fmt.Println(l.Size())

	l.Reverse()
	// [6, 5, 2]
	fmt.Println(l.Data())
	// [2, 5, 6]
	fmt.Println(l.DataReverse())
}
