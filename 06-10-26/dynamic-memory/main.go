package main

import "fmt"

func main() {
	p := new(int) // type of p is *int, allocates one int to the Heap memory

	x := 9

	p = &x // redirecting p to point the address of x.

	fmt.Println(p, *p)

	p = nil // The Garbage Collector will free the memory on next GC cycle.

	a := make([]int, 10, 10)       // Allocates a slice with len and cap of 10 consisting of 10 0 (the zero value for int)
	fmt.Println(a, len(a), cap(a)) // [] 0 10

	a = nil // Freeing the memory backing a.
	fmt.Println(a, len(a), cap(a)) // [] 0 0

	fmt.Println(8, "abc") // 8 abc
}
