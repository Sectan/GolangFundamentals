// This example shows the underlying type of a custom type based on a struct

package main

import "fmt"
import "reflect"

type player struct {
	name string
}

func main() {
	var player1 player
	fmt.Println("Defined type:", reflect.TypeOf(player1))
	fmt.Println("Underlying type:", reflect.TypeOf(player1).Kind())

	player2 := new(player)
	fmt.Println("Defined type:", reflect.TypeOf(player2))
	fmt.Println("Underlying type:", reflect.TypeOf(player2).Kind())
}
