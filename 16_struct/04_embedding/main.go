// A field declared with a type but no explicit field name is called an embedded field.
// An embedded field must be specified as a type name T or as a pointer to a non-interface
// type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name.
// Golang spec: https://golang.org/ref/spec#Struct_types

package main

import "fmt"

func main() {
	type person struct {
		firstName, lastName string
		age                 int
	}

	type player struct {
		person
		number int
	}

	player1 := player{
		person{
			"Francesco",
			"Totti",
			41},
		10}
	fmt.Println(player1)

	player2 := player{
		person: person{
			firstName: "Daniele",
			lastName:  "de Rossi",
			age:       34},
		number: 16}
	fmt.Println(player2)

	var player3 player
	player3.person.firstName = "Alessandro"
	player3.person.lastName = "Florenzi"
	player3.person.age = 41
	player3.number = 24
	fmt.Println(player3)
}
