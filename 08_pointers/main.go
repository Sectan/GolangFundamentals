// A pointer holds the memory address of a value.
// More information:
//  - https://golang.org/ref/spec#Pointer_types
//  - https://golang.org/doc/faq#Pointers
//  - https://tour.golang.org/moretypes/1
package main

import "fmt"

func zeroWithPointer(x *int) {
	*x = 0
}

func zeroWithoutPointer(x int) {
	x = 0 // x is a local variable that is only visible inside the function
}

func main() {
	a := 10
	fmt.Println("Content of a:", a)
	fmt.Println("Address of a:", &a)
	var b = &a             // b is of type *int (address operator always returns type *T), b references a (referencing)
	fmt.Println("b:", b)   // Prints the address of a
	fmt.Println("*b:", *b) // Prints the content of a (dereferencing)

	// Assigning a value to a dereferenced pointer changes the value at the referenced address. (https://gobyexample.com/pointers)
	*b = 20
	fmt.Println("Content of a:", a)
	fmt.Println("*b:", *b)

	x := 5
	zeroWithPointer(&x)
	fmt.Println(x) // x is 0

	x = 5
	zeroWithoutPointer(x)
	fmt.Println(x) // x is 5
}

// When should we use pointers? (source: https://www.goinggo.net/2014/12/using-pointers-in-go.html)
// 1) In general, don’t share built-in type values with a pointer.
// 2) In general, share struct type values with a pointer unless the struct type has been implemented to behave like a primitive data value.
// 3) In general, don’t share reference type values with a pointer unless you are implementing an unmarshal type of functionality.
// 4) In general, create slices and maps of values when you can.

// Pointer use cases (source: http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/)
// 1) Variable must not be modified:
//     We do not have other option, but pass variable by value. So that variable cannot be modified downstream.
//
// 2) Variable is a large struct:
//     If variable is a large struct and performance is an issue, it's preferable to pass variable by pointer.
//
// 3) Variable is a map or slice:
//     Maps and slices are reference types in Go and should be passed by values.
