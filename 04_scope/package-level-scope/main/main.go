package main

import (
	"fmt"
	"github.com/sectan/GolangFundamentals/04_scope/package-level-scope/valpkg"
)

var foo = 1

func main() {
	fmt.Println(valpkg.FirstName)
	valpkg.PrintValues()
	fmt.Println(foo)
	fmt.Println(bar)
}

var bar = 1