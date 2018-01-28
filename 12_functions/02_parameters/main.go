// A function can take zero or more arguments.
// Everything in Go is passed by value.
// That is, a function always gets a copy of the thing being passed,
// as if there were an assignment statement assigning the value
// to the parameter. For instance, passing an int value to a function
// makes a copy of the int, and passing a pointer value makes a copy of
// the pointer, but not the data it points to.
// More information:
//  https://golang.org/ref/spec#Calls
//  https://golang.org/doc/faq#pass_by_value

package main

import (
	"fmt"
)

func printFullName(firstName string, lastName string) {
	fmt.Println("Full name:", firstName, lastName)
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func printFName(firstName, lastName string) {
	fmt.Println("Full name:", firstName, lastName)
}

func main() {
	printFullName("Francesco", "Totti")
	printFName("Francesco", "Totti")
}
