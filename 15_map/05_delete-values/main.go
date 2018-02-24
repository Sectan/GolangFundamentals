// The built-in function delete removes the element with key k from a map m.
// The type of k must be assignable to the key type of m.
// Golang Spec: https://golang.org/ref/spec#Deletion_of_map_elements

package main

import "fmt"

func main() {
	players := map[int]string{
		10: "Totti",
		16: "de Rossi",
		24: "Florenzi",
		25: "Bruno Peres",
	}
	fmt.Println(players)
	delete(players, 25) // Remove element players[25] from players
	fmt.Println(players)

	var foo map[string]int
	// If the map m is nil or the element m[k] does not exist, delete is a no-op.
	delete(foo, "foo")
}
