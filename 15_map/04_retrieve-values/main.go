// Elements may be retrieved with index expressions.

package main

import "fmt"

func main() {
	players := map[int]string{
		10: "Totti",
		16: "de Rossi",
		24: "Florenzi",
	}

	player, ok := players[10]
	fmt.Println("Value retrieved:", ok, "\nValue:", player)

	// Sometimes you need to distinguish a missing entry from a zero value.
	// You can discriminate with a form of multiple assignment.
	// For obvious reasons this is called the “comma ok” idiom.
	player, ok = players[100]
	fmt.Println("Value retrieved:", ok, "\nValue:", player)
	fmt.Printf("Type of player: %T\nType of ok %T\n", player, ok)

	_, present := players[100]
	fmt.Println("Is value present:", present)
}
