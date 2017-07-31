// The Go for loop is similar to—but not the same as—C's.
// It unifies for and while and there is no do-while.
// There are three forms, only one of which has semicolons.
//
// Like a C for
// for init; condition; post { }
//
// Like a C while
// for condition { }
//
// Like a C for(;;)
// for { }

// Source inspired by https://gobyexample.com/for

package main

import "fmt"

func main() {
	// A classic initial/condition/after `for` loop.
	for i := 0; i <= 9; i++ {
		fmt.Println("i = ", i)
	}

	// The most basic type, with a single condition.
	j := 1
	for j <= 3 {
		fmt.Println("j = ", j)
		j = j + 1
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}