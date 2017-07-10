// Every Go program is made up of packages.
// In Go, a name is exported if it begins with a capital letter.
// When importing a package, you can refer only to its exported names.
// Any "unexported" names are not accessible from outside the package.

package kriens

// Gemeinde is exported because it starts with a capital letter.
var Gemeinde string = "Gemeinde Kriens"

// zipCode is not exported because it doesn't start with a capital letter.
var zipCode uint16 = 6010
