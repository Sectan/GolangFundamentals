// An array is a numbered sequence of elements of a single type,
// called the element type. The number of elements is called
// the length and is never negative.
// More info: https://golang.org/ref/spec#Array_types
// Copied some doc from: https://gobyexample.com/range

// Arrays are useful when planning the detailed layout of memory
// and sometimes can help avoid allocation, but primarily they are a
// building block for slices, the subject of the next section.
// An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types).
// More info: https://golang.org/doc/effective_go.html#arrays

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
	for i, v := range oneD {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	// Even multi dimensional arrays are possible
	var twoD [2][3]int
	fmt.Println(twoD)

	// The notation ... specifies an array length equal to the maximum element index plus one.
	players := [...]string{"Totti", "de Rossi"} // The maximum element index in this example is 1
	fmt.Println(len(players))                   // len = maximum element index + 1 = 1 + 1 = 2

	// Arrays can be compared
	fmt.Println(players == [...]string{"Totti", "de Rossi"})
}
