// A function literal can invoked directly.
// Golang Spec: https://golang.org/ref/spec#Function_literals
// Interesting stackoverflow post: https://stackoverflow.com/questions/16008604/why-add-after-closure-body-in-golang

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Anonymous function with defer")
	}()

	func() {
		fmt.Println("Anonymous function")
	}()
}
