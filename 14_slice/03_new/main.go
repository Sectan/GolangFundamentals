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
}

https://tip.golang.org/doc/go1.2#three_index
https://medium.com/golangspec/slice-expressions-in-go-963368c20765