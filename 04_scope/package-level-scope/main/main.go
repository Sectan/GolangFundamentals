package main

import (
	"fmt"
	"github.com/sectan/GolangFundamentals/04_scope/package-level-scope/valpkg"
)

func main() {
	fmt.Println(valpkg.FirstName)
	valpkg.PrintValues()
}