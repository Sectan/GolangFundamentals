// https://golang.org/ref/spec#Function_literals
// https://gobyexample.com/closures
// https://newfivefour.com/golang-closures-anonymous-functions.html
// https://www.goinggo.net/2014/06/pitfalls-with-closures-in-go.html

// A function literal represents an anonymous function.
// A function literal can be assigned to a variable or invoked directly.
// More info: https://golang.org/ref/spec#Function_literals

package main

import (
	"fmt"
	"time"
)

func startGoroutine() {
	msg := "Anonymous function"

	for i := 0; i < 5; i++ {
		go func(number int) {
			fmt.Println(msg, number)
		} (i) // Here we pass the parameter to the anonymous function
	}

	// Make sure all goroutines finish properly
	time.Sleep(3 * time.Second)
}

func main()  {
	name := func() string {
		fmt.Println("inside the anonymous function")
		return "Totti"
	}
	fmt.Println(name())
	// The type of the variable is the function signature of the anonymous function
	fmt.Printf("%T\n", name)

	startGoroutine()
}