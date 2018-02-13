// A slice is a descriptor for a contiguous segment of an underlying array and
// provides access to a numbered sequence of elements from that array. A slice
// type denotes the set of all slices of arrays of its element type. The value
// of an uninitialized slice is nil.
// Golang Spec: https://golang.org/ref/spec#Slice_types
//
// Slices wrap arrays to give a more general, powerful, and convenient interface
// to sequences of data. Most array programming in Go is done with slices
// rather than simple arrays.
// More info: https://golang.org/doc/effective_go.html#slices

package main

import "fmt"

func main() {
	var player []string
	fmt.Println("Content of slice player:", player)
	fmt.Println(player == nil)
	fmt.Println("Length of slice player:", len(player))
	fmt.Println("Capacity of slice player:", cap(player))
	player = append(player, "Totti")
	fmt.Println("Content of slice player:", player)
	fmt.Println(player == nil)
	fmt.Println("Length of slice player:", len(player))
	fmt.Println("Capacity of slice player:", cap(player))

	players := []string{"Totti"}
	fmt.Println("Content of slice players:", players)
	fmt.Println("Length of slice players:", len(players))
	fmt.Println("Capacity of slice players:", cap(players))
	players = append(players, "de Rossi")
	fmt.Println("Content of slice players:", players)
	fmt.Println("Length of slice players:", len(players))
	fmt.Println("Capacity of slice players:", cap(players))
}
