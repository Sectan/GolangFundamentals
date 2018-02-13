// The variadic function append appends zero or
// more values x to s of type S, which must be
// a slice type, and returns the resulting slice, also of type S.
// Golang Spec: https://golang.org/ref/spec#Appending_and_copying_slices

// The append built-in function appends elements to the end of a slice.
// More info: https://golang.org/pkg/builtin/#append

// The built-in append function appends elements to the end of a slice:
// 	- if there is enough capacity, the underlying array is reused;
// 	- if not, a new underlying array is allocated and the data is copied over.
// More info: http://yourbasic.org/golang/append-explained/

package main

import "fmt"

func main() {
	fistName := []string{"Francesco"}
	lastName := []string{"Totti"}
	fullName := append(fistName, lastName...) // append a slice
	fmt.Println(fullName)

	fullName = append(fistName, "Totti")
	fmt.Println(fullName) // append a single element

	msg := []string{"these", "aren't"}
	fullMsg := append(msg, "the", "droids", "you're", "looking", "for")
	fmt.Println(fullMsg)

	fullMsg2 := append(fullMsg[:3], fullMsg[3:]...)
	fmt.Println(fullMsg2)

	// As a special case, it is legal to append a string to a byte slice, like this:
	someSlice := append([]byte("hello "), "world"...)
	fmt.Println(someSlice)
}
