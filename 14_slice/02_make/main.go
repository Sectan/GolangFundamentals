// The built-in function make takes a type T, which must be a slice, map or 
// channel type, optionally followed by a type-specific list of expressions. 
// It returns a value of type T (not *T). The memory is initialized.
// Golang Spec: https://golang.org/ref/spec#Making_slices_maps_and_channels

package main

import "fmt"

func main() {
	players := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	fmt.Println(players == nil)
	players[0] = "Totti"
	players[1] = "de Rossi"
	players[2] = "Florenzi"
	fmt.Println(players)
	fmt.Println("Length of slice players:", len(players))

	// This does not work. It exceeds the lenght of the slice.
	//players[3] = "Pellegrini"

	// We need to use append.
	// The variadic function append appends zero or more values x to s of type S, 
	// which must be a slice type, and returns the resulting slice, also of type S.
	// Golang Spec: https://golang.org/ref/spec#Appending_and_copying_slices
	players = append(players, "Pellegrini", "Montella")
	fmt.Println(players)
	fmt.Println("Length of slice players:", len(players))

	// This is the idiomatic way of using make
	newSlice := make([]string, 30)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}