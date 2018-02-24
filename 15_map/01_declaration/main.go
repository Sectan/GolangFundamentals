// A map is an unordered group of elements of one type, called the element type,
// indexed by a set of unique keys of another type, called the key type.
// The value of an uninitialized map is nil.
// Golang Spec: https://golang.org/ref/spec#Map_types

package main

import "fmt"

func main() {
	var map1 = map[int]string{}
	fmt.Println(map1)
	fmt.Println(map1 == nil)

	map2 := make(map[int]string)
	fmt.Println(map2)
	fmt.Println(map2 == nil)

	// A nil map is equivalent to an empty map except that no elements may be added.
	var nilMap map[int]string
	fmt.Println(nilMap)
	fmt.Println(nilMap == nil)
}
