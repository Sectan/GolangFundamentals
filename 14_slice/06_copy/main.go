// The function copy copies slice elements from a source src to a
// destination dst and returns the number of elements copied. Both
// arguments must have identical element type T and must be
// assignable to a slice of type []T.
// Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
//
// copy(dst, src []T) int
// copy(dst []byte, src string) int
//
// Golang spec: https://golang.org/ref/spec#Appending_and_copying_slices

package main

import "fmt"

func main() {
	base := [...]int{0, 1, 2, 3, 4, 5}
	i := make([]int, 5)
	b := make([]byte, 9)

	c1 := copy(i, base[0:])
	fmt.Println(i, "Number of elements copied from base:", c1)

	c2 := copy(b, "Francesco Totti")
	fmt.Println(b, "Number of elements copied from string \"Francesco Totti\"", c2)
}
