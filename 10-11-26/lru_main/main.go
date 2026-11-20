package main

import (
	"Nov_10_26/lru"
	"fmt"
)

func main() {
	cache := lru.NewCache[string, int](5)

	cache.Set("abc", 9)
	cache.Set("a", 65)
	cache.Set("bc", 11)
	cache.Set("iu", 12)
	cache.Set("cap", 5)

	fmt.Println(cache.Get("cap"))

	cache.Set("123", 123)

	fmt.Println(cache.Get("abc"))

}
