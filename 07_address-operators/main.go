// For an operand x of type T, the address operation &x generates a pointer of type *T to x.
// More information: https://golang.org/ref/spec#Address_operators
package main

import "fmt"

func main() {
	var player string
	fmt.Print("Enter your favourite football player: ")
	fmt.Scan(&player)
	fmt.Println(player, "is your favourite football player!")

	fmt.Println("The memory address of the variable player is:", &player)

	/*
	var x *int = nil
	*x   // causes a run-time panic
	&*x  // causes a run-time panic
	*/
}
