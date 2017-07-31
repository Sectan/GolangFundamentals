package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if true {
		fmt.Println("Condition true")
	}

	// The expression may be preceded by a simple statement,
	// which executes before the expression is evaluated.
	if x := rand.Intn(15); x < 5 {
		fmt.Println(x)
	} else if x < 10 {
		fmt.Println(x)
	} else {
		fmt.Println(x)
	}
}
