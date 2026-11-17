package main

import (
	"Nov_10_26/list"
	"fmt"
)

func merge(left *list.List[int], right *list.List[int]) *list.List[int] {
	result := list.NewList[int]()
	for !left.IsEmpty() && !right.IsEmpty() {
		fmt.Println(result.Data(), left.Data(), right.Data())

		firstLeft := left.First().Element()
		firstRight := right.First().Element()
		if firstLeft < firstRight {
			result.InsertNodeAsLast(left.DetachHead())
		} else {
			result.InsertNodeAsLast(right.DetachHead())
		}
	}
	if !left.IsEmpty() {
		result.Join(left)
	}
	if !right.IsEmpty() {
		result.Join(right)
	}
	return result
}

func main() {
	left := list.NewList[int]()
	left.AddBack(2)
	left.AddBack(3)
	left.AddBack(5)
	left.AddBack(6)

	fmt.Println(left.Data())

	right := list.NewList[int]()
	right.AddBack(1)
	right.AddBack(2)
	right.AddBack(4)

	fmt.Println(right.Data())

	result := merge(left, right)
	fmt.Println(result.Data())
}
