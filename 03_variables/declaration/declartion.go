// For more information: https://golang.org/ref/spec#Variable_declarations
package main

import "fmt"

// A variable declaration creates one or more variables, binds
// corresponding identifiers to them, and gives each a type and an initial value.

// A full variable declartion looks like that: var name type
var foo int

// It is also possible to declare multiple variables at once
var a, b, c int

// Or like that
var (
	a1 string
	a2 uint64
)

// If a type is present, each variable is given that type.
// Otherwise, each variable is given the type of the corresponding
// initialization value in the assignment.
var i = 42 // i is int

// Declare and initialize a variable
var test int = 10

// We can also declare and initialise multiple variables
var m, n int = 1, 2 // m = 1, n = 2

func main() {
	// There is also a short variable declaration

	// Short variable declarations may appear only inside functions.
	// In some contexts such as the initializers for "if", "for", or
	// "switch" statements, they can be used to declare
	// local temporary variables.

	j := 10 // j is int
	fmt.Println(j)
	fmt.Println(n)
}
