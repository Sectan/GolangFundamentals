// nil is a predeclared identifier representing the zero value for 
// a pointer, channel, func, interface, map, or slice type.
// More info: https://golang.org/pkg/builtin/#pkg-variables



package main

import "fmt"

func main() {
	var nilSlice []string // Idiomatic way to declare an empty slice
	zeroSlice := []string{}

	// The former declares a nil slice value, 
	// while the latter is non-nil but zero-length. 
	// They are functionally equivalent—their len and 
	// cap are both zero—but the nil slice is the preferred style.
	// More info: https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices

	fmt.Println(len(nilSlice))
	fmt.Println(nilSlice == nil)
	fmt.Printf("%#v\n", nilSlice)

	fmt.Println(len(zeroSlice))
	fmt.Println(zeroSlice == nil)
	fmt.Printf("%#v\n", zeroSlice)
}
