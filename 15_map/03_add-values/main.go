package main

import "fmt"

func main() {
	players := make(map[int]string)
	players[10] = "Totti"
	players[16] = "de Rossi"
	players[24] = "Florenzi"

	fmt.Println(players)
}
