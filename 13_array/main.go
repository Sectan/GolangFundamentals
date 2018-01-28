// An array is a numbered sequence of elements of a single type, 
// called the element type. The number of elements is called 
// the length and is never negative.
// More info: https://golang.org/ref/spec#Array_types
// Copied some doc from: https://gobyexample.com/range

package main

import "fmt"

func main() {
	var oneD [10]int
	fmt.Println(oneD)
	fmt.Println("Length of oneD: ", len(oneD))

	for i := 0; i < 10; i++ {
		oneD[i] = i
	}

	// range on arrays provides both the index and value for each entry.
	// See Golang Spec: "For statements with range clause"
	for i, v := range(oneD) {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	// Even multi dimensional arrays are possible
	var twoD [2][3]int
	fmt.Println(twoD)
	}