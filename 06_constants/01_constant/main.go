// A constant is just a simple, unchanging value

package main

import "fmt"

const city = "Roma"

const (
	pi       = 3.14
	language = "Go"
)

func main() {
	const gemeinde string = "Kriens"
	fmt.Println(pi)
	fmt.Println(language)
	fmt.Println(city)
	fmt.Println(gemeinde)

	const foo  = "bar" // It is allowed to declare a constant and not using it
}