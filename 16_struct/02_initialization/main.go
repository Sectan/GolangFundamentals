package main

import "fmt"

func main() {
	type person struct {
		firstName, lastName string
		age                 int
	}

	// The keyword var is used to indicate that a variable is being set to its zero value.
	// Source: Bill Kennedy
	var person1 person
	fmt.Printf("Type: %T\n", person1)
	fmt.Printf("Value: %+v\n", person1)

	// We can also use the new function.
	// The new function returns a pointer (*person)
	person2 := new(person)
	fmt.Printf("Type: %T\n", person2)
	fmt.Printf("Value: %+v\n", *person2)

	// We can directly specify the values of the field
	person3 := person{"Francesco", "Totti", 41}
	fmt.Printf("Type: %T\n", person3)
	fmt.Printf("Value: %+v\n5zt", person3)

	// We can also use this way
	person4 := person{firstName: "Francesco", lastName: "Totti", age: 41}
	fmt.Printf("Type: %T\n", person4)
	fmt.Printf("Value: %+v\n5zt", person4)
}
