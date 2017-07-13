package valpkg

var FirstName string = "Francesco" // Is exported
var lastName string = "Totti"

// Within a package, package-level variables are initialized
// in declaration order but after any of the variables they depend on.
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)
// the initialization order is d, b, c, a.