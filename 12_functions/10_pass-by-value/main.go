// Everything in Go is passed by value.
// Golang Spec: "the parameters of the call are passed by value
// to the function and the called function begins execution."
// Golang Spec: https://golang.org/ref/spec#Calls

// There are no reference variables in Golang: https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
// Therefore pass by reference does not exist.

package main

import "fmt"

func main() {
	player := "Totti"
	changePlayer(player)
	fmt.Println(player)

	changePlayer2(&player)
	fmt.Println(player)
}

func changePlayer(p string) {
	p = "de Rossi"
}

func changePlayer2(p *string) {
	fmt.Println("Address of player: ", p)
	fmt.Println("Address of p: ", &p)
	*p = "de Rossi"
}
