// Slice expressions construct a substring or slice
// from a string, array, pointer to array, or slice.
//
// Simple slice expressions
// For a string, array, pointer to array, or slice a, the primary expression:
// 	a[low : high]
//
// Full slice expressions
// For an array, pointer to array, or slice a (but not a string), the primary expression
//  a[low : high : max]
//
// Golang spec: https://golang.org/ref/spec#Slice_expressions

package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a[2:]) // same as a[2 : len(a)]
	fmt.Println(a[:3]) // same as a[0 : 3]
	fmt.Println(a[:])  // same as a[0 : len(a)]

	t := a[1:3:5]
	fmt.Println("len of t:", len(t))
	fmt.Println("cap of t:", cap(t))
	fmt.Println("elements of t:", t)
}
