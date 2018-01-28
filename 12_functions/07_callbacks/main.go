// Callbacks are not idiomatic Go and should be avoided if possible

package main

import "fmt"

func printWihtCallback(values []string, callback func(string)) {
	for _, n := range values {
		callback(n)
	}
}

func main() {
	printWihtCallback([]string{"Francesco", "Daniele", "Alessandro"}, func(value string) {
		fmt.Println(value)
	})
}