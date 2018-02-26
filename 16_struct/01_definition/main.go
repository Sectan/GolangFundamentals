// A struct is a sequence of named elements, called fields, each of which has a name and a type.
// Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).
// Golang spec: https://golang.org/ref/spec#Struct_types

package main

func main() {
	// A struct with 3 fields
	type person struct {
		firstName, lastName string
		age                 int
	}
}
