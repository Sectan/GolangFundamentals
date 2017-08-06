// Within a constant declaration, the predeclared identifier iota
// represents successive untyped integer constants. It is reset to 0
// whenever the reserved word const appears in the source
// and increments after each constant specification.
// More information: https://golang.org/ref/spec#Iota
// English definition of the word iota: an extremely small amount
package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

// Since iota can be part of an expression and expressions can be
// implicitly repeated, it is easy to build intricate sets of values.
// Within a parenthesized const declaration list the expression list may be omitted from any but the first declaration.
// More information: https://golang.org/ref/spec#Constant_declarations
type ByteSize float64
const (
	_           = iota // ignore first value by assigning to blank identifier
	kb ByteSize = 1 << (10 * iota)
	mb
	gb
	tb
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("")
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}