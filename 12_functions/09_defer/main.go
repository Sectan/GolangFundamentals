// A defer statement defers the execution of a
// function until the surrounding function returns.
// More information:
//	- https://golang.org/ref/spec#Defer_statements

package main

import "fmt"

func withoutDefer() {
	fmt.Println("### withoutDefer")
	fmt.Println("world")
	fmt.Println("hello")

}

func withDefer() {
	fmt.Println("### withDefer")
	defer fmt.Println("world")
	fmt.Println("hello")
}

func main() {
	withoutDefer()
	withDefer()

	defer fmt.Println("Exiting func main")
	panic(-1) // Defer is executed, even if a panic occurs
}
