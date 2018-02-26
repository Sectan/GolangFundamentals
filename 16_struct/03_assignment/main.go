// Struct fields are accessed using a dot.

package main

import "fmt"

func main() {
	type person struct {
		firstName, lastName string
		age                 int
	}

	var person1 person
	person1.firstName = "Francesco"
	person1.lastName = "Totti"
	person1.age = 41
	fmt.Println(person1)

	person2 := new(person)
	person2.firstName = "Francesco"
	person2.lastName = "Totti"
	person2.age = 41
	fmt.Println(*person2)
	// This was the short form for:
	(*person2).firstName = "Daniele"
	(*person2).lastName = "de Rossi"
	(*person2).age = 34
	fmt.Println(*person2)
}
