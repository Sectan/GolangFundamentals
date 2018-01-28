// To declare a function we use the func keyword.
// It is not possible to overload a function, see: https://golang.org/doc/faq#overloading
// The parameters and the results form the signature of a function
// More info: https://golang.org/ref/spec#Function_types

package main

import (
	"fmt"
)

func aFunction() {
	fmt.Println("func aFunction called")
}

/*
This does not compile, because function overloading is not permitted.

func aFunction(i float32, j float32) float32 {

}*/


func main() {
	aFunction()
}