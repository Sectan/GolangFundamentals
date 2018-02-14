// new is a built-in function that allocates memory, but unlike its namesakes
// in some other languages it does not initialize the memory, it only zeros it.
// Golang Spec: https://golang.org/ref/spec#Allocation
// Effecitve Go: https://golang.org/doc/effective_go.html#allocation_new

package main

import "fmt"

func main() {
	playersPointer := new([]string)
	players := *playersPointer
	fmt.Println(players == nil)
	fmt.Println(len(players))
	fmt.Println(cap(players))
	players = append(players, "Totti", "de Rossi", "Florenzi")
	fmt.Println(len(players))
	fmt.Println(cap(players))

	var someSlice = []int{}
	someSlice = new([100]int)[:50] // same as make([]int, 50, 100)
	// new returns a pointer to an array
	// Golang Spec states: If a is a pointer to an array, a[low : high] is shorthand for (*a)[low : high].

	fmt.Println(someSlice)
	fmt.Println(len(someSlice))
	fmt.Println(cap(someSlice))
}
