package main

import "fmt"

func main() {
	players := make(map[int]string)
	players[10] = "?"
	players[16] = "de Rossi"
	players[24] = "Florenzi"
	fmt.Println(players)

	players[10] = "Totti" // Replace the old value
	fmt.Println(players)
}
