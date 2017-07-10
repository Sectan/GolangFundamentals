// Programs start running in package main.
// Article on Go package names: https://blog.golang.org/package-names

package main

import "fmt"

// An import path is the string with which users import a package.
// It specifies the directory (relative to $GOROOT/src/pkg or $GOPATH/src)
// in which the package's source code resides.
// Therefore every Go project should be created inside $GOPATH/src .
// A common practice for package paths are:
//  - github.com/username/projectname/packagename
//  - projectwebsite/packagename
import "github.com/sectan/GolangFundamentals/02_packages/kriens"

// Example of the factored import statement.
// It is good style to use the factored import statement.
import (
	"github.com/sectan/GolangFundamentals/02_packages/roma"
	"github.com/sectan/GolangFundamentals/02_packages/roma/testaccio"
)

func main () {
	fmt.Println(kriens.Gemeinde)
	fmt.Println(roma.CaputMundi)
	fmt.Println(testaccio.Rione20)
}