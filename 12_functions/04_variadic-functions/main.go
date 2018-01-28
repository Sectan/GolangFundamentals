// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.

// The final incoming parameter in a function
// signature may have a type prefixed with ... .
// A function with such a parameter is called variadic
// and may be invoked with zero or more arguments for that parameter.

// If f is variadic with a final parameter p of type ...T,
// then within f the type of p is equivalent to type []T.
// If f is invoked with no actual arguments for p, the value passed to p is nil.
// Otherwise, the value passed is a new slice of type []T.

package main

import (
	"fmt"
)

func Greeting(greetingMsg string, names ...string) {
	for _, name := range names {
		fmt.Println(greetingMsg, name)
	}
}

/*
This won't compile, because ... is not used as the final parameter.
func Foo(a ...string, b string) {

}
*/

func main() {
	Greeting("Ciao", "Francesco", "Daniele", "Alessandro")

	var names []string
	names = []string{"Francesco", "Daniele", "Alessandro"}
	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	Greeting("Ciao", names...)
}
