// The blank identifier can be assigned or declared with
// any value of any type, with the value discarded harmlessly.
// It's a bit like writing to the Unix /dev/null file: it
// represents a write-only value to be used as a place-holder
// where a variable is needed but the actual value is irrelevant.

// More information: https://golang.org/doc/effective_go.html#blank

package main

import (
	"fmt"
	"os"
	_ "log" // Unused package declaration
)

func main() {
	path := "Foooooo"

	// Use the blank identifier to discard the irrelevant value
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}

	var i int  // Variable declared but not used
	_ = i      // Make sure that the code compiles
}