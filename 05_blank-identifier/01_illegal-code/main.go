package main

// IMPORTANT: This code won't compile

// The presence of an unused variable may indicate a bug, while unused
// imports just slow down compilation, an effect that can become substantial
// as a program accumulates code and programmers over time. For these reasons,
// Go refuses to compile programs with unused variables or imports, trading
// short-term convenience for long-term build speed and program clarity.
//
// More information: https://golang.org/doc/faq#unused_variables_and_imports

import "fmt" // Illegal: Unused import declaration

func main() {
	x := 1 // Illegal: Unused variable declaration
}