// More information on this topic: https://blog.golang.org/constants
package main

import "fmt"

// untyped string constant
const hello  = "Hello, 世界"
// After this declaration, hello is also an untyped string constant.
// An untyped constant is just a value, one not yet given a defined
// type that would force it to obey the strict rules that prevent
// combining differently typed values.

// typed string constant
const typedHello string = "Hello, 世界"
// Notice that the declaration of typedHello has an explicit
// string type before the equals sign. This means that typedHello
// has Go type string, and cannot be assigned to a
// Go variable of a different type.

func main() {
	type MyString string // Define an own type
	var m MyString
	// m = typedHello // Type error
	// This is because typedHello has an explicit string type
	// and does not work with the type MyString
	m = hello
	fmt.Println(m)
}
