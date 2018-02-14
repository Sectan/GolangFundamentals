// The bytes package provides some useful functions for handling byte slices
// More info: https://golang.org/pkg/bytes/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	byte1 := []byte{3, 2, 1}
	byte2 := []byte{84, 111, 116, 116, 105}
	byte3 := []byte{3, 2, 1}

	// Check if slices are equal
	fmt.Println(bytes.Equal(byte1, byte3))
	fmt.Println(bytes.Equal(byte2, []byte("Totti")))
	fmt.Println(bytes.Equal(byte1, byte2))

	// Check if the subslice is contained in the slice
	fmt.Println(bytes.Contains(byte2, []byte("ti")))
}
