package main

import (
	"sort"
	"github.com/neilramaswamy/generics-playground/rope"
)

func sortFunc[T any](a []T, less func(x, y T) bool) {	
	sort.Slice(a, func(i int, j int) bool {
		return less(a[i], a[j])
	})
}

// func printThing[T MyUnion](thing T) {
// 	fmt.Printf("thing is %v", thing)
	
// 	// x := []uint32{"hello", "goodbye", "aaa", "ccc"}
// 	x := []uint32{0-1, 1, 0}
// 	sortFunc(x, func(a, b uint32) bool{ return a < b })
// 	fmt.Printf("%#v", x)
// }

func main() {
	// printThing(uint32(2))
	 _ = rope.Rope[string]{}
}