// Go functions may be closures.
// A closure is a function value that references
//  variables from outside its body. The function
// may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.
// More information:
//	- https://tour.golang.org/moretypes/25
//	- https://golang.org/ref/spec#Declarations_and_scope

package main

import (
	"fmt"
)

// package scope variable
var x int

func increment() int {
	x++
	return x
}

func main() {
	a := 10
	fmt.Println(a)
	{
		fmt.Println(a)
		obiWan := "these aren't the droids you're looking for"
		fmt.Println(obiWan)
	}
	// fmt.Println(obiWan) // obiWan is not available in this scope

	fmt.Println(increment())
	fmt.Println(increment())
	
	x++
	fmt.Println(x)
}
