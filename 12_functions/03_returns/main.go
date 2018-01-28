// A function can return any number of results.

// Go's return values may be named.  If so,
// they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// More information: https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
// More information: https://golang.org/ref/spec#Return_statements

package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func values() (valA, valB int) {
	return 10, 16
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// A return statement without arguments returns the named return values.
	// This is known as a "naked" return.
	// Naked return statements should be used only in short functions.
	// Naked return statements only work with named returns.
	return
}

func main() {
	a, b := swap("Totti", "Francesco")
	fmt.Println(a,b)
	valA, valB := values()
	fmt.Println(valA, valB)
	fmt.Println(split(17))
}